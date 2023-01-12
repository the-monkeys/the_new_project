package main

import (
	"log"
	"net"

	"github.com/89minutes/the_new_project/auth_service/pkg/config"
	"github.com/89minutes/the_new_project/auth_service/pkg/db"
	"github.com/89minutes/the_new_project/auth_service/pkg/pb"
	"github.com/89minutes/the_new_project/auth_service/pkg/services"
	"github.com/89minutes/the_new_project/auth_service/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(cfg.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       cfg.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", cfg.AuthAddr)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	logrus.Info("starting the authentication server at address: ", cfg.AuthAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
