package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/89minutes/the_new_project/common"
	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/models"
	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BlogService struct {
	osClient openSearchClient
	logger   *logrus.Logger
	pb.UnimplementedBlogsAndPostServiceServer
}

func NewBlogService(client openSearchClient,
	logger *logrus.Logger) *BlogService {
	return &BlogService{osClient: client, logger: logger}
}

func (blog *BlogService) CreateABlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog.logger.Infof("received a create blog request from user: %v", req.AuthorId)

	// Lower cased tags and trim spaces
	for i, v := range req.Tags {
		req.Tags[i] = strings.ToLower(strings.TrimSpace(v))
	}

	// Trim spaces from fields
	req.Title = strings.TrimSpace(req.Title)
	req.AuthorName = strings.TrimSpace(req.AuthorName)
	req.Content = strings.TrimSpace(req.Content)
	req.AuthorId = strings.TrimSpace(req.AuthorId)

	req.CanEdit = true
	req.Ownership = pb.CreateBlogRequest_THE_USER

	// Assign to models struct
	post := models.Blogs{
		Id:          req.Id,
		Title:       req.Title,
		Content:     req.Content,
		Author:      req.AuthorName,
		AuthorEmail: req.AuthorId,
		Published:   &req.Published,
		Tags:        req.Tags,
		CreateTime:  req.CreateTime.AsTime().Format("2006-01-02T15:04:05Z07:00"),
		UpdateTime:  req.UpdateTime.AsTime().Format("2006-01-02T15:04:05Z07:00"),
		CanEdit:     &req.CanEdit,
		OwnerShip:   req.Ownership,
		FolderPath:  "",
	}

	// Create the articles
	resp, err := blog.osClient.CreateAnArticle(post)
	if err != nil {
		blog.logger.Errorf("cannot save the blog, error: %+v", err)
		return nil, err
	}

	if resp.StatusCode == http.StatusBadRequest {
		blog.logger.Errorf("cannot save the blog bad request, error: %+v", err)
		return nil, common.BadRequest
	}

	blog.logger.Infof("user %v created a blog successfully: %v", req.GetAuthorId(), req.GetId())

	return &pb.CreateBlogResponse{
		Status: int64(resp.StatusCode),
		Id:     int64(resp.StatusCode),
	}, nil
}

func (blog *BlogService) Get100Blogs(req *emptypb.Empty, stream pb.BlogsAndPostService_Get100BlogsServer) error {
	searchResponse, err := blog.osClient.GetLast100Articles()
	var result map[string]interface{}

	// logrus.Infof("Response: %+v", searchResponse)
	decoder := json.NewDecoder(searchResponse.Body)
	if err := decoder.Decode(&result); err != nil {
		blog.logger.Error("error while decoding, error", err)
	}

	bx, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		blog.logger.Errorf("cannot marshal map[string]interface{}, error: %+v", err)
		return err
	}

	arts := models.Last100Articles{}
	if err := json.Unmarshal(bx, &arts); err != nil {
		blog.logger.Errorf("cannot unmarshal byte slice, error: %+v", err)
		return err
	}

	articles := parseToStruct(arts)
	for _, article := range articles {
		logrus.Printf("Article: %+v\n", article)
		if err := stream.Send(&article); err != nil {
			blog.logger.Errorf("error while sending stream, error %+v", err)
		}
	}

	return nil
}
