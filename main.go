package main

import (
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/login"
	"github.com/dkzhang/RmsGo/webapi/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	webapi.InitInfrastructure()

	r := gin.Default()
	r.Use(middleware.LoggerGinToFile())

	/////////////////////////////////////////////////////////////

	r.POST("/ApplyLogin", login.ApplyLogin)
	r.POST("/Login", login.Login)

	/////////////////////////////////////////////////////////////
	r.Run()
}
