package main

import (
	"fmt"
	"log"
	"net"

	"github.com/89minutes/the_new_project/article_and_post/pkg/config"
	"github.com/89minutes/the_new_project/article_and_post/pkg/pb"
	"github.com/89minutes/the_new_project/article_and_post/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", cfg.ArticleServerPort)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Article Service is on port: ", cfg.ArticleServerPort)

	articleServer, err := service.NewArticleServer(cfg.OSAddress, cfg.OSUsername, cfg.OSPassword)
	grpcServer := grpc.NewServer()

	pb.RegisterArticleServiceServer(grpcServer, articleServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
