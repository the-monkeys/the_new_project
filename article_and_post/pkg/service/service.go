package service

import (
	"context"
	"net/http"

	"github.com/89minutes/the_new_project/article_and_post/pkg/database"
	"github.com/89minutes/the_new_project/article_and_post/pkg/models"
	"github.com/89minutes/the_new_project/article_and_post/pkg/pb"
	"github.com/opensearch-project/opensearch-go"
	"github.com/sirupsen/logrus"
)

type ArticleServer struct {
	osClient *opensearch.Client
	pb.UnimplementedArticleServiceServer
}

func NewArticleServer(url, username, password string) (*ArticleServer, error) {
	client, err := database.NewOSClient(url, username, password)
	if err != nil {
		logrus.Errorf("Failed to connect to opensearch instance, error: %+v", err)
		return nil, err
	}

	return &ArticleServer{
		osClient: client,
	}, nil
}

func (srv *ArticleServer) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	var article models.Article

	article.Title = req.Title
	article.Content = req.Content
	article.Author = req.Author

	// Store into the opensearch db
	logrus.Info("Getting the article: %+v", article)
	// srv.osClient.Create("article")
	// if result := s.H.DB.Create(&product); result.Error != nil {
	// 	return &pb.CreateProductResponse{
	// 		Status: http.StatusConflict,
	// 		Error:  result.Error.Error(),
	// 	}, nil
	// }

	return &pb.CreateArticleResponse{
		Status: http.StatusCreated,
		Id:     article.Id,
	}, nil
}
