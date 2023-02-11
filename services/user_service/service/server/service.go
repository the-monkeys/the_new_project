package server

import (
	"context"
	"log"
	"net/http"

	"database/sql"

	"github.com/89minutes/the_new_project/services/user_service/service/database"
	"github.com/89minutes/the_new_project/services/user_service/service/pb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	db  *database.UserDbHandler
	log *logrus.Logger
	pb.UnimplementedUserServiceServer
}

func NewUserService(usd *database.UserDbHandler, log *logrus.Logger) *UserService {
	return &UserService{db: usd, log: log}
}

func (us *UserService) GetMyProfile(ctx context.Context, req *pb.GetMyProfileReq) (*pb.GetMyProfileRes, error) {
	us.log.Infof("user %v has requested for their profile", req.GetId())

	resp, err := us.db.GetMyProfile(req.GetId())
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			us.log.Infof("cannot fine the user with id %v ", req.GetId())
			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
		case sql.ErrTxDone:
			log.Println("The transaction has already been committed or rolled back.")
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		case sql.ErrConnDone:
			log.Println("The database connection has been closed.")
			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
		default:
			log.Printf("An internal server error occurred: %v\n", err)
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		}
	}

	return resp, nil
}

func (us *UserService) SetMyProfile(ctx context.Context, req *pb.SetMyProfileReq) (*pb.SetMyProfileRes, error) {
	us.log.Infof("the user %s has requested to update profile", req.GetEmail())
	if err := us.db.UpdateMyProfile(req); err != nil {
		switch err {
		case sql.ErrNoRows:
			us.log.Infof("cannot fine the user with id %v ", req.GetEmail())
			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
		case sql.ErrTxDone:
			log.Println("The transaction has already been committed or rolled back.")
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		case sql.ErrConnDone:
			log.Println("The database connection has been closed.")
			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
		default:
			log.Printf("An internal server error occurred: %v\n", err)
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		}
	}
	return &pb.SetMyProfileRes{
		Status: http.StatusOK,
	}, nil
}
