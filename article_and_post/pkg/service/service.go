package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/89minutes/the_new_project/article_and_post/pkg/database"
	"github.com/89minutes/the_new_project/article_and_post/pkg/models"
	"github.com/89minutes/the_new_project/article_and_post/pkg/pb"
	"github.com/89minutes/the_new_project/article_and_post/pkg/utils"
	"github.com/google/uuid"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
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

	if req.GetId() == "" {
		req.Id = uuid.New().String()
	}

	req.CanEdit = true
	req.ContentOwnerShip = pb.CreateArticleRequest_THE_USER

	// Store into the opensearch db
	document := strings.NewReader(ArticleToString(req))

	osReq := opensearchapi.IndexRequest{
		Index:      utils.OpensearchArticleIndex,
		DocumentID: req.Id,
		Body:       document,
	}

	insertResponse, err := osReq.Do(context.Background(), srv.osClient)
	if err != nil {
		logrus.Errorf("cannot create a new document for user: %s, error: %v", req.GetAuthor(), err)
		return nil, err
	}

	if insertResponse.IsError() {
		logrus.Errorf("opensearch apt failed to create a new document for user: %s, error: %v",
			req.GetAuthor(), insertResponse.Status())
		return nil, err
	}

	logrus.Infof("successfully created an article for user: %s, insert response: %+v",
		req.Author, insertResponse)

	return &pb.CreateArticleResponse{
		Status: http.StatusCreated,
		Id:     article.Id,
	}, nil
}

func ArticleToString(ip *pb.CreateArticleRequest) string {
	return fmt.Sprintf(`{
		"id":         			"%s",
		"title":      			"%s",
		"content":     			"%s",
		"author":   			"%s",
		"is_draft":    			"%v",
		"tags": 				"%v",
		"create_time": 			"%v",
		"update_time": 			"%v",
		"quick_read": 			"%v",
		"content_ownership": 	"%v",
		"can_edit": 			"%v",
		"viewed_by":			"%v",
		"comments":				"%v"
	}`, ip.Id, ip.Title, ip.Content, ip.Author, ip.IsDraft,
		ip.Tags, ip.CreateTime, ip.UpdateTime, ip.QuickRead, ip.ContentOwnerShip,
		ip.CanEdit, ip.ViewBy, ip.Comment)
}
