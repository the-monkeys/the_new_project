package user_service

import (
	"context"
	"net/http"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/user_service/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func NewUserServiceClient(cfg *config.Config) pb.UserServiceClient {
	cc, err := grpc.Dial(cfg.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc user server: %v", err)
	}
	logrus.Infof("The Gateway is dialing to user gRPC server at: %v", cfg.UserSvcUrl)
	return pb.NewUserServiceClient(cc)
}

func RegisterUserRouter(router *gin.Engine, cfg *config.Config, authClient *auth.ServiceClient) *UserServiceClient {
	mware := auth.InitAuthMiddleware(authClient)

	usc := &UserServiceClient{
		Client: NewUserServiceClient(cfg),
	}
	routes := router.Group("/api/v1/profile")
	routes.Use(mware.AuthRequired)
	routes.GET("/user", usc.GetProfile)

	return usc
}

func (asc *UserServiceClient) GetProfile(ctx *gin.Context) {

	body := ProfileRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.GetMyProfile(context.Background(), &pb.GetMyProfileReq{
		Id: body.Id,
	})

	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusAccepted, &res)

}
