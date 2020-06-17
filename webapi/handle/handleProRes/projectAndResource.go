package handleProRes

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ApplyNewProRes(infra *infrastructure.Infrastructure, c *gin.Context) {
	// * Parsing request information
	appWA := application.GeneralApplication{}
	if err := c.ShouldBindJSON(&appWA); err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"error": err,
		}).Error("GeneralApplication BindJSON error.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无法解析json"})
		return
	}

	///////////////////////////////////////////////////////////////////////////
	// * Check and Use LoginUser info
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	// only RoleProjectChief can apply new Project and Resource
	if userLoginInfo.Role != user.RoleProjectChief {
	}

	///////////////////////////////////////////////////////////////////////////
	// * Perform the appropriate action
}

func UpdateNewProRes(infra *infrastructure.Infrastructure, c *gin.Context) {
	// Parsing request information
	appWA := application.GeneralApplication{}
	if err := c.ShouldBindJSON(&appWA); err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"error": err,
		}).Error("GeneralApplication BindJSON error.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无法解析json"})
		return
	}

	// Check LoginUser info

	// Perform the appropriate action

}

type AppExResWA struct {
}

type AppReComWA struct {
}

type AppReStoWA struct {
}

type AppMeterWA struct {
}
