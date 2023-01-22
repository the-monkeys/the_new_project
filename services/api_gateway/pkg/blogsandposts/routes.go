package blogsandposts

import (
	"context"
	"io"
	"net/http"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/blogsandposts/pb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BlogServiceClient struct {
	Client pb.BlogsAndPostServiceClient
}

func NewUserServiceClient(cfg *config.Config) pb.BlogsAndPostServiceClient {
	cc, err := grpc.Dial(cfg.BlogAndPostSvcURL, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc user server: %v", err)
	}

	return pb.NewBlogsAndPostServiceClient(cc)
}

func RegisterBlogRouter(router *gin.Engine, cfg *config.Config, authClient *auth.ServiceClient) *BlogServiceClient {
	mware := auth.InitAuthMiddleware(authClient)

	blogCli := &BlogServiceClient{
		Client: NewUserServiceClient(cfg),
	}
	routes := router.Group("/api/v1/post")
	routes.GET("/", blogCli.Get100Blogs)
	routes.GET("/:id", blogCli.GetArticleById)

	routes.Use(mware.AuthRequired)

	routes.POST("/create", blogCli.CreateABlog)

	return blogCli
}

func (asc *BlogServiceClient) CreateABlog(ctx *gin.Context) {

	body := CreatePostRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		logrus.Errorf("cannot bind json to struct, error: %v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.CreateABlog(context.Background(), &pb.CreateBlogRequest{
		Id:         uuid.NewString(),
		Title:      body.Title,
		Content:    body.Content,
		AuthorName: body.Author,
		AuthorId:   body.AuthorId,
		Published:  body.Published,
		Tags:       body.Tags,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	logrus.Errorf("Response: %+v", res)
	ctx.JSON(http.StatusAccepted, &res)

}

func (svc *BlogServiceClient) Get100Blogs(ctx *gin.Context) {
	logrus.Infof("the page is visited from ip: %s", "192.168.0.3")

	stream, err := svc.Client.Get100Blogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		logrus.Errorf("cannot connect to article stream rpc server, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
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

		logrus.Infof("Got response: %+v", resp)
		response = append(response, resp)
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *BlogServiceClient) GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := svc.Client.GetBlogById(context.Background(), &pb.GetBlogByIdRequest{Id: id})
	if err != nil {
		logrus.Errorf("cannot connect to article rpc server, error: %v", err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
