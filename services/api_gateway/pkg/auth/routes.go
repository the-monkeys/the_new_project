package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
	Log    logrus.Logger
}

func InitServiceClient(cfg *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(cfg.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc auth server: %v", err)
	}

	logrus.Infof("gateway is dialing to the auth server at: %v", cfg.AuthSvcUrl)
	return pb.NewAuthServiceClient(cc)
}

func RegisterRouter(router *gin.Engine, cfg *config.Config) *ServiceClient {
	asc := &ServiceClient{
		Client: InitServiceClient(cfg),
		Log:    *logrus.New(),
	}
	routes := router.Group("/api/v1/auth")
	routes.POST("/register", asc.Register)
	routes.POST("/login", asc.Login)
	routes.POST("/forgot_pass", asc.ForgotPassword)
	routes.GET("/resetpassword", asc.ResetPassword)

	return asc
}

func (asc *ServiceClient) Register(ctx *gin.Context) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		asc.Log.Errorf("json body is not correct, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logrus.Infof("traffic is coming from ip: %v", ctx.ClientIP())

	body.FirstName = strings.TrimSpace(body.FirstName)
	body.LastName = strings.TrimSpace(body.LastName)
	body.Email = strings.TrimSpace(body.Email)

	res, err := asc.Client.Register(context.Background(), &pb.RegisterRequest{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	})

	logrus.Infof("Response: %+v", res)
	logrus.Infof("Error: %+v", err)

	if err != nil {
		asc.Log.Errorf("rpc auth server returned error, error: %v", err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if res.Status == http.StatusConflict {
		ctx.JSON(http.StatusConflict, nil)
		return
	}

	ctx.JSON(int(res.Status), &res)

}

func (asc *ServiceClient) Login(ctx *gin.Context) {
	body := LoginRequestBody{}

	logrus.Infof("traffic is coming from ip: %v", ctx.ClientIP())

	if err := ctx.BindJSON(&body); err != nil {
		asc.Log.Errorf("json body is not correct, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// do the trimming
	body.Email = strings.TrimSpace(body.Email)

	res, err := asc.Client.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		asc.Log.Errorf("internal server error, user containing email: %s cannot login", body.Email)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if res.Status == http.StatusNotFound || res.Error == "user doesn't exists" {
		asc.Log.Errorf("user containing email: %s, doesn't exists", body.Email)
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(res.Error))
		return
	}

	if res.Status == http.StatusBadRequest || res.Error == "incorrect password" {
		asc.Log.Errorf("incorrect password given for the user containing email: %s", body.Email)
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(res.Error))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (asc *ServiceClient) ForgotPassword(ctx *gin.Context) {
	body := ForgetPass{}

	body.Email = strings.TrimSpace(body.Email)

	if err := ctx.BindJSON(&body); err != nil {
		asc.Log.Errorf("json body is not correct, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.ForgotPassword(context.Background(), &pb.ForgotPasswordReq{
		Email: body.Email,
	})

	if err != nil {
		asc.Log.Errorf("internal server error, user containing email: %s cannot login", body.Email)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if res.Status == http.StatusNotFound || res.Error == "user doesn't exists" {
		asc.Log.Infof("user containing email: %s, doesn't exists", body.Email)
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(res.Error))
		return
	}

	if res.Status == http.StatusBadRequest || res.Error == "incorrect password" {
		asc.Log.Infof("incorrect password given for the user containing email: %s", body.Email)
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(res.Error))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (asc *ServiceClient) ResetPassword(ctx *gin.Context) {
	userAny := ctx.Query("user")
	secretAny := ctx.Query("evpw")

	res, err := asc.Client.ResetPassword(context.Background(), &pb.ResetPasswordReq{
		Email: userAny,
		Token: secretAny,
	})

	if err != nil {
		asc.Log.Errorf("rpc auth server returned error: %v", err)
		_ = ctx.AbortWithError(http.StatusForbidden, err)
		return
	}

	if res.Status == http.StatusNotFound || res.Error == "user doesn't exists" {
		asc.Log.Infof("user containing email: %s, doesn't exists", userAny)
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(res.Error))
		return
	}

	if res.Status == http.StatusBadRequest || res.Error == "incorrect password" {
		asc.Log.Infof("incorrect password given for the user containing email: %s", userAny)
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(res.Error))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
