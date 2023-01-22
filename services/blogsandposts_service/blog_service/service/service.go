package service

import (
	"context"
	"strings"

	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/models"
	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/pb"
	"github.com/sirupsen/logrus"
)

type BlogService struct {
	client openSearchClient
	logger *logrus.Logger
	pb.UnimplementedBlogsAndPostServiceServer
}

func NewBlogService(client openSearchClient,
	logger *logrus.Logger) *BlogService {
	return &BlogService{client: client, logger: logger}
}

// func (us *BlogService) CreateABlog(ctx context.Context, req *pb.CreateBlogReq) (*pb.CreateBlogRes, error) {

// 	res := &pb.CreateBlogRes{}

// 	logrus.Infof("fetched profile for the user containing id: %d", res.Id)
// 	return res, nil
// }

func (us *BlogService) CreateABlog(ctx context.Context, req *pb.CreateBlogReq) (*pb.CreateBlogRes, error) {

	var article models.Blogs
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
	req.Ownership = pb.CreateBlogReq_THE_USER

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

	logrus.Infof("The blog: %v", post)
	// Create the articles
	// resp, err := srv.OsClient.CreateAnArticle(post)
	// if err != nil {
	// 	srv.Log.Errorf("cannot save the post, error: %+v", err)
	// }

	// srv.Log.Infof("The status code for the save post is: %v", resp.StatusCode)

	return &pb.CreateBlogRes{
		Message: "Successful",
		Id:      article.Id,
	}, nil
}
