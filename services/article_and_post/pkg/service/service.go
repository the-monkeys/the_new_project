package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/89minutes/the_new_project/services/article_and_post/pkg/database"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/models"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/pb"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/utils"
	"github.com/google/uuid"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleServer struct {
	Log      logrus.Logger
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
		Log:      *logrus.New(),
	}, nil
}

func (srv *ArticleServer) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	var article models.Article

	if req.GetId() == "" {
		req.Id = uuid.New().String()
	}

	// Lower cased tags
	for i, v := range req.Tags {
		req.Tags[i] = strings.ToLower(strings.TrimSpace(v))
	}
	req.Title = strings.TrimSpace(req.Title)
	req.Author = strings.TrimSpace(req.Author)
	req.Content = strings.TrimSpace(req.Content)

	req.CanEdit = true
	req.ContentOwnerShip = pb.CreateArticleRequest_THE_USER

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
	bs, err := json.Marshal(post)
	if err != nil {
		logrus.Errorf("cannot marshal the request, error: %v", err)
		return nil, err
	}

	document := strings.NewReader(string(bs))

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

func (srv *ArticleServer) GetArticles(req *pb.GetArticlesRequest, stream pb.ArticleService_GetArticlesServer) error {
	// Search for the document.
	content := strings.NewReader(`
	{
		"size": 100,
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
				"quick_read"
			]
		}
	}`)

	search := opensearchapi.SearchRequest{
		Index: []string{utils.OpensearchArticleIndex},
		Body:  content,
	}

	searchResponse, err := search.Do(context.Background(), srv.osClient)
	if err != nil {
		srv.Log.Errorf("failed to search document, error: %+v", err)
		return status.Errorf(codes.Internal, "failed to find the document, error: %v", err)
	}

	var result map[string]interface{}

	decoder := json.NewDecoder(searchResponse.Body)
	if err := decoder.Decode(&result); err != nil {
		srv.Log.Error("Error while decoding, error", err)
	}

	bx, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		srv.Log.Errorf("cannot marshal map[string]interface{}, error: %+v", err)
		return err
	}

	arts := models.ArticlesForTheMainPage{}
	if err := json.Unmarshal(bx, &arts); err != nil {
		srv.Log.Errorf("cannot unmarshal byte slice, error: %+v", err)
		return err
	}

	articles := ParseToStruct(arts)
	for _, article := range articles {
		if err := stream.Send(&article); err != nil {
			srv.Log.Errorf("error while sending stream, error %+v", err)
		}
	}

	return nil
}

func ParseToStruct(result models.ArticlesForTheMainPage) []pb.GetArticlesResponse {
	var resp []pb.GetArticlesResponse

	for _, val := range result.Hits.Hits {
		t, err := time.Parse("2006-01-02T15:04:05Z", val.Source.CreateTime)
		if err != nil {
			logrus.Errorf("cannot parse the time, error: %v", err)
		}

		res := pb.GetArticlesResponse{
			Id:         val.Source.ID,
			Title:      val.Source.Title,
			Author:     val.Source.Author,
			CreateTime: timestamppb.New(t),
			QuickRead:  val.Source.QuickRead,
			// ViewBy:    instance.ViewedBy,
		}
		resp = append(resp, res)
	}

	return resp
}

func (srv *ArticleServer) GetArticleById(ctx context.Context, req *pb.GetArticleByIdReq) (*pb.GetArticleByIdResp, error) {
	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"id": "%s"
			}
		}
	}`, req.GetId())
	content := strings.NewReader(query)

	search := opensearchapi.SearchRequest{
		Index: []string{utils.OpensearchArticleIndex},
		Body:  content,
	}

	searchResponse, err := search.Do(ctx, srv.osClient)
	if err != nil {
		logrus.Errorf("failed to find document, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "failed to find the document, error: %v", err)
	}

	var result map[string]interface{}

	decoder := json.NewDecoder(searchResponse.Body)
	if err := decoder.Decode(&result); err != nil {
		logrus.Error("error while decoding result, error", err)
		return nil, status.Errorf(codes.Internal, "cannot decode opensearch response: %v", err)
	}

	bx, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		logrus.Errorf("cannot marshal map[string]interface{}, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "cannot marshal opensearch response: %v", err)
	}

	art := models.GetArticleById{}
	if err := json.Unmarshal(bx, &art); err != nil {
		logrus.Errorf("cannot unmarshal byte slice, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "cannot unmarshal opensearch response: %v", err)
	}

	// tStamp, err := SplitSecondsAndNanos(art.Hits.Hits[0].Source.CreateTime)
	if err != nil {
		logrus.Errorf("cannot parse string timestamp to timestamp, error %v", err)
	}
	t, err := time.Parse("2006-01-02T15:04:05Z", art.Hits.Hits[0].Source.CreateTime)
	if err != nil {
		logrus.Errorf("cannot parse the time, error: %v", err)
	}
	return &pb.GetArticleByIdResp{
		Id:         art.Hits.Hits[0].Source.ID,
		Title:      art.Hits.Hits[0].Source.Title,
		Author:     art.Hits.Hits[0].Source.Author,
		Content:    art.Hits.Hits[0].Source.Content,
		CreateTime: timestamppb.New(t),
	}, nil
}

func (srv *ArticleServer) EditArticle(ctx context.Context, req *pb.EditArticleReq) (*pb.EditArticleRes, error) {
	// Get the document from opensearch
	existingArticle, err := srv.GetArticleById(ctx, &pb.GetArticleByIdReq{Id: req.GetId()})
	if err != nil {
		logrus.Errorf("cannot get the existing article, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "cannot get the existing article, error: %v", err)
	}

	// Check if partial then fill a new struct
	artToBeUpdated := PartialOrAllUpdate(req.GetMethod(), existingArticle, req)
	logrus.Infof("Article to be updated: %+v", artToBeUpdated)

	// Update query
	updateQuery := fmt.Sprintf(`
	
	{
		"query": {
			"match": {
				"id": "%s"
			}
		},
		"script": {
			"source": "ctx._source.title = params.title; ctx._source.content = params.content",
			"lang": "painless",
			"params": {
				"title": "%s",
				"content": "%s"
			}
		}
	}`, artToBeUpdated.GetId(), artToBeUpdated.Title, artToBeUpdated.Content)
	document := strings.NewReader(updateQuery)

	updateReq := opensearchapi.UpdateByQueryRequest{
		Index: []string{utils.OpensearchArticleIndex},
		Body:  document,
	}

	fmt.Println(updateQuery)
	updateRes, err := updateReq.Do(ctx, srv.osClient)
	if err != nil {
		logrus.Errorf("failed to update the document, error: %+v", err)
		return nil, status.Errorf(codes.Internal, "cannot update the document, error: %v", err)
	}
	logrus.Infof("Updateed: %+v", updateRes)
	if updateRes.IsError() {
		logrus.Errorf("failed to update the document, bad request, error: %+v", err)
		return nil, status.Errorf(codes.InvalidArgument, "cannot update the document, error: %v", err)
	}

	return &pb.EditArticleRes{
		Status: http.StatusCreated,
		Id:     artToBeUpdated.Id,
	}, nil
}

func PartialOrAllUpdate(method string, existingArt *pb.GetArticleByIdResp, reqArt *pb.EditArticleReq) *pb.EditArticleReq {
	procdArt := &pb.EditArticleReq{Id: reqArt.Id}

	if method == http.MethodPatch {
		if reqArt.Title == "" {
			procdArt.Title = existingArt.Title
		} else {
			procdArt.Title = reqArt.Title
		}
		if reqArt.Content == "" {
			procdArt.Content = existingArt.Content
		} else {
			procdArt.Content = reqArt.Content
		}
	} else {
		procdArt.Title = reqArt.Title
		procdArt.Content = reqArt.Content
	}

	return procdArt
}
