package server

import (
	"context"
	"io"
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

func (us *UserService) UploadProfile(stream pb.UserService_UploadProfileServer) error {
	var imageData []byte
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		imageData = append(imageData, chunk.Data...)
	}

	err := us.db.UploadProfilePic(imageData, 2)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.ProfileId{Id: 1})
}
func (us *UserService) Download(req *pb.ProfileId, stream pb.UserService_DownloadServer) error {
	xb := []byte{}
	if err := us.db.Psql.QueryRow("SELECT profile_pic from the_monkeys_user WHERE id=$1", 2).Scan(&xb); err != nil {
		us.log.Errorf("cannot get the profile pic, error: %v", err)
		return err
	}

	if err := stream.Send(&pb.ProfilePicChunk{
		Data: xb,
	}); err != nil {
		us.log.Errorf("error while sending stream, error %+v", err)
	}

	return nil
}
