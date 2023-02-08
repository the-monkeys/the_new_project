package user_service

import (
	"context"
	"log"
	"net/http"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/user_service/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		s, ok := status.FromError(err)
		if !ok {
			log.Printf("Unexpected error from gRPC server: %v", err)
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		switch s.Code() {
		case codes.NotFound:
			log.Printf("Error from gRPC server: %s", http.StatusText(http.StatusNotFound))
			_ = ctx.AbortWithError(http.StatusNotFound, err)
			return
		case codes.InvalidArgument:
			log.Printf("Error from gRPC server: %s", http.StatusText(http.StatusBadRequest))
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			log.Printf("Error from gRPC server: %s", http.StatusText(http.StatusInternalServerError))
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	ctx.JSON(http.StatusAccepted, &res)
}
