package service

import (
	"context"
	"errors"

	"github.com/89minutes/the_new_project/user_profile_svc/user_service/database"
	"github.com/89minutes/the_new_project/user_profile_svc/user_service/models"
	"github.com/89minutes/the_new_project/user_profile_svc/user_service/pb"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	DbClient database.UserHandler
	pb.UnimplementedUserServiceServer
}

func (us *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileReq) (*pb.GetUserProfileRes, error) {
	var user models.User

	if result := us.DbClient.DB.Where(&models.User{Id: req.GetId()}).First(&user); result.Error != nil {
		logrus.Infof("profile for user containing email: %s, doesn't exists", user.Email)

		return nil, errors.New("cannot get the profile")
	}

	logrus.Infof("got the user: %+v", user)
	return &pb.GetUserProfileRes{}, nil
}
