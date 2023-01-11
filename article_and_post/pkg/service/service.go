package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func (srv *ArticleServer) GetArticles(req *pb.GetArticlesRequest, stream pb.ArticleService_GetArticlesServer) error {
	// TODO: Get All the articles and stream in the for loop

	// Search for the document.
	content := strings.NewReader(`{
		"query": {
			"match": {
				"is_draft": "false"
			}
		},
		"_source": {
			"includes": [
				"id",
				"title",
				"author",
				"create_time",
				"quick_read",
				"viewed_by"
			],
			"excludes": [
				"content"
			]
		}
	}`)

	search := opensearchapi.SearchRequest{
		Index: []string{utils.OpensearchArticleIndex},
		Body:  content,
	}

	searchResponse, err := search.Do(context.Background(), srv.osClient)
	if err != nil {
		fmt.Println("failed to search document ", err)
		os.Exit(1)
	}
	fmt.Println("Searching for a document")
	// fmt.Println(searchResponse.)
	var result map[string]interface{}
	// resp := []pb.GetArticlesResponse{}

	decoder := json.NewDecoder(searchResponse.Body)
	if err := decoder.Decode(&result); err != nil {
		logrus.Error("Error while decoding, error", err)
	}

	documents := result["hits"].(map[string]interface{})["hits"].([]interface{})
	fmt.Println("Length of doc: ", len(documents))
	for _, doc := range documents {
		resp := ParseToStruct(doc)
		if err := stream.Send(&resp); err != nil {
			logrus.Errorf("error while sending stream, error %+v", err)
		}
	}

	return nil
}

func ParseToStruct(result interface{}) pb.GetArticlesResponse {
	// logrus.Infof("Struct**********: %+v", result.(map[string]interface{}))
	instance := models.GetArticleResp{}
	layer1 := result.(map[string]interface{})
	logrus.Infof("Layer 1: %+v", layer1["_source"])
	byteSlice, err := json.MarshalIndent(layer1["_source"], "", "    ")
	if err != nil {
		logrus.Errorf("cannot marshal map[string]interface{}, error: %+v", err)
		return pb.GetArticlesResponse{}
	}

	if err := json.Unmarshal(byteSlice, &instance); err != nil {
		logrus.Errorf("cannot unmarshal byte slice, error: %+v", err)
		return pb.GetArticlesResponse{}
	}
	logrus.Infof("byteslice: %s", string(byteSlice))

	qRead := false
	if instance.QuickRead == "true" {
		qRead = true
	}
	return pb.GetArticlesResponse{
		Id:     instance.ID,
		Title:  instance.Title,
		Author: instance.Author,
		// CreateTime: instance.CreateTime,
		QuickRead: qRead,
		// ViewBy:    instance.ViewedBy,
	}
}
