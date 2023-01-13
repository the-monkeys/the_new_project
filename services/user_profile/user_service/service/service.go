package service

import (
	"context"
	"errors"

	"github.com/89minutes/the_new_project/services/user_profile/user_service/database"
	"github.com/89minutes/the_new_project/services/user_profile/user_service/pb"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	DbClient database.UserHandler
	pb.UnimplementedUserServiceServer
}

func (us *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileReq) (*pb.GetUserProfileRes, error) {

	res := &pb.GetUserProfileRes{}
	err := us.DbClient.Psql.QueryRow("select id, first_name, last_name, email, profile_pic from users where id=$1;", req.GetId()).Scan(
		&res.Id, &res.FirstName, &res.LastName, &res.Email, &res.ProfilePic)
	if err != nil {
		logrus.Infof("cannot get profile for user containing id: %v, error: ", req.GetId(), err)
		return nil, errors.New("cannot get the profile")
	}

	logrus.Infof("fetched profile for the user containing id: %d", res.Id)
	return res, nil
}
