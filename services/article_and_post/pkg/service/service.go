package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/89minutes/the_new_project/services/article_and_post/pkg/models"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/pb"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/utils"
	"github.com/google/uuid"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleServer struct {
	Log      *logrus.Logger
	osClient *openSearchClient
	pb.UnimplementedArticleServiceServer
}

func NewArticleServer(url, username, password string, logrus *logrus.Logger) (*ArticleServer, error) {
	client, err := newOpenSearchClient(url, username, password, logrus)
	if err != nil {
		logrus.Errorf("Failed to connect to opensearch instance, error: %+v", err)
		return nil, err
	}

	return &ArticleServer{
		osClient: client,
		Log:      logrus,
	}, nil
}

func (srv *ArticleServer) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	var article models.Article

	if req.GetId() == "" {
		req.Id = uuid.New().String()
	}

	// Lower cased tags and trim spaces
	for i, v := range req.Tags {
		req.Tags[i] = strings.ToLower(strings.TrimSpace(v))
	}

	// Trim spaces from fields
	req.Title = strings.TrimSpace(req.Title)
	req.Author = strings.TrimSpace(req.Author)
	req.Content = strings.TrimSpace(req.Content)
	req.AuthorEmail = strings.TrimSpace(req.AuthorEmail)

	req.CanEdit = true
	req.ContentOwnerShip = pb.CreateArticleRequest_THE_USER

	// Assign to models struct
	post := models.Article{
		Id:          req.Id,
		Title:       req.Title,
		Content:     req.Content,
		Author:      req.Author,
		AuthorEmail: req.AuthorEmail,
		IsDraft:     &req.IsDraft,
		Tags:        req.Tags,
		CreateTime:  req.CreateTime.AsTime().Format("2006-01-02T15:04:05Z07:00"),
		UpdateTime:  req.UpdateTime.AsTime().Format("2006-01-02T15:04:05Z07:00"),
		QuickRead:   req.QuickRead,
		CanEdit:     &req.CanEdit,
		OwnerShip:   pb.CreateArticleRequest_THE_USER,
		FolderPath:  "",
	}

	// Create the articles
	resp, err := srv.osClient.CreateAnArticle(post)
	if err != nil {
		srv.Log.Errorf("cannot save the post, error: %+v", err)
	}

	srv.Log.Infof("The status code for the save post is: %v", resp.StatusCode)

	return &pb.CreateArticleResponse{
		Status: int64(resp.StatusCode),
		Id:     article.Id,
	}, nil
}

//
func (srv *ArticleServer) GetArticles(req *pb.GetArticlesRequest, stream pb.ArticleService_GetArticlesServer) error {
	searchResponse, err := srv.osClient.GetLast100Articles()
	if err != nil {
		srv.Log.Error("error while getting 100 articles, error", err)
		return err
	}

	var result map[string]interface{}

	decoder := json.NewDecoder(searchResponse.Body)
	if err := decoder.Decode(&result); err != nil {
		srv.Log.Error("error while decoding, error", err)
		return err
	}

	bx, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		srv.Log.Errorf("cannot marshal map[string]interface{}, error: %+v", err)
		return err
	}

	arts := models.Last100Articles{}
	if err := json.Unmarshal(bx, &arts); err != nil {
		srv.Log.Errorf("cannot unmarshal byte slice, error: %+v", err)
		return err
	}

	articles := ParseToStruct(arts)
	for _, art := range articles {
		if err := stream.Send(&art); err != nil {
			srv.Log.Errorf("error while sending stream, error %+v", err)
		}
	}

	return nil
}

func (srv *ArticleServer) GetArticleById(ctx context.Context, req *pb.GetArticleByIdReq) (*pb.GetArticleByIdResp, error) {

	searchResponse, err := srv.osClient.GetArticleById(ctx, req.GetId())
	if err != nil {
		srv.Log.Errorf("failed to find document, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "failed to find the document, error: %v", err)
	}

	var result map[string]interface{}

	// logrus.Infof("Response: %+v", searchResponse)

	decoder := json.NewDecoder(searchResponse.Body)
	if err := decoder.Decode(&result); err != nil {
		srv.Log.Error("error while decoding result, error", err)
		return nil, status.Errorf(codes.Internal, "cannot decode opensearch response: %v", err)
	}

	bx, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		srv.Log.Errorf("cannot marshal map[string]interface{}, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "cannot marshal opensearch response: %v", err)
	}

	art := models.GetArticleById{}
	if err := json.Unmarshal(bx, &art); err != nil {
		srv.Log.Errorf("cannot unmarshal byte slice, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "cannot unmarshal opensearch response: %v", err)
	}

	if err != nil {
		srv.Log.Errorf("cannot parse string timestamp to timestamp, error %v", err)
	}

	return &pb.GetArticleByIdResp{
		Id:         art.Hits.Hits[0].Source.ID,
		Title:      art.Hits.Hits[0].Source.Title,
		Author:     art.Hits.Hits[0].Source.Author,
		Content:    art.Hits.Hits[0].Source.Content,
		CreateTime: timestamppb.New(art.Hits.Hits[0].Source.CreateTime),
		Tags:       art.Hits.Hits[0].Source.Tags,
	}, nil
}

func (srv *ArticleServer) EditArticle(ctx context.Context, req *pb.EditArticleReq) (*pb.EditArticleRes, error) {
	// Lower cased tags and trim spaces
	for i, v := range req.Tags {
		req.Tags[i] = strings.ToLower(strings.TrimSpace(v))
	}

	// Trim spaces from fields
	req.Title = strings.TrimSpace(req.Title)
	req.Content = strings.TrimSpace(req.Content)

	// Get the document from opensearch
	existingArticle, err := srv.GetArticleById(ctx, &pb.GetArticleByIdReq{Id: req.GetId()})
	if err != nil {
		srv.Log.Errorf("cannot get the existing article, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "cannot get the existing article, error: %v", err)
	}

	// Check if partial then fill a new struct
	toBeUpdated := partialOrAllUpdate(req.GetMethod(), existingArticle, req)
	logrus.Infof("Article to be updated: %+v", toBeUpdated.Id)

	document := strings.NewReader(updateArticleById(toBeUpdated.Id, toBeUpdated.Title, toBeUpdated.Content, toBeUpdated.Tags))

	updateReq := opensearchapi.UpdateByQueryRequest{
		Index: []string{utils.OpensearchArticleIndex},
		Body:  document,
	}

	updateRes, err := updateReq.Do(ctx, srv.osClient.client)
	if err != nil {
		srv.Log.Errorf("failed to update the document, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "cannot update the document, error: %v", err)
	}

	if updateRes.IsError() {
		srv.Log.Errorf("cannot update the document, error: %+v", updateRes)
		return nil, status.Errorf(codes.Internal, "cannot update the document, error: %v", err)
	}

	if updateRes.StatusCode == http.StatusBadRequest {
		srv.Log.Errorf("cannot update the document, bad request, error: %+v", updateRes)
		return nil, status.Errorf(codes.Internal, "cannot update the document, error: %v", err)
	}

	logrus.Infof("Updated the article %s", req.Id)

	if updateRes.IsError() {
		srv.Log.Errorf("failed to update the document, bad request, error: %+v", err)
		return nil, status.Errorf(codes.InvalidArgument, "cannot update the document, error: %v", err)
	}

	return &pb.EditArticleRes{
		Status: http.StatusCreated,
		Id:     toBeUpdated.Id,
	}, nil
}

func (srv *ArticleServer) DeleteArticleById(ctx context.Context, req *pb.GetArticleByIdReq) (*pb.DeleteArticleByIdRes, error) {

	deleteResp, err := srv.osClient.DeleteArticleById(ctx, req.GetId())
	if deleteResp.StatusCode == http.StatusNotFound {
		srv.Log.Errorf("cannot find the article %s, error: %+v", req.GetId(), err)
		return &pb.DeleteArticleByIdRes{
			Status: int64(http.StatusNotFound),
		}, errors.New("cannot find the document")
	}

	if err != nil {
		srv.Log.Errorf("cannot delete the article %s, error: %+v", req.GetId(), err)
		return &pb.DeleteArticleByIdRes{
			Status: int64(http.StatusInternalServerError),
		}, errors.New("internal server error")
	}

	logrus.Info(("RPC server cannot take the error"))
	return &pb.DeleteArticleByIdRes{
		Status: int64(deleteResp.StatusCode),
		Id:     req.GetId(),
	}, nil
}
