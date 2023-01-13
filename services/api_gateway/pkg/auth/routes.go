package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(cfg.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc auth server: %v", err)
	}

	return pb.NewAuthServiceClient(cc)
}

func RegisterRouter(router *gin.Engine, cfg *config.Config) *ServiceClient {
	asc := &ServiceClient{
		Client: InitServiceClient(cfg),
	}
	routes := router.Group("/api/v1/auth")
	routes.POST("/register", asc.Register)
	routes.POST("/login", asc.Login)

	return asc
}

func (asc *ServiceClient) Register(ctx *gin.Context) {
	// Register(ctx, asc.Client)

	body := RegisterRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.Register(context.Background(), &pb.RegisterRequest{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}

func (asc *ServiceClient) Login(ctx *gin.Context) {
	// Login(ctx, asc.Client)
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		logrus.Errorf("internal server error, user containing email: %s cannot login", b.Email)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if res.Status == http.StatusNotFound || res.Error == "user doesn't exists" {
		logrus.Infof("user containing email: %s, doesn't exists", b.Email)
		ctx.AbortWithError(http.StatusNotFound, errors.New(res.Error))
		return
	}

	if res.Status == http.StatusBadRequest || res.Error == "incorrect password" {
		logrus.Infof("incorrect password given for the user containing email: %s", b.Email)
		ctx.AbortWithError(http.StatusNotFound, errors.New(res.Error))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
