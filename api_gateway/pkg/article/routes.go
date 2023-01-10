package article

import (
	"context"
	"net/http"
	"time"

	"github.com/89minutes/the_new_project/api_gateway/config"
	"github.com/89minutes/the_new_project/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/article_and_post/pkg/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RegisterArticleRoutes(r *gin.Engine, cfg *config.Config, authClient *auth.ServiceClient) {
	mware := auth.InitAuthMiddleware(authClient)

	svc := &ArticleServiceClient{
		Client: InitArticleServiceClient(cfg),
	}

	routes := r.Group("/api/v1/article")
	routes.Use(mware.AuthRequired)
	routes.POST("/", svc.CreateArticle)

}

func (svc *ArticleServiceClient) CreateArticle(ctx *gin.Context) {
	body := CreateArticleRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		logrus.Info("incomplete body, error", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logrus.Infof("received a post by: %s", body.Author)

	isShort := len(body.Content) <= 1000
	res, err := svc.Client.CreateArticle(context.Background(), &pb.CreateArticleRequest{
		Title:      body.Title,
		Content:    body.Content,
		Author:     body.Author,
		IsDraft:    body.IsDraft,
		Tags:       body.Tags,
		CreateTime: timestamppb.New(time.Now()),
		UpdateTime: timestamppb.New(time.Now()),
		QuickRead:  isShort,
	})

	if err != nil {
		logrus.Info("cannot connect to article rpc server, error", err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
