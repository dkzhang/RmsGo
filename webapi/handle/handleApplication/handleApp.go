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
	"github.com/dkzhang/RmsGo/webapi/workflow"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type HandleApp struct {
	theAppWorkflow map[int]workflow.GeneralWorkflow

	theAppDM applicationDM.ApplicationDM

	theExtractor extractLoginUserInfo.Extractor

	theLogMap logMap.LogMap
}

func NewHandleApp(adm applicationDM.ApplicationDM,
	ext extractLoginUserInfo.Extractor, lm logMap.LogMap) HandleApp {
	return HandleApp{
		theAppWorkflow: make(map[int]workflow.GeneralWorkflow),
		theAppDM:       adm,
		theExtractor:   ext,
		theLogMap:      lm,
	}
}

func (h HandleApp) RegisterWorkflow(t int, wl workflow.GeneralWorkflow) {
	h.theAppWorkflow[t] = wl
}

func (h HandleApp) Create(c *gin.Context) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	permission := authApplication.AuthorityCheck(h.theLogMap,
		userLoginInfo, application.Application{},
		authApplication.OPS_CREATE)
	if permission == false {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&GeneralFormDraftCreated) error.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "GeneralForm解析失败",
		})
		return
	}

	appType := gfc.Type
	wf, ok := h.theAppWorkflow[appType]
	if !ok {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"appType": appType,
		}).Error("unsupported application type for create.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("不支持的appType: %d", appType),
		})
		return
	}

	appID, waErr := wf.Apply(gfc, userLoginInfo)
	if waErr != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
}

func (h HandleApp) Update(c *gin.Context) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	app, err := h.extractAccessedApplication(c)
	if err != nil {
		return
	}

	permission := authApplication.AuthorityCheck(h.theLogMap, userLoginInfo, app, authApplication.OPS_UPDATE)
	if permission == false {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON GeneralForm error.")
		return
	}

	// fill attribute
	gf.FormID = app.ApplicationID

	appType := gf.Type
	wf, ok := h.theAppWorkflow[appType]
	if !ok {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"appType": appType,
		}).Error("unsupported application type for update.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("不支持的appType: %d", appType),
		})
		return
	}

	waErr := wf.Process(gf, app, userLoginInfo)

	if waErr != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       waErr.Error(),
		}).Error("TheAppNewWorkflow.Process error.")

		c.JSON(waErr.Type(), gin.H{
			"msg": waErr.Msg(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "新建项目的申请单审批操作成功",
	})
	return
}

func (h HandleApp) RetrieveByUserLogin(c *gin.Context) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	appType, errT := strconv.Atoi(c.DefaultQuery("type", "-1"))
	appStatus, errS := strconv.Atoi(c.DefaultQuery("status", "-1"))

	if errT != nil || errS != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"errT": errT,
			"errS": errS,
		}).Error("Parse appType and appStatus from URL error")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "type和status参数无效",
		})
		return
	}

	switch userLoginInfo.Role {
	case user.RoleProjectChief:
		apps, err := h.theAppDM.QueryByOwner(userLoginInfo.UserID, appType, appStatus)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
		apps, err := h.theAppDM.QueryByDepartmentCode(userLoginInfo.DepartmentCode, appType, appStatus)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
		apps, err := h.theAppDM.QueryAll(appType, appStatus)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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

func (h HandleApp) RetrieveJTBD(c *gin.Context) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	switch userLoginInfo.Role {
	case user.RoleProjectChief:
		apps, err := h.theAppDM.QueryByOwner(userLoginInfo.UserID, application.AppTypeALL, application.AppStatusProjectChief)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
		apps, err := h.theAppDM.QueryByDepartmentCode(userLoginInfo.DepartmentCode, application.AppTypeALL, application.AppStatusApprover)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
		apps, err := h.theAppDM.QueryAll(application.AppTypeALL, application.AppStatusController)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	app, err := h.extractAccessedApplication(c)
	if err != nil {
		return
	}

	permission := authApplication.AuthorityCheck(h.theLogMap, userLoginInfo, app, authApplication.OPS_RETRIEVE)
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
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get Application ID from gin.Context failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟操作的Application ID无效",
		})
		return application.Application{},
			fmt.Errorf("get Application ID from gin.Context failed: %v", err)
	}

	app, err = h.theAppDM.QueryByID(appID)
	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"appID": appID,
		}).Error("theAppDM.QueryByID (using appID from gin.Context) failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该申请表单",
		})
		return application.Application{},
			fmt.Errorf("theAppDM.QueryByID (using appID from gin.Context) error: %v", err)
	}
	return app, nil
}
