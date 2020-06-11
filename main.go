package main

import (
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/handleLogin"
	"github.com/dkzhang/RmsGo/webapi/handleUser"
	"github.com/dkzhang/RmsGo/webapi/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	webapi.InitInfrastructure()

	r := gin.Default()
	r.Use(middleware.LoggerGinToFile())

	/////////////////////////////////////////////////////////////

	r.POST("/ApplyLogin", handleLogin.ApplyLogin)
	r.POST("/Login", handleLogin.Login)

	r.GET("/AllUsers", middleware.TokenAuth(), handleUser.AllUsers)

	r.PUT("/User", middleware.TokenAuth(), handleUser.Create)
	r.GET("/User/:id", middleware.TokenAuth(), handleUser.Retrieve)
	r.POST("/User/:id", middleware.TokenAuth(), handleUser.Update)
	r.DELETE("/User/:id", middleware.TokenAuth(), handleUser.Delete)

	/////////////////////////////////////////////////////////////
	r.Run()
}
