package handleAppTemp

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	authAppTemp "github.com/dkzhang/RmsGo/webapi/authority/appTempCRUD"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/model/appTemp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Create(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	// Load appTemp CreatedInfo from Request
	appTempCreated := appTemp.AppTemp{}
	err = c.BindJSON(&appTempCreated)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&appTempCreated) error.")
		return
	}

	// fill attribute
	appTempCreated.UserID = userLoginInfo.UserID
	appTempCreated.ApplicationID = -1

	// authentication
	permission := authAppTemp.AppTempAuthorityCheck(infra.TheLogMap,
		userLoginInfo, appTempCreated, authAppTemp.OPS_CREATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"appTempCreated": appTempCreated,
		}).Error("Create AppTemp failed, since AppTempAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	// Insert into DB
	id, err := infra.TheAppTempDB.InsertAppTemp(appTempCreated)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"appTempCreated": appTempCreated,
			"error":          err,
		}).Error("TheAppTempDB.InsertAppTemp error.")
		return
	}

	// success
	appTempCreated.ApplicationID = id
	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":  userLoginInfo,
		"appTempCreated": appTempCreated,
	}).Info("Create appTemp success.")
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("创建申请表草稿成功: %v", appTempCreated),
		"id":  id,
	})
	return
}

func Retrieve(infra *infrastructure.Infrastructure, c *gin.Context) {

}

func Update(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	appTempAccessedInfo, err := extractAccessedAppTempInfo(infra, c)
	if err != nil {
		return
	}

	// Load AppTemp Update Info from Request
	appTempUpdated := appTemp.AppTemp{}
	err = c.BindJSON(&appTempUpdated)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&appTempUpdated) error.")
		return
	}

	// fill attribute
	appTempUpdated.UserID = userLoginInfo.UserID
	appTempUpdated.ApplicationID = appTempAccessedInfo.ApplicationID

	// authentication
	permission := authAppTemp.AppTempAuthorityCheck(infra.TheLogMap,
		userLoginInfo, appTempUpdated, authAppTemp.OPS_UPDATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"appTempUpdated": appTempUpdated,
		}).Error("Update AppTemp failed, since AppTempAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	// Update to the DB
	err = infra.TheAppTempDB.UpdateAppTemp(appTempUpdated)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"appTempUpdated": appTempUpdated,
			"error":          err,
		}).Error("TheAppTempDB.UpdateAppTemp error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	// all success
	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":  userLoginInfo,
		"appTempUpdated": appTempUpdated,
	}).Info("Update AppTemp success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":            fmt.Sprintf("更新申请表草稿(id=%d)信息成功", appTempUpdated.ApplicationID),
		"appTempUpdated": appTempUpdated,
	})
	return
}

func Delete(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	appTempAccessedInfo, err := extractAccessedAppTempInfo(infra, c)
	if err != nil {
		return
	}

	permission := authAppTemp.AppTempAuthorityCheck(infra.TheLogMap,
		userLoginInfo, appTempAccessedInfo, authAppTemp.OPS_DELETE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":         userLoginInfo.UserID,
			"appTempAccessedInfo": appTempAccessedInfo,
		}).Error("Delete AppTemp failed, since AppTemp AuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	err = infra.TheAppTempDB.DeleteAppTemp(appTempAccessedInfo.ApplicationID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":         userLoginInfo.UserID,
			"appTempAccessedInfo": appTempAccessedInfo,
			"error":               err,
		}).Error("TheAppTempDB.DeleteAppTemp error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":       userLoginInfo,
		"appTempAccessedInfo": appTempAccessedInfo,
	}).Info("Delete appTemp success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":     fmt.Sprintf("删除申请表草稿(id=%d)信息成功", appTempAccessedInfo.ApplicationID),
		"appTemp": appTempAccessedInfo,
	})
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func extractAccessedAppTempInfo(infra *infrastructure.Infrastructure, c *gin.Context) (appTempAccessedInfo appTemp.AppTemp, err error) {
	idStr := c.Param("id")
	appTempAccessedID, err := strconv.Atoi(idStr)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get appTemp ID from gin.Context failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟操作的appTempID无效",
		})
		return appTemp.AppTemp{}, fmt.Errorf("get appTemp ID from gin.Context failed: %v", err)
	}

	appTempAccessedInfo, err = infra.TheAppTempDB.QueryAppTempByID(appTempAccessedID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"appTempAccessedID": appTempAccessedID,
		}).Error("TheAppTempDB.QueryAppTempByID (using appTempAccessedID from gin.Context) failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该AppTemp",
		})
		return appTemp.AppTemp{}, fmt.Errorf("TheUserDM.QueryUserByID (using userAccessedID from gin.Context) error: %v", err)
	}
	return appTempAccessedInfo, nil
}
