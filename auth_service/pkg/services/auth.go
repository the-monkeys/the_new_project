package services

import (
	"context"
	"net/http"

	"github.com/89minutes/the_new_project/auth_service/pkg/db"
	"github.com/89minutes/the_new_project/auth_service/pkg/models"
	"github.com/89minutes/the_new_project/auth_service/pkg/pb"
	"github.com/89minutes/the_new_project/auth_service/pkg/utils"
	"github.com/sirupsen/logrus"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	s.H.DB.Create(&user)

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
