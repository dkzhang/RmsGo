package handleAppTemp

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
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

	appTempCreated.UserID = userLoginInfo.UserID
	appTempCreated.ApplicationID = -1

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
	}).Info("Delete user success.")
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("创建临时申请单成功: %v", appTempCreated),
		"id":  id,
	})
	return
}
