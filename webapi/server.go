package webapi

import (
	"github.com/dkzhang/RmsGo/webapi/handle/handleApplication"
	"github.com/dkzhang/RmsGo/webapi/handle/handleGeneralFormDraft"
	"github.com/dkzhang/RmsGo/webapi/handle/handleLogin"
	"github.com/dkzhang/RmsGo/webapi/handle/handleProject"
	"github.com/dkzhang/RmsGo/webapi/handle/handleProjectRes"
	"github.com/dkzhang/RmsGo/webapi/handle/handleUser"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/middleware"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/workflow/applyChangeResource"
	"github.com/dkzhang/RmsGo/webapi/workflow/applyProjectAndResource"
	"github.com/dkzhang/RmsGo/webapi/workflow/applyReturnComputeRes"
	"github.com/dkzhang/RmsGo/webapi/workflow/applyReturnStorageRes"
	"github.com/dkzhang/RmsGo/webapi/workflow/browseMetering"
	"github.com/gin-gonic/gin"
	"os"
)

func Serve() {
	infra := infrastructure.NewInfrastructure(infrastructure.InfraConfigFile{
		LogMapConf: os.Getenv("LogMapConf"),
		SmsSE:      os.Getenv("SmsSE"),
		DbSE:       os.Getenv("DbSE"),
		LoginConf:  os.Getenv("LoginConf"),
		LoginSec:   os.Getenv("LoginSec"),
	})

	r := gin.Default()
	r.Use(middleware.LoggerGinToFile(infra.TheLogMap))

	/////////////////////////////////////////////////////////////
	theHandleApp := handleApplication.NewHandleApp(infra.TheApplicationDM,
		infra.TheExtractor, infra.TheLogMap)
	theHandleApp.RegisterWorkflow(application.AppTypeNew,
		ApplyProjectAndResource.NewWorkflow(infra.TheApplicationDM, infra.TheProjectDM))
	theHandleApp.RegisterWorkflow(application.AppTypeChange,
		applyChangeResource.NewWorkflow(infra.TheApplicationDM, infra.TheProjectDM, infra.TheLogMap))
	theHandleApp.RegisterWorkflow(application.AppTypeReturnCompute,
		applyReturnComputeRes.NewWorkflow(infra.TheApplicationDM, infra.TheProjectDM, infra.TheProjectResDM, infra.TheLogMap))

	bmwf := browseMetering.NewWorkflow(infra.TheApplicationDM, infra.TheProjectDM, infra.MetClient)
	theHandleApp.RegisterWorkflow(application.AppTypeBrowseMetering, bmwf)
	theHandleApp.RegisterWorkflow(application.AppTypeReturnStorage,
		applyReturnStorageRes.NewWorkflow(infra.TheApplicationDM, infra.TheProjectDM, infra.TheProjectResDM,
			bmwf, infra.TheLogMap))

	theHandleProject := handleProject.NewHandleProject(infra.TheProjectDM,
		infra.TheExtractor, infra.TheLogMap)

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

		hApp := webAPIv1.Group("/Application")
		{
			hApp.GET("/", middleware.TokenAuth(infra), theHandleApp.RetrieveByUserLogin)
			hApp.GET("/ByID/:id", middleware.TokenAuth(infra), theHandleApp.RetrieveByID)
			hApp.GET("/JTBD", middleware.TokenAuth(infra), theHandleApp.RetrieveJTBD)
			hApp.GET("/OpsRecords/:id", middleware.TokenAuth(infra), theHandleApp.RetrieveAppOpsByAppId)

			hApp.POST("/", middleware.TokenAuth(infra), theHandleApp.Create)
			hApp.PUT("/:id", middleware.TokenAuth(infra), theHandleApp.Update)
		}

		hProject := webAPIv1.Group("/Project")
		{
			hProject.GET("/", middleware.TokenAuth(infra), theHandleProject.RetrieveByUserLogin)
			hProject.GET("/:id", middleware.TokenAuth(infra), theHandleProject.RetrieveByID)
			hProject.PUT("/AllocInfo/:id", middleware.TokenAuth(infra), theHandleProject.UpdateAllocInfo)
		}

		theHandleProjectRes := handleProjectRes.NewHandleProjectRes(infra.TheProjectResDM,
			infra.TheProjectDM, infra.TheExtractor, infra.TheLogMap)
		hProjectRes := webAPIv1.Group("/ProjectRes")
		{
			hProjectRes.GET("/Tree/CpuOccupied/:id", middleware.TokenAuth(infra),
				theHandleProjectRes.QueryCpuTreeOccupied)
			hProjectRes.GET("/Tree/CpuAvailable/:id", middleware.TokenAuth(infra),
				theHandleProjectRes.QueryCpuTreeAvailable)
			hProjectRes.GET("/Tree/GpuOccupied/:id", middleware.TokenAuth(infra),
				theHandleProjectRes.QueryGpuTreeOccupied)
			hProjectRes.GET("/Tree/GpuAvailable/:id", middleware.TokenAuth(infra),
				theHandleProjectRes.QueryGpuTreeAvailable)

			hProjectRes.POST("/Schedule/CPU/:id", middleware.TokenAuth(infra),
				theHandleProjectRes.SchedulingCpu)
			hProjectRes.POST("/Schedule/GPU/:id", middleware.TokenAuth(infra),
				theHandleProjectRes.SchedulingGpu)
			hProjectRes.POST("/Schedule/Storage/:id", middleware.TokenAuth(infra),
				theHandleProjectRes.SchedulingStorage)
		}
	}

	/////////////////////////////////////////////////////////////
	r.Run(":8083")
}
