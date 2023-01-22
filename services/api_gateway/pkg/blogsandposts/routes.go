package blogsandposts

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BlogServiceClient struct {
	Client pb.BlogServiceClient
}

func InitBlogServiceClient(cfg *config.Config) pb.BlogServiceClient {
	// using WithInsecure() because no SSL running
	logrus.Infof("dialing to blog server at: %v", cfg.BlogAndPostSvcURL)
	cc, err := grpc.Dial(cfg.BlogAndPostSvcURL, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil
	}

	return pb.NewBlogServiceClient(cc)
}

func RegisterBlogsAndPostsRoutes(r *gin.Engine, cfg *config.Config, authClient *auth.ServiceClient) {
	mware := auth.InitAuthMiddleware(authClient)

	svc := &BlogServiceClient{
		Client: InitBlogServiceClient(cfg),
	}

	routes := r.Group("/api/v1/post")
	routes.GET("/", svc.GetArticles)
	routes.GET("/:id", svc.GetArticleById)

	routes.Use(mware.AuthRequired)

	routes.POST("/", svc.CreateArticle)
	routes.PUT("/post/:id/", svc.EditArticles)
	routes.PATCH("/post/:id/", svc.EditArticles)

}

// TODO: Check for all the errors being returned by the gRPC servers...

func (svc *BlogServiceClient) CreateArticle(ctx *gin.Context) {
	body := CreateArticleRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		logrus.Errorf("incomplete body, error", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logrus.Infof("received a post by: %s", body.Author)

	isShort := len(body.Content) <= 1000
	res, err := svc.Client.CreateBlog(context.Background(), &pb.CreateBlogRequest{
		Title:       body.Title,
		Content:     body.Content,
		Author:      body.Author,
		IsDraft:     body.IsDraft,
		Tags:        body.Tags,
		CreateTime:  timestamppb.New(time.Now()),
		UpdateTime:  timestamppb.New(time.Now()),
		QuickRead:   isShort,
		AuthorEmail: body.AuthorEmail,
	})

	if err != nil {
		logrus.Infof("cannot connect to article rpc server, error: %v", err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func (svc *BlogServiceClient) GetArticles(ctx *gin.Context) {
	logrus.Infof("the page is visited from ip: %s", "192.168.0.3")

	stream, err := svc.Client.GetBlogs(context.Background(), &pb.GetBlogsRequest{})
	if err != nil {
		logrus.Errorf("cannot connect to article stream rpc server, error: %v", err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	response := []*pb.GetBlogsResponse{}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Errorf("cannot get the stream data, error: %+v", err)
		}

		logrus.Info("Got response: %+v", resp)
		response = append(response, resp)
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *BlogServiceClient) GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := svc.Client.GetBlogById(context.Background(), &pb.GetBlogByIdReq{Id: id})
	if err != nil {
		logrus.Errorf("cannot connect to article rpc server, error: %v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (svc *BlogServiceClient) EditArticles(ctx *gin.Context) {
	id := ctx.Param("id")

	reqObj := EditArticleRequestBody{}

	if err := ctx.BindJSON(&reqObj); err != nil {
		logrus.Errorf("invalid body, error", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := svc.Client.EditBlog(context.Background(), &pb.EditBlogReq{
		Id:      id,
		Title:   reqObj.Title,
		Content: reqObj.Content,
		Tags:    reqObj.Tags,
		Method:  ctx.Request.Method,
	})

	if err != nil {
		logrus.Errorf("cannot connect to article rpc server, error: %v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
