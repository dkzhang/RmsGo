package handleProject

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/authority/authApplication"
	"github.com/dkzhang/RmsGo/webapi/authority/authProject"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

	piStatus, err := strconv.Atoi(c.DefaultQuery("status", "-1"))

	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"errS": err,
		}).Error("Parse project status from URL error")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "status参数无效",
		})
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
				"msg": fmt.Sprintf("数据库中查询%s相关项目信息失败", user.RoleStrProjectChief),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"projects": projectFilterByStatus(pros, piStatus),
			"msg":      fmt.Sprintf("查询%s相关项目信息成功", user.RoleStrProjectChief),
		})
	case user.RoleApprover:
		pros, err := h.theProjectDM.QueryByDepartmentCode(userLoginInfo.DepartmentCode)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Project By DC error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("数据库中查询%s相关项目信息失败", user.RoleStrApprover),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"projects": projectFilterByStatus(pros, piStatus),
			"msg":      fmt.Sprintf("查询%s相关项目信息成功", user.RoleStrApprover),
		})
	case user.RoleApprover2:
		pros, err := h.theProjectDM.QueryAllInfo()
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Project All error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("数据库中查询%s相关项目信息失败", user.RoleStrApprover2),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"projects": projectFilterByStatus(pros, piStatus),
			"msg":      fmt.Sprintf("查询%s相关项目信息成功", user.RoleStrApprover2),
		})
	case user.RoleController:
		pros, err := h.theProjectDM.QueryAllInfo()
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
				"userID": userLoginInfo.UserID,
				"error":  err,
			}).Error("Query Project All error")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("数据库中查询%s相关项目信息失败", user.RoleStrController),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"projects": projectFilterByStatus(pros, piStatus),
			"msg":      fmt.Sprintf("查询%s相关项目信息成功", user.RoleStrController),
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

func projectFilterByStatus(pros []project.Info, qStatus int) (prosF []project.Info) {
	for _, pi := range pros {
		if pi.BasicStatus&qStatus != 0 {
			prosF = append(prosF, pi)
		}
	}
	return prosF
}

func (h HandleProject) RetrieveByID(c *gin.Context) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	pi, err := h.extractAccessedProject(c)
	if err != nil {
		return
	}

	permission := authProject.AuthorityCheck(h.theLogMap, userLoginInfo, pi, authApplication.OPS_RETRIEVE)
	if permission == true {
		c.JSON(http.StatusOK, gin.H{
			"pi":  pi,
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

func (h HandleProject) extractAccessedProject(c *gin.Context) (pi project.Info, err error) {
	idStr := c.Param("id")
	pid, err := strconv.Atoi(idStr)

	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get Project ID from gin.Context failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟操作的Project ID无效",
		})
		return project.Info{},
			fmt.Errorf("get Project ID from gin.Context failed: %v", err)
	}

	pi, err = h.theProjectDM.QueryByID(pid)
	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"pid": pid,
		}).Error("ProjectDM.QueryByID (using pid from gin.Context) failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该项目",
		})
		return project.Info{},
			fmt.Errorf("ProjectDM.QueryByID (using pid from gin.Context) error: %v", err)
	}
	return pi, nil
}
