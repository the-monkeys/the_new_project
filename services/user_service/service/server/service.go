package server

import (
	"context"

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
		if err == sql.ErrNoRows {
			us.log.Infof("cannot fine the user with id %v ", req.GetId())
			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
		}

		us.log.Errorf("cannot get user profile for id %v, error: %v", req.GetId(), err)
		return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
	}
	return resp, nil
}
