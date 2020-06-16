package main

import (
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/handle/handleLogin"
	"github.com/dkzhang/RmsGo/webapi/handle/handleUser"
	"github.com/dkzhang/RmsGo/webapi/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	infra := webapi.NewInfrastructure(webapi.InfraConfigFile{
		LogMapConf: os.Getenv("LogMapConf"),
		SmsSE:      os.Getenv("SmsSE"),
		DbSE:       os.Getenv("DbSE"),
		LoginConf:  os.Getenv("LoginConf"),
	})

	r := gin.Default()
	r.Use(middleware.LoggerGinToFile())

	/////////////////////////////////////////////////////////////

	r.POST("/ApplyLogin", func(c *gin.Context) { handleLogin.ApplyLogin(infra, c) })
	r.POST("/Login", func(c *gin.Context) { handleLogin.Login(infra, c) })

	r.GET("/AllUsers", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.AllUsers(infra, c) })

	r.GET("/User", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.RetrieveUserLogin(infra, c) })

	r.POST("/User", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.Create(infra, c) })
	r.GET("/User/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.Retrieve(infra, c) })
	r.PUT("/User/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.Update(infra, c) })
	r.DELETE("/User/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.Delete(infra, c) })

	/////////////////////////////////////////////////////////////
	r.Run(":8080")
}
