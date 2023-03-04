package auth

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/89minutes/the_new_project/common"
	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/errors"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
	Log    logrus.Logger
}

func InitServiceClient(cfg *config.Address) pb.AuthServiceClient {
	cc, err := grpc.Dial(cfg.AuthService, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc auth server: %v", err)
	}

	logrus.Infof("The Gateway is dialing to auth gRPC server at: %v", cfg.AuthService)
	return pb.NewAuthServiceClient(cc)
}

func RegisterRouter(router *gin.Engine, cfg *config.Address) *ServiceClient {

	asc := &ServiceClient{
		Client: InitServiceClient(cfg),
		Log:    *logrus.New(),
	}
	routes := router.Group("/api/v1/auth")

	routes.POST("/register", asc.Register)
	routes.POST("/login", asc.Login)

	// Forgot password
	routes.POST("/forgot-pass", asc.ForgotPassword)
	routes.POST("/reset-password", asc.ResetPassword)

	routes.POST("/verify-email", asc.VerifyEmail)

	mware := InitAuthMiddleware(asc)
	routes.Use(mware.AuthRequired)
	routes.POST("/update-password", asc.UpdatePassword)
	routes.POST("/req-email-verification", asc.ReqEmailVerification)

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

	if res.Status == http.StatusNotFound {
		asc.Log.Errorf("user containing email: %s, doesn't exists", body.Email)
		_ = ctx.AbortWithError(http.StatusNotFound, common.NotFound)
		return
	}

	if res.Status == http.StatusBadRequest {
		asc.Log.Errorf("incorrect password given for the user containing email: %s", body.Email)
		_ = ctx.AbortWithError(http.StatusNotFound, common.BadRequest)
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
		errors.RestError(ctx, err, "user")
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// TODO: Rename it to Password Reset Email Verification
func (asc *ServiceClient) ResetPassword(ctx *gin.Context) {
	userAny := ctx.Query("user")
	secretAny := ctx.Query("evpw")

	userId, err := strconv.ParseInt(userAny, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.ResetPassword(context.Background(), &pb.ResetPasswordReq{
		Id:    userId,
		Token: secretAny,
	})

	if err != nil {
		asc.Log.Errorf("rpc auth server returned error: %v", err)
		_ = ctx.AbortWithError(http.StatusForbidden, err)
		return
	}

	if res.Status == http.StatusNotFound || res.Error == "user doesn't exists" {
		asc.Log.Infof("user containing email: %s, doesn't exists", userAny)
		_ = ctx.AbortWithError(http.StatusNotFound, common.NotFound)
		return
	}

	if res.Status == http.StatusBadRequest || res.Error == "incorrect password" {
		asc.Log.Infof("incorrect password given for the user containing email: %s", userAny)
		_ = ctx.AbortWithError(http.StatusNotFound, common.BadRequest)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (asc *ServiceClient) UpdatePassword(ctx *gin.Context) {

	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := asc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	pass := UpdatePassword{}
	if err := ctx.BindJSON(&pass); err != nil {
		asc.Log.Errorf("json body is not correct, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// logrus.Infof("Password: %v", pass.Password)
	// logrus.Infof("res: %+v", res)
	passResp, err := asc.Client.UpdatePassword(context.Background(), &pb.UpdatePasswordReq{
		Password: pass.Password,
		Email:    res.User,
	})
	if err != nil {
		errors.RestError(ctx, err, "user")
		return
	}

	ctx.JSON(http.StatusOK, passResp)
}

// To verify email
func (asc *ServiceClient) VerifyEmail(ctx *gin.Context) {
	userAny := ctx.Query("user")
	secretAny := ctx.Query("evpw")

	res, err := asc.Client.VerifyEmail(context.Background(), &pb.VerifyEmailReq{
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
		_ = ctx.AbortWithError(http.StatusNotFound, common.NotFound)
		return
	}

	if res.Status == http.StatusBadRequest || res.Error == "incorrect password" {
		asc.Log.Infof("incorrect password given for the user containing email: %s", userAny)
		_ = ctx.AbortWithError(http.StatusNotFound, common.BadRequest)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (asc *ServiceClient) ReqEmailVerification(ctx *gin.Context) {
	var vrEmail VerifyEmail

	if err := ctx.BindJSON(&vrEmail); err != nil {
		asc.Log.Errorf("json body is not correct, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := asc.Client.RequestForEmailVerification(context.Background(), &pb.EmailVerificationReq{
		Email: vrEmail.Email,
	})

	if err != nil {
		asc.Log.Errorf("rpc auth server returned error: %v", err)
		_ = ctx.AbortWithError(http.StatusForbidden, err)
		return
	}

	if res.Status == http.StatusNotFound || res.Error == "user doesn't exists" {
		asc.Log.Infof("user containing email: %s, doesn't exists", vrEmail.Email)
		_ = ctx.AbortWithError(http.StatusNotFound, common.NotFound)
		return
	}

	if res.Status == http.StatusBadRequest || res.Error == "incorrect password" {
		asc.Log.Infof("incorrect password given for the user containing email: %s", vrEmail.Email)
		_ = ctx.AbortWithError(http.StatusNotFound, common.BadRequest)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
