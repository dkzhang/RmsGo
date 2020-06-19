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
)

func Create(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	// Load userCreatedInfo from Request
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
	permission := authAppTemp.AppTempAuthorityCheck(infra.TheLogMap, userLoginInfo, appTempCreated, authAppTemp.OPS_UPDATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"appTempCreated": appTempCreated,
		}).Error("Update userInfo failed, since UserAuthorityCheck permission not allowed.")
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
		"msg": fmt.Sprintf("创建临时申请单成功: %v", appTempCreated),
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

	// Load userCreatedInfo from Request
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

	// authentication
	permission := authAppTemp.AppTempAuthorityCheck(infra.TheLogMap, userLoginInfo, appTempUpdated, authAppTemp.OPS_UPDATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"appTempUpdated": appTempUpdated,
		}).Error("Update userInfo failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}
}

func Delete(infra *infrastructure.Infrastructure, c *gin.Context) {

}
