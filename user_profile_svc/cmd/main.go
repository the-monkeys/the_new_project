package main

import (
	"log"
	"net"

	"github.com/89minutes/the_new_project/user_profile_svc/user_service/config"
	"github.com/89minutes/the_new_project/user_profile_svc/user_service/database"
	"github.com/89minutes/the_new_project/user_profile_svc/user_service/pb"
	"github.com/89minutes/the_new_project/user_profile_svc/user_service/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadUserConfig()
	if err != nil {
		logrus.Errorf("failed to load user config, error: %+v", err)
	}

	handler := database.Init(cfg.DBUrl)

	lis, err := net.Listen("tcp", cfg.UserSrvPort)
	if err != nil {
		logrus.Errorf("failed to listen at port %v, error: %+v", cfg.UserSrvPort, err)
	}

	userService := service.UserService{
		DbClient: handler,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &userService)

	logrus.Infof("user service started at port: %v", cfg.UserSrvPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
