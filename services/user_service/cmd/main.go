package main

import (
	"net"

	"github.com/89minutes/the_new_project/services/user_service/service/config"
	"github.com/89minutes/the_new_project/services/user_service/service/database"
	"github.com/89minutes/the_new_project/services/user_service/service/pb"
	"github.com/89minutes/the_new_project/services/user_service/service/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadUserConfig()
	if err != nil {
		logrus.Errorf("failed to load user config, error: %+v", err)
	}
	log := logrus.New()

	db := database.NewUserDbHandler(cfg.DBUrl, log)

	lis, err := net.Listen("tcp", cfg.UserSrvPort)
	if err != nil {
		log.Errorf("failed to listen at port %v, error: %+v", cfg.UserSrvPort, err)
	}

	userService := server.NewUserService(db, log)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userService)
	if err = BlogServiceConn(cfg.BlogAndPostSvcURL); err != nil {
		logrus.Fatalln("could not connect to the blog service: %v", err)
	}

	log.Infof("the user service started at: %v", cfg.UserSrvPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

func BlogServiceConn(addr string) error {
	logrus.Infof("gRPC dialing to the blog server: %v", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}
