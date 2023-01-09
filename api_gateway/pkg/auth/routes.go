package auth

import (
	"context"
	"net/http"

	"github.com/89minutes/the_new_project/api_gateway/config"
	"github.com/89minutes/the_new_project/api_gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type serviceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(cfg.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc auth server: %v", err)
	}

	return pb.NewAuthServiceClient(cc)
}

func RegisterRouter(router *gin.Engine, cfg *config.Config) *serviceClient {
	asc := &serviceClient{
		Client: InitServiceClient(cfg),
	}
	routes := router.Group("/api/v1/auth")
	routes.POST("/register", asc.Register)
	routes.POST("/login", asc.Login)

	return asc
}

func (asc *serviceClient) Register(ctx *gin.Context) {
	// Register(ctx, asc.Client)

	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}

func (asc *serviceClient) Login(ctx *gin.Context) {
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
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

// func Register(ctx *gin.Context, c pb.AuthServiceClient) {
// 	body := RegisterRequestBody{}

// 	if err := ctx.BindJSON(&body); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	res, err := c.Register(context.Background(), &pb.RegisterRequest{
// 		Email:    body.Email,
// 		Password: body.Password,
// 	})

// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadGateway, err)
// 		return
// 	}

// 	ctx.JSON(int(res.Status), &res)
// }

// func Login(ctx *gin.Context, c pb.AuthServiceClient) {
// 	b := LoginRequestBody{}

// 	if err := ctx.BindJSON(&b); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	res, err := c.Login(context.Background(), &pb.LoginRequest{
// 		Email:    b.Email,
// 		Password: b.Password,
// 	})

// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadGateway, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, &res)
// }
