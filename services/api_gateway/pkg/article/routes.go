package article

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/article/pb"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServiceClient struct {
	Client pb.ArticleServiceClient
}

func NewUserServiceClient(cfg *config.Config) pb.ArticleServiceClient {
	logrus.Infof("Dialing to grpc user service: %v", cfg.ArticleSvcUrl)
	cc, err := grpc.Dial(cfg.ArticleSvcUrl, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc user server: %v", err)
	}

	return pb.NewArticleServiceClient(cc)
}

func RegisterArticleRoutes(r *gin.Engine, cfg *config.Config, authClient *auth.ServiceClient) {
	mware := auth.InitAuthMiddleware(authClient)

	svc := &ArticleServiceClient{
		Client: NewUserServiceClient(cfg),
	}

	routes := r.Group("/api/v1/article")
	routes.GET("/", svc.GetArticles)
	routes.GET("/:id", svc.GetArticleById)

	routes.Use(mware.AuthRequired)

	routes.POST("/", svc.CreateArticle)
	routes.PUT("/post/:id/", svc.EditArticles)
	routes.PATCH("/post/:id/", svc.EditArticles)
	// routes.DELETE("/:id", svc.DeleteArticleById)

}

// TODO: Check for all the errors being returned by the gRPC servers...

func (svc *ArticleServiceClient) CreateArticle(ctx *gin.Context) {
	body := CreateArticleRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		logrus.Errorf("incomplete body, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logrus.Infof("received a post by: %s", body.Author)

	isShort := len(body.Content) <= 1000
	res, err := svc.Client.CreateArticle(context.Background(), &pb.CreateArticleRequest{
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
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func (svc *ArticleServiceClient) GetArticles(ctx *gin.Context) {
	logrus.Infof("the page is visited from ip: %s", "192.168.0.3")

	stream, err := svc.Client.GetArticles(context.Background(), &pb.GetArticlesRequest{})
	if err != nil {
		logrus.Errorf("cannot connect to article stream rpc server, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	response := []*pb.GetArticlesResponse{}
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

func (svc *ArticleServiceClient) GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := svc.Client.GetArticleById(context.Background(), &pb.GetArticleByIdReq{Id: id})
	if err != nil {
		logrus.Errorf("cannot connect to article rpc server, error: %v", err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (svc *ArticleServiceClient) EditArticles(ctx *gin.Context) {
	id := ctx.Param("id")

	reqObj := EditArticleRequestBody{}

	if err := ctx.BindJSON(&reqObj); err != nil {
		logrus.Errorf("invalid body, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := svc.Client.EditArticle(context.Background(), &pb.EditArticleReq{
		Id:      id,
		Title:   reqObj.Title,
		Content: reqObj.Content,
		Tags:    reqObj.Tags,
		Method:  ctx.Request.Method,
	})

	if err != nil {
		logrus.Errorf("cannot connect to article rpc server, error: %v", err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (svc *ArticleServiceClient) DeleteArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	logrus.Infof("id : %+v", id)
	res, err := svc.Client.DeleteArticleById(context.Background(), &pb.GetArticleByIdReq{Id: id})

	logrus.Infof("err : %+v", err)
	logrus.Infof("res : %+v", res)
	if err != nil {
		logrus.Errorf("cannot delete the article rpc server sent, error: %v", err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
