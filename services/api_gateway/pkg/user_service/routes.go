package user_service

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/errors"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/user_service/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func NewUserServiceClient(cfg *config.Address) pb.UserServiceClient {
	cc, err := grpc.Dial(cfg.UserService, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc user server: %v", err)
	}
	logrus.Infof("The Gateway is dialing to user gRPC server at: %v", cfg.UserService)
	return pb.NewUserServiceClient(cc)
}

func RegisterUserRouter(router *gin.Engine, cfg *config.Address, authClient *auth.ServiceClient) *UserServiceClient {
	mware := auth.InitAuthMiddleware(authClient)

	usc := &UserServiceClient{
		Client: NewUserServiceClient(cfg),
	}
	routes := router.Group("/api/v1/profile")
	routes.Use(mware.AuthRequired)
	routes.GET("/user/:id", usc.GetProfile)
	routes.POST("/user/:id", usc.UpdateProfile)
	routes.POST("/user/pic/:id", usc.UpdateProfilePic)
	routes.GET("/user/pic/:id", usc.GetProfilePic)

	return usc
}

func (asc *UserServiceClient) GetProfile(ctx *gin.Context) {
	// get id
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.GetMyProfile(context.Background(), &pb.GetMyProfileReq{
		Id: userId,
	})

	if err != nil {
		errors.RestError(ctx, err, "user")
		return
	}

	ctx.JSON(http.StatusAccepted, &res)
}

func (asc *UserServiceClient) UpdateProfile(ctx *gin.Context) {
	// get id
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body := UpdateProfile{}
	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.SetMyProfile(context.Background(), &pb.SetMyProfileReq{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		CountryCode: body.CountryCode,
		MobileNo:    body.MobileNo,
		About:       body.About,
		Instagram:   body.Instagram,
		Twitter:     body.Twitter,
		Email:       body.Email,
		Id:          userId,
	})

	if err != nil {
		errors.RestError(ctx, err, "user")
		return
	}

	ctx.JSON(http.StatusAccepted, &res)
}

func (asc *UserServiceClient) UpdateProfilePic(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading image data:", err)

	}

	// get id
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stream, err := asc.Client.UploadProfile(context.Background())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	chunk := &pb.UploadProfilePicReq{
		Data: imageData,
		Id:   userId,
	}
	err = stream.Send(chunk)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// log.Printf("%+v\n", response)
	ctx.JSON(http.StatusAccepted, "uploaded")
}

// TODO: Handle 404 error it's throwing error
func (asc *UserServiceClient) GetProfilePic(ctx *gin.Context) {
	// get id
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stream, err := asc.Client.Download(context.Background(), &pb.GetProfilePicReq{Id: userId})
	if err != nil {
		logrus.Errorf("cannot connect to user rpc server, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	resp, err := stream.Recv()
	if err == io.EOF {

	}
	if err != nil {
		errors.RestError(ctx, err, "user_service")
		logrus.Errorf("cannot get the stream data, error: %+v", err)
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename=file-name.txt")
	ctx.Data(http.StatusOK, "application/octet-stream", resp.Data)
}
