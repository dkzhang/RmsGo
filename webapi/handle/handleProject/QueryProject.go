package handleProject

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HandleProject struct {
	theProjectDM projectDM.ProjectDM

	theExtractor extractLoginUserInfo.Extractor

	theLogMap logMap.LogMap
}

func NewHandleProject(pdm projectDM.ProjectDM,
	ext extractLoginUserInfo.Extractor, lm logMap.LogMap) HandleProject {
	return HandleProject{
		theProjectDM: pdm,
		theExtractor: ext,
		theLogMap:    lm,
	}
}

func (h HandleProject) RetrieveByUserLogin(c *gin.Context) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	switch userLoginInfo.Role {
	case user.RoleProjectChief:
		pros, err := h.theProjectDM.QueryByOwner(userLoginInfo.UserID)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Project By Owner error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "数据库中查询项目长相关项目信息失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"projects": pros,
			"msg":      "查询项目长相关申请单成功",
		})
	case user.RoleApprover:
		pros, err := h.theProjectDM.QueryByDepartmentCode(userLoginInfo.DepartmentCode)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Project By DC error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "数据库中查询审批人相关项目信息失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"projects": pros,
			"msg":      "查询项目长相关申请单成功",
		})
	case user.RoleController:
		pros, err := h.theProjectDM.QueryAllInfo()
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Project All error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "数据库中查询调度员相关申请单失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"projects": pros,
			"msg":      "查询项目长相关申请单成功",
		})
	default:
		h.theLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
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

//func (h HandleProject) RetrieveByID(c *gin.Context) {
//
//}
