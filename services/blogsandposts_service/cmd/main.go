package main

import (
	"log"
	"net"

	isv "github.com/89minutes/the_new_project/apis/interservice/blogs/pb"
	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/config"
	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/pb"
	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadArtNPostConfig()

	if err != nil {
		log.Fatalln("failed to load the config file, error: ", err)
	}

	lis, err := net.Listen("tcp", cfg.BlogAndPostSvcURL)
	if err != nil {
		log.Fatalf("article and service server failed to listen at port %v, error: %v",
			cfg.BlogAndPostSvcURL, err)
	}

	logger := logrus.New()

	osClient, err := service.NewOpenSearchClient(cfg.OSAddress, cfg.OSUsername, cfg.OSPassword, logger)
	if err != nil {
		logger.Fatalf("cannot get the opensearch client, error: %v", err)
	}

	blogService := service.NewBlogService(*osClient, logger)

	grpcServer := grpc.NewServer()

	pb.RegisterBlogsAndPostServiceServer(grpcServer, blogService)
	isv.RegisterBlogServiceServer(grpcServer, nil)

	logrus.Info("art and post service is running on address: ", cfg.BlogAndPostSvcURL)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
