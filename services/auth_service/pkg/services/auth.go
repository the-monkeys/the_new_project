package services

import (
	"context"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"github.com/89minutes/the_new_project/common"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/config"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/db"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/models"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/pb"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	dbCli  *db.AuthDBHandler
	jwt    utils.JwtWrapper
	config config.Config
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(dbCli *db.AuthDBHandler, jwt utils.JwtWrapper, config config.Config) *AuthServer {
	return &AuthServer{
		dbCli:  dbCli,
		jwt:    jwt,
		config: config,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.TheMonkeysUser

	// Check if the user exists with the same email id return conflict
	err := s.dbCli.PsqlClient.QueryRow("SELECT email, deactivated FROM the_monkeys_user WHERE email=$1;", req.GetEmail()).
		Scan(&user.Email, &user.Deactivated)
	if err == nil {
		logrus.Errorf("cannot register the user, as the email %s is existing already", req.Email)
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "the email is already registered",
		}, nil
	}

	hash := string(utils.GenHash())
	encHash := utils.HashPassword(hash)

	user.UUID = utils.GetUUID()
	user.FirstName = req.GetFirstName()
	user.LastName = req.GetLastName()
	user.Email = req.GetEmail()
	user.Password = utils.HashPassword(req.Password)
	user.CreateTime = time.Now().Format(common.DATE_TIME_FORMAT)
	user.UpdateTime = time.Now().Format(common.DATE_TIME_FORMAT)
	user.IsActive = true
	user.Role = int32(pb.UserRole_USER_NORMAL)
	user.EmailVerificationToken = encHash
	user.EmailVerificationTimeout = time.Now().Add(time.Hour * 24)
	user.Deactivated = false
	user.LoginMethod = pb.RegisterRequest_LoginMethod_name[pb.RegisterRequest_LoginMethod_value[req.LoginMethod.String()]]

	logrus.Infof("registering the user with email %v", req.Email)
	if err := s.dbCli.RegisterUser(user); err != nil {
		return nil, err
	}

	// Send email verification mail as a routine else the register api gets slower
	emailBody := utils.EmailVerificationHTML(user.Email, hash)
	go s.SendMail(user.Email, emailBody)

	logrus.Infof("user %s is successfully registered.", user.Email)

	// Generate and return token
	token, err := s.jwt.GenerateToken(user)
	if err != nil {
		logrus.Errorf("cannot create a token for %s, error: %+v", req.Email, err)
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "something went wrong",
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		Error:  "",
		Token:  token,
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	logrus.Infof("user %s has requested to login", req.Email)
	var user models.TheMonkeysUser

	// Check if the email exists
	err := s.dbCli.PsqlClient.QueryRow("SELECT email, password, deactivated FROM the_monkeys_user WHERE email=$1;", req.GetEmail()).
		Scan(&user.Email, &user.Password, &user.Deactivated)
	if err != nil {
		logrus.Errorf("cannot login as the email %s doesn't exist, error: %+v", req.Email, err)
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "the email isn't registered",
		}, nil
	}

	// If user is deactivated then no login
	if user.Deactivated == true {
		logrus.Errorf("user %s cannot login as it's deactivated, error: %+v", req.Email, err)
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "the email isn't registered",
		}, nil
	}

	// Check if the password match
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		logrus.Errorf("cannot login as the email/password doesn't match for: %s", req.Email)
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "the email/password incorrect",
		}, nil
	}

	// Generate and return token
	token, err := s.jwt.GenerateToken(user)
	if err != nil {
		logrus.Errorf("cannot create a token for %s, error: %+v", req.Email, err)
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "the email/password incorrect",
		}, nil
	}

	logrus.Infof("user containing email: %s, has been assigned a token", req.Email)
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *AuthServer) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.jwt.ValidateToken(req.Token)
	if err != nil {
		logrus.Errorf("cannot validate the json token, error: %v", err)
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.TheMonkeysUser
	// Check if the email exists
	if err := s.dbCli.PsqlClient.QueryRow("SELECT email, password FROM the_monkeys_user WHERE email=$1;", claims.Email).
		Scan(&user.Email, &user.Password); err != nil {
		logrus.Errorf("cannot validate token as the email %s doesn't exist, error: %+v", claims.Email, err)
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
		User:   claims.Email,
	}, nil
}

func (s *AuthServer) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	logrus.Infof("user %s has forgotten their password", req.Email)

	user, err := s.dbCli.GetNamesEmailFromEmail(req)
	if err != nil {
		logrus.Errorf("error occurred while finding the user %s, error: %v", req.Email, err)
		return nil, err
	}

	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	randomHash := make([]rune, 64)
	for i := 0; i < 64; i++ {
		// Intn() returns, as an int, a non-negative pseudo-random number in [0,n).
		randomHash[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}

	emailVerifyHash := utils.HashPassword(string(randomHash))

	if err = s.dbCli.UpdatePasswordRecoveryToken(emailVerifyHash, user); err != nil {
		logrus.Errorf("error occurred while updating email verification token for %s, error: %v", req.Email, err)
		return nil, err
	}

	emailBody := utils.ResetPasswordTemplate(user.FirstName, user.LastName, string(randomHash), user.Id)
	go s.SendMail(req.Email, emailBody)

	return &pb.ForgotPasswordRes{
		Status: 200,
		Error:  "",
	}, nil
}

func (s *AuthServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.ResetPasswordRes, error) {
	logrus.Infof("user %d has requested to reset their password", req.Id)
	var pwr models.PasswordReset
	var user models.TheMonkeysUser
	var timeOut string
	if err := s.dbCli.PsqlClient.QueryRow("SELECT email,recovery_hash, time_out FROM pw_reset WHERE user_id=$1 ORDER BY id DESC LIMIT 1; ", req.GetId()).
		Scan(&pwr.Email, &pwr.RecoveryHash, &timeOut); err != nil {
		logrus.Errorf("cannot get the password recovery details, error: %v", err)
		return nil, err
	}

	// TODO: Remove the following two line
	// logrus.Infof("PWR: %+v", pwr)
	// logrus.Infof("timeOut: %+v", timeOut)

	timeTill, err := time.Parse(time.RFC3339, timeOut)

	if timeTill.Before(time.Now()) {
		logrus.Errorf("the token has already expired, error: %+v", err)
		return nil, status.Errorf(codes.Unauthenticated, "token expired already")
	}

	// Verify reset token
	if ok := utils.CheckPasswordHash(req.Token, pwr.RecoveryHash); !ok {
		logrus.Errorf("the token didn't match, error: %+v", err)
		return nil, status.Errorf(codes.Unauthenticated, "token didn't match")
	}

	if err := s.dbCli.PsqlClient.QueryRow("SELECT id, email FROM the_monkeys_user WHERE id=$1;", req.GetId()).
		Scan(&user.Id, &user.Email); err != nil {
		logrus.Errorf("cannot get the password recovery details, error: %v", err)
		return nil, err
	}

	token, _ := s.jwt.GenerateToken(user)
	logrus.Infof("password is set and token generated for %d", req.Id)

	return &pb.ResetPasswordRes{
		Status: 200,
		Error:  "",
		Token:  token,
	}, nil
}

func (srv *AuthServer) SendMail(email, emailBody string) error {
	logrus.Infof("Send mail routine triggered")

	fromEmail := srv.config.SMTPMail        //ex: "John.Doe@gmail.com"
	smtpPassword := srv.config.SMTPPassword // ex: "ieiemcjdkejspqz"
	address := srv.config.SMTPAddress
	to := []string{email}

	subject := "Subject: The Monkeys support\n"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + emailBody)

	auth := smtp.PlainAuth("", fromEmail, smtpPassword, srv.config.SMTPHost)

	if err := smtp.SendMail(address, auth, fromEmail, to, message); err != nil {
		logrus.Errorf("error occurred while sending verification email, error: %+v", err)
		return nil

	}

	return nil
}

func (s *AuthServer) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordReq) (*pb.UpdatePasswordRes, error) {
	logrus.Infof("updating password for: %+v", req.Email)

	encHash := utils.HashPassword(req.Password)
	if err := s.dbCli.UpdatePassword(encHash, req.Email); err != nil {
		return nil, err
	}
	logrus.Infof("updated password for: %+v", req.Email)
	return &pb.UpdatePasswordRes{
		Status: http.StatusOK,
	}, nil
}

// Verify email
func (s *AuthServer) VerifyEmail(ctx context.Context, req *pb.VerifyEmailReq) (*pb.VerifyEmailRes, error) {
	var user models.TheMonkeysUser
	var timeOut string
	logrus.Infof("verifying email: %s", req.GetEmail())

	if err := s.dbCli.PsqlClient.QueryRow(`SELECT email, email_verification_token, email_verification_timeout 
		FROM the_monkeys_user WHERE email=$1;`, req.GetEmail()).
		Scan(&user.Email, &user.EmailVerificationToken, &timeOut); err != nil {
		logrus.Errorf("cannot get the user details to verify email, error: %v", err)
		return nil, err
	}
	timeTill, err := time.Parse(time.RFC3339, timeOut)

	if timeTill.Before(time.Now()) {
		logrus.Errorf("the token has already expired, error: %+v", err)
		return nil, status.Errorf(codes.Unauthenticated, "token expired already")
	}

	// Verify reset token
	if ok := utils.CheckPasswordHash(req.Token, user.EmailVerificationToken); !ok {
		logrus.Errorf("the token didn't match, error: %+v", err)
		return nil, status.Errorf(codes.Unauthenticated, "token didn't match")
	}

	res, err := s.dbCli.PsqlClient.Exec("UPDATE the_monkeys_user SET email_verified=true WHERE email=$1;", req.GetEmail())
	if err != nil {
		logrus.Errorf("cannot update the verification details for %s, error: %v", req.Email, err)
		return nil, err
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %s, error: %v", user.Email, err)
		return nil, err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %s, error: %v", user.Email, err)
		return nil, err
	}

	logrus.Infof("verified email: %s", req.GetEmail())
	return &pb.VerifyEmailRes{
		Status: 200,
		Error:  "",
	}, nil

}

func (s *AuthServer) RequestForEmailVerification(ctx context.Context, req *pb.EmailVerificationReq) (*pb.EmailVerificationRes, error) {
	if req.Email == "" {
		return nil, common.BadRequest
	}
	logrus.Infof("user %v has requested for email verification", req.Email)
	var user models.TheMonkeysUser
	var timeOut string

	if err := s.dbCli.PsqlClient.QueryRow(`SELECT id, email, email_verification_token, email_verification_timeout 
		FROM the_monkeys_user WHERE email=$1;`, req.GetEmail()).
		Scan(&user.Id, &user.Email, &user.EmailVerificationToken, &timeOut); err != nil {
		logrus.Errorf("cannot get the user details to verify email, error: %v", err)
		return nil, err
	}

	logrus.Infof("generating verification email token for: %s", req.GetEmail())
	hash := string(utils.GenHash())
	encHash := utils.HashPassword(hash)

	user.EmailVerificationToken = encHash
	user.EmailVerificationTimeout = time.Now().Add(time.Hour * 24)

	if err := s.dbCli.UpdateEmailVerToken(user); err != nil {
		return nil, err
	}

	emailBody := utils.EmailVerificationHTML(user.Email, hash)
	logrus.Infof("Sending verification email to: %s", req.GetEmail())
	// TODO: Handle error of the go routine
	go s.SendMail(user.Email, emailBody)

	return &pb.EmailVerificationRes{
		Status:  http.StatusOK,
		Message: "Check your email and click on the verify link",
	}, nil
}
