package webapi

import (
	"github.com/dkzhang/RmsGo/webapi/handle/handleGeneralFormDraft"
	"github.com/dkzhang/RmsGo/webapi/handle/handleLogin"
	"github.com/dkzhang/RmsGo/webapi/handle/handleUser"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func Serve() {
	infra := infrastructure.NewInfrastructure(infrastructure.InfraConfigFile{
		LogMapConf: os.Getenv("LogMapConf"),
		SmsSE:      os.Getenv("SmsSE"),
		DbSE:       os.Getenv("DbSE"),
		LoginConf:  os.Getenv("LoginConf"),
	})

	r := gin.Default()
	r.Use(middleware.LoggerGinToFile(infra.TheLogMap))

	/////////////////////////////////////////////////////////////

	webAPIv1 := r.Group("/webapi")
	{
		webAPIv1.POST("/ApplyLogin", func(c *gin.Context) { handleLogin.ApplyLogin(infra, c) })
		webAPIv1.POST("/Login", func(c *gin.Context) { handleLogin.Login(infra, c) })
		webAPIv1.GET("/Logout", middleware.TokenAuth(infra), func(c *gin.Context) { handleLogin.Logout(infra, c) })

		webAPIv1.GET("/AllUsers", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.AllUsers(infra, c) })

		hUser := webAPIv1.Group("/User")
		{
			hUser.GET("/", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.RetrieveUserLogin(infra, c) })

			hUser.POST("/", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.Create(infra, c) })
			hUser.GET("/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.Retrieve(infra, c) })
			hUser.PUT("/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.Update(infra, c) })
			hUser.DELETE("/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleUser.Delete(infra, c) })
		}

		hAppTemp := webAPIv1.Group("/GeneralFormDraft")
		{
			hAppTemp.GET("/", middleware.TokenAuth(infra), func(c *gin.Context) { handleGeneralFormDraft.RetrieveByOwner(infra, c) })

			hAppTemp.POST("/", middleware.TokenAuth(infra), func(c *gin.Context) { handleGeneralFormDraft.Create(infra, c) })
			hAppTemp.GET("/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleGeneralFormDraft.RetrieveByID(infra, c) })
			hAppTemp.PUT("/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleGeneralFormDraft.Update(infra, c) })
			hAppTemp.DELETE("/:id", middleware.TokenAuth(infra), func(c *gin.Context) { handleGeneralFormDraft.Delete(infra, c) })
		}
	}

	/////////////////////////////////////////////////////////////
	r.Run(":8083")
}
