package article

import (
	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/article/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ArticleServiceClient struct {
	Client pb.ArticleServiceClient
	Log    *logrus.Logger
}

func InitArticleServiceClient(c *config.Config) pb.ArticleServiceClient {
	// using WithInsecure() because no SSL running
	logrus.Infof("Dialing to grpc art service: %v", c.ArticleSvcUrl)
	cc, err := grpc.Dial(c.ArticleSvcUrl, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("Could not connect to the art and post grpc server, error: %v", err)
		return nil
	}

	return pb.NewArticleServiceClient(cc)
}
