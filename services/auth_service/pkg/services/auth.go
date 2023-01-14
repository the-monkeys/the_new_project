package services

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/89minutes/the_new_project/services/auth_service/pkg/db"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/models"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/pb"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User
	var passReset models.PasswordReset

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
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

	s.H.DB.Create(&user)
	s.H.DB.Create(&passReset)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
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

	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
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

func (s *Server) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.ResetPasswordRes, error) {
	var user models.User
	var pass models.PasswordReset
	// Check if email exists or not
	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.ResetPasswordRes{
			Status: http.StatusNotFound,
			Error:  "the email doesn't exist",
		}, nil
	}
	//

	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	emailVerRandRune := make([]rune, 64)
	for i := 0; i < 64; i++ {
		// Intn() returns, as an int, a non-negative pseudo-random number in [0,n).
		emailVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}

	emailVerPassword := string(emailVerRandRune)
	// var emailVerPWhash []byte
	// generate emailVerPassword hash for db
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	emailVerPWhash, err := bcrypt.GenerateFromPassword([]byte(emailVerPassword), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("bcrypt err:", err)
		return nil, err
	}

	if res := s.H.DB.Model(&pass).Where("email = ?", req.Email).Updates(map[string]interface{}{"recovery_hash": string(emailVerPWhash),
		"time_out": time.Now().Add(time.Minute * 5), "last_pass_reset": time.Now()}); res.Error != nil {
		return &pb.ResetPasswordRes{
			Status: http.StatusNotFound,
			Error:  "the email doesn't exist",
		}, nil
	}

	// **********************************************************************************************
	// send email with hyperlink
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
	// receiver address privided through toEmail argument
	to := []string{req.Email}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: The Monkeys Account Recovery\n"
	// localhost:8080 will be removed by many email service but works with online sites
	// https must be used since we are sending personal data through url parameters
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"http://localhost:5001/forgotpwchange?u=" + req.Email + "&evpw=" + emailVerPassword + "\">Change Password</a></body>"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	fmt.Println("message:", string(message))
	err = smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("error sending reset password email, err:", err)

	}

	return &pb.ResetPasswordRes{
		Status: 200,
	}, nil
}
