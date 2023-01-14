package services

import (
	"context"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"github.com/89minutes/the_new_project/services/auth_service/pkg/config"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/db"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/models"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/pb"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	H      db.Handler
	Jwt    utils.JwtWrapper
	Config config.Config
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User
	var passReset models.PasswordReset

	if result := s.H.GormConn.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email already exists",
		}, nil
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)
	user.CreateTime = timestamppb.New(time.Now()).String()
	user.UpdateTime = timestamppb.New(time.Now()).String()
	user.IsActive = true
	user.Role = int32(pb.UserRole_USER_NORMAL)

	passReset.Email = req.Email

	s.H.GormConn.Create(&user)
	s.H.GormConn.Create(&passReset)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.GormConn.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		logrus.Infof("user containing email: %s, doesn't exists", req.Email)
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "user doesn't exists",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		logrus.Infof("incorrect password given for the user containing email: %s", req.Email)
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "incorrect password",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	logrus.Infof("user containing email: %s, can successfully login", req.Email)
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := s.H.GormConn.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}

func (s *Server) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	var user models.User

	if err := s.H.Psql.QueryRow("SELECT first_name, last_name, email from users where email=$1;", req.GetEmail()).Scan(
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

	logrus.Infof("Email randomHash hash: %+v", string(randomHash))
	logrus.Infof("Email verification hash: %+v", string(emailVerifyHash))
	// TODO: start a database transaction from here till all the process are complete
	sqlStmt, err := s.H.Psql.Prepare("UPDATE password_resets SET recovery_hash=$1, time_out=$2, last_password_reset=$3 WHERE email=$4")
	if err != nil {
		logrus.Errorf("cannot prepare the reset link, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}

	result, err := sqlStmt.Exec(emailVerifyHash, time.Now().Add(time.Minute*5), time.Now(), req.Email)
	if err != nil {
		logrus.Errorf("cannot sent the reset link, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		logrus.Errorf("cannot check for affected, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}
	if affectedRows != 1 {
		logrus.Errorf("more than 1 rows are getting affected, error: %v", req.Email, err)
		return nil, status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}

	// **********************************************************************************************
	// send email with hyperlink
	// sender data
	fromEmail := s.Config.SMTPMail        //ex: "John.Doe@gmail.com"
	smtpPassword := s.Config.SMTPPassword // ex: "ieiemcjdkejspqz"
	address := s.Config.SMTPAddress
	to := []string{req.Email}

	subject := "Subject: The Monkeys Account Recovery\n"

	emailBody := utils.ResetPasswordTemplate(user.FirstName, user.LastName, user.Email, string(randomHash))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + emailBody)

	auth := smtp.PlainAuth("", fromEmail, smtpPassword, s.Config.SMTPHost)

	// fmt.Println("message:", string(message))
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
