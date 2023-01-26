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
	"github.com/google/uuid"
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
	err := s.dbCli.PsqlClient.QueryRow("SELECT email FROM the_monkeys_user WHERE email=$1;", req.GetEmail()).
		Scan(&user.Email)
	if err == nil {
		logrus.Errorf("cannot register the user, as the email %s is existing already", req.Email)
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "the email is already registered",
		}, nil
	}

	user.UUID = uuid.NewString()
	user.FirstName = req.GetFirstName()
	user.LastName = req.GetLastName()
	user.Email = req.GetEmail()
	user.Password = utils.HashPassword(req.Password)
	user.CreateTime = time.Now().Format(common.DATE_TIME_FORMAT)
	user.UpdateTime = time.Now().Format(common.DATE_TIME_FORMAT)
	user.IsActive = true
	user.Role = int32(pb.UserRole_USER_NORMAL)

	logrus.Infof("registering the user with email %v", req.Email)
	// else create the user
	if err := s.dbCli.RegisterUser(user); err != nil {
		return nil, err
	}
	logrus.Infof("user %s is successfully registered.", user.Email)
	return &pb.RegisterResponse{Status: http.StatusOK, Error: "registered successfully"}, nil
}

// func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
// 	var user models.TheMonkeysUser

// 	if result := s.dbCli.GormClient.Where(&models.TheMonkeysUser{Email: req.Email}).First(&user); result.Error == nil {
// 		return &pb.RegisterResponse{
// 			Status: http.StatusConflict,
// 			Error:  "email already exists",
// 		}, nil
// 	}

// 	return &pb.RegisterResponse{
// 		Status: http.StatusOK,
// 		Error:  "no error",
// 	}, nil

// 	user.FirstName = req.FirstName
// 	user.LastName = req.LastName
// 	user.Email = req.Email
// 	user.Password = utils.HashPassword(req.Password)
// 	user.CreateTime = time.Now().Format(common.DATE_TIME_FORMAT)
// 	user.UpdateTime = time.Now().Format(common.DATE_TIME_FORMAT)
// 	user.IsActive = true
// 	user.Role = int32(pb.UserRole_USER_NORMAL)

// 	s.dbCli.GormClient.Create(&user)

// 	return &pb.RegisterResponse{
// 		Status: http.StatusCreated,
// 	}, nil
// }

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.TheMonkeysUser

	// Check if the email exists
	err := s.dbCli.PsqlClient.QueryRow("SELECT email, password FROM the_monkeys_user WHERE email=$1;", req.GetEmail()).
		Scan(&user.Email, &user.Password)
	if err != nil {
		logrus.Errorf("cannot login as the email %s doesn't exist, error: %+v", req.Email, err)
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

// func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
// 	var user models.TheMonkeysUser

// 	if result := s.dbCli.GormClient.Where(&models.TheMonkeysUser{Email: req.Email}).First(&user); result.Error != nil {
// 		logrus.Infof("user containing email: %s, doesn't exists", req.Email)
// 		return &pb.LoginResponse{
// 			Status: http.StatusNotFound,
// 			Error:  "user doesn't exists",
// 		}, nil
// 	}

// 	match := utils.CheckPasswordHash(req.Password, user.Password)

// 	if !match {
// 		logrus.Infof("incorrect password given for the user containing email: %s", req.Email)
// 		return &pb.LoginResponse{
// 			Status: http.StatusBadRequest,
// 			Error:  "incorrect password",
// 		}, nil
// 	}

// 	token, _ := s.jwt.GenerateToken(user)

// 	logrus.Infof("user containing email: %s, can successfully login", req.Email)
// 	return &pb.LoginResponse{
// 		Status: http.StatusOK,
// 		Token:  token,
// 	}, nil
// }

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

	// if result := s.dbCli.GormClient.Where(&models.TheMonkeysUser{Email: claims.Email}).First(&user); result.Error != nil {
	// 	return &pb.ValidateResponse{
	// 		Status: http.StatusNotFound,
	// 		Error:  "User not found",
	// 	}, nil
	// }

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}

func (s *AuthServer) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	var user models.TheMonkeysUser

	if err := s.dbCli.PsqlClient.QueryRow("SELECT first_name, last_name, email from users where email=$1;", req.GetEmail()).Scan(
		&user.FirstName, &user.LastName, &user.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			logrus.Errorf("cannot find the email %s, error: %v", req.Email, err)
			return nil, status.Errorf(codes.NotFound, "the email isn't registered: %v", err)
		}

		logrus.Errorf("cannot fine the email %s, internal server error, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}

	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	randomHash := make([]rune, 64)
	for i := 0; i < 64; i++ {
		// Intn() returns, as an int, a non-negative pseudo-random number in [0,n).
		randomHash[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}

	emailVerifyHash := utils.HashPassword(string(randomHash))

	// TODO: start a database transaction from here till all the process are complete
	sqlStmt, err := s.dbCli.PsqlClient.Prepare("UPDATE password_resets SET recovery_hash=$1, time_out=$2, last_password_reset=$3 WHERE email=$4")
	if err != nil {
		logrus.Errorf("cannot prepare the reset link for %s, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}

	result, err := sqlStmt.Exec(emailVerifyHash, time.Now().Add(time.Minute*5), time.Now(), req.Email)
	if err != nil {
		logrus.Errorf("cannot sent the reset link for %s, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		logrus.Errorf("cannot check for affected rows for %s, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}
	if affectedRows != 1 {
		logrus.Errorf("more than 1 rows are getting affected for %s, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}

	// **********************************SEND EMAIL WITH PW RESET LINK***************************************
	fromEmail := s.config.SMTPMail        //ex: "John.Doe@gmail.com"
	smtpPassword := s.config.SMTPPassword // ex: "ieiemcjdkejspqz"
	address := s.config.SMTPAddress
	to := []string{req.Email}

	subject := "Subject: The Monkeys Account Recovery\n"

	emailBody := utils.ResetPasswordTemplate(user.FirstName, user.LastName, user.Email, string(randomHash))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + emailBody)

	auth := smtp.PlainAuth("", fromEmail, smtpPassword, s.config.SMTPHost)

	if err = smtp.SendMail(address, auth, fromEmail, to, message); err != nil {
		logrus.Errorf("error occurred while sending verification email, error: %+v", err)
		return &pb.ForgotPasswordRes{
			Status: int64(codes.Internal),
			Error:  "cannot send email, please provide correct email id",
		}, nil

	}

	return &pb.ForgotPasswordRes{
		Status: 200,
		Error:  "",
	}, nil
}

func (s *AuthServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.ResetPasswordRes, error) {
	var pwr models.PasswordReset
	var user models.TheMonkeysUser
	var timeOut string
	if err := s.dbCli.PsqlClient.QueryRow("SELECT email,recovery_hash, time_out FROM password_resets WHERE email=$1;", req.GetEmail()).
		Scan(&pwr.Email, &pwr.RecoveryHash, &timeOut); err != nil {
		logrus.Errorf("cannot get the password recovery details, error: %v", err)
		return nil, err
	}

	logrus.Infof("PWR: %+v", pwr)
	logrus.Infof("timeOut: %+v", timeOut)

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

	if err := s.dbCli.PsqlClient.QueryRow("SELECT id, email FROM users WHERE email=$1;", req.GetEmail()).
		Scan(&user.Id, &user.Email); err != nil {
		logrus.Errorf("cannot get the password recovery details, error: %v", err)
		return nil, err
	}

	token, _ := s.jwt.GenerateToken(user)
	return &pb.ResetPasswordRes{
		Status: 200,
		Error:  "",
		Token:  token,
	}, nil
}
