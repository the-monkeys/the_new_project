package main

import (
	"net"

	"github.com/89minutes/the_new_project/services/article_and_post/pkg/config"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/pb"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadArtNPostConfig()

	if err != nil {
		logrus.Fatalln("failed to load the config file, error: ", err)
	}

	lis, err := net.Listen("tcp", cfg.ArticleSvcUrl)
	if err != nil {
		logrus.Fatalf("article and service server failed to listen at port %v, error: %v",
			cfg.ArticleServerPort, err)
	}

	// artServer, err := service.NewArticleServer(cfg.OSAddress, cfg.OSUsername, cfg.OSPassword, logrus.New())
	// if err != nil {
	// 	artServer.Log.Errorf("cannot create to the art server, error: %v", err)
	// }

	client, _ := service.NewOpenSearchClient(cfg.OSAddress, cfg.OSUsername, cfg.OSPassword, logrus.New())
	s := service.ArticleServer{
		Log:      logrus.New(),
		OsClient: client,
	}
	// artServer.Log.SetReportCaller(true)
	// artServer.Log.SetFormatter(&logrus.TextFormatter{
	// 	DisableColors: false,
	// 	FullTimestamp: false,
	// })
	grpcServer := grpc.NewServer()

	pb.RegisterArticleServiceServer(grpcServer, &s)

	logrus.Info("art and post service is running on address: ", cfg.ArticleSvcUrl)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalln("Failed to serve:", err)
	}
}
