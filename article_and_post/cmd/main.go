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

	// h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", cfg.ArticleServerPort)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", cfg.ArticleServerPort)

	// s := services.Server{
	// 	H: h,
	// }
	articleServer, err := service.NewArticleServer(cfg.OSAddress, cfg.Username, cfg.Password)
	grpcServer := grpc.NewServer()

	pb.RegisterArticleServiceServer(grpcServer, articleServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
