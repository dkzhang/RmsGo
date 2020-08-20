package handleApplication

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/authority/authApplication"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDM"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/generalForm"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/dkzhang/RmsGo/webapi/workflow/ApplyProjectAndResource"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type HandleApp struct {
	TheAppNewWorkflow           ApplyProjectAndResource.Workflow
	TheAppChangeWorkflow        interface{}
	TheAppReturnComputeWorkflow interface{}
	TheAppReturnStorageWorkflow interface{}

	TheAppDM applicationDM.ApplicationDM

	TheExtractor extractLoginUserInfo.Extractor

	TheLogMap logMap.LogMap
}

func NewHandleApp(nwf ApplyProjectAndResource.Workflow) (h HandleApp) {
	return h
}

func (h HandleApp) Create(c *gin.Context) {
	userLoginInfo, err := h.TheExtractor.Extract(c)
	if err != nil {
		return
	}

	permission := authApplication.AuthorityCheck(h.TheLogMap,
		userLoginInfo, application.Application{},
		authApplication.OPS_CREATE)
	if permission == false {
		h.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
		}).Error(" AuthorityCheck failed.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "当前用户无权提交申请表单",
		})
		return
	}

	// Load GeneralFormDraft CreatedInfo from Request
	gfc := generalForm.GeneralForm{}
	err = c.BindJSON(&gfc)
	if err != nil {
		h.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&GeneralFormDraftCreated) error.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "GeneralForm解析失败",
		})
		return
	}

	appType := c.GetInt("type")
	switch appType {
	case application.AppTypeNew:
		appID, waErr := h.TheAppNewWorkflow.Apply(gfc, userLoginInfo)
		if waErr != nil {
			h.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"userLoginID": userLoginInfo.UserID,
				"error":       waErr.Error(),
			}).Error("TheAppNewWorkflow.Apply error.")

			c.JSON(waErr.Type(), gin.H{
				"msg": waErr.Msg(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"appID": appID,
			"msg":   "新建项目的申请单创建成功",
		})
		return
	case application.AppTypeChange:
	case application.AppTypeReturnCompute:
	case application.AppTypeReturnStorage:
	default:
		h.TheLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
			"appType": appType,
		}).Error("unsupported application type for create.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("不支持的appType: %s", appType),
		})
		return
	}
}

func (h HandleApp) Update(c *gin.Context) {
	userLoginInfo, err := h.TheExtractor.Extract(c)
	if err != nil {
		return
	}

	app, err := h.extractAccessedApplication(c)
	if err != nil {
		return
	}

	permission := authApplication.AuthorityCheck(h.TheLogMap, userLoginInfo, app, authApplication.OPS_UPDATE)
	if permission == false {
		h.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"appID":       app.ApplicationID,
		}).Error(" AuthorityCheck failed.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "当前用户无权操作该表单",
		})
		return
	}

	gf := generalForm.GeneralForm{}
	err = c.BindJSON(&gf)
	if err != nil {
		h.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON GeneralForm error.")
		return
	}

	// fill attribute
	gf.FormID = app.ApplicationID

	waErr := h.TheAppNewWorkflow.Process(gf, userLoginInfo)

	if waErr != nil {
		h.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       waErr.Error(),
		}).Error("TheAppNewWorkflow.Process error.")

		c.JSON(waErr.Type(), gin.H{
			"msg": waErr.Msg(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "新建项目的申请单操作成功",
	})
	return
}

func (h HandleApp) RetrieveByUserLogin(c *gin.Context) {
	userLoginInfo, err := h.TheExtractor.Extract(c)
	if err != nil {
		return
	}

	appType := c.GetInt("type")
	appStatus := c.GetInt("status")

	switch userLoginInfo.Role {
	case user.RoleProjectChief:
		apps, err := h.TheAppDM.QueryByOwner(userLoginInfo.UserID, appType, appStatus)
		if err != nil {
			h.TheLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Application By Owner error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "数据库中查询项目长相关申请单失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"apps": apps,
			"msg":  "查询项目长相关申请单成功",
		})
	case user.RoleApprover:
		apps, err := h.TheAppDM.QueryByDepartmentCode(userLoginInfo.DepartmentCode, appType, appStatus)
		if err != nil {
			h.TheLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Application By Owner error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "数据库中查询审批人相关申请单失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"apps": apps,
			"msg":  "查询项目长相关申请单成功",
		})
	case user.RoleController:
		apps, err := h.TheAppDM.QueryAll(appType, appStatus)
		if err != nil {
			h.TheLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Application By Owner error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "数据库中查询调度员相关申请单失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"apps": apps,
			"msg":  "查询项目长相关申请单成功",
		})
	default:
		h.TheLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
			"userID": userLoginInfo.UserID,
			"Role":   userLoginInfo.Role,
			"error":  err,
		}).Error("error: unsupported user type")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户身份代码不支持",
		})
		return
	}
}

func (h HandleApp) RetrieveByID(c *gin.Context) {
	userLoginInfo, err := h.TheExtractor.Extract(c)
	if err != nil {
		return
	}

	app, err := h.extractAccessedApplication(c)
	if err != nil {
		return
	}

	permission := authApplication.AuthorityCheck(h.TheLogMap, userLoginInfo, app, authApplication.OPS_RETRIEVE)
	if permission == true {
		c.JSON(http.StatusOK, gin.H{
			"app": app,
			"msg": "查询成功",
		})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "无权访问",
		})
		return
	}
}

func (h HandleApp) extractAccessedApplication(c *gin.Context) (app application.Application, err error) {
	idStr := c.Param("id")
	appID, err := strconv.Atoi(idStr)

	if err != nil {
		h.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get GeneralFormDraft ID from gin.Context failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟操作的GeneralFormDraftID无效",
		})
		return application.Application{},
			fmt.Errorf("get Application ID from gin.Context failed: %v", err)
	}

	app, err = h.TheAppDM.QueryByID(appID)
	if err != nil {
		h.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"appID": appID,
		}).Error("TheAppDM.QueryByID (using appID from gin.Context) failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该申请表单",
		})
		return application.Application{},
			fmt.Errorf("TheAppDM.QueryByID (using appID from gin.Context) error: %v", err)
	}
	return app, nil
}
