package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	// auth.RegisterRouter(router, &cfg)
	logrus.Info("starting the article server at port: ", 50002)
	router.Run(":50002")
}
