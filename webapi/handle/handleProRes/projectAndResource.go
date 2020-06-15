package handleProRes

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"github.com/dkzhang/RmsGo/webapi/model/resourceApplication"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type AppNewProRes struct {
	ProjectName string `json:"project_name"`
	resource.Resource
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func ApplyNewProRes(c *gin.Context) {
	// * Parsing request information
	appWA := resourceApplication.ApplicationWA{}
	if err := c.ShouldBindJSON(&appWA); err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"error": err,
		}).Error("ApplicationWA BindJSON error.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无法解析json"})
		return
	}

	///////////////////////////////////////////////////////////////////////////
	// * Check and Use LoginUser info
	userLoginInfo, err := extractLoginUserInfo.Extract(c)
	if err != nil {
		return
	}

	// only RoleProjectChief can apply new Project and Resource
	if userLoginInfo.Role != user.RoleProjectChief {
	}

	///////////////////////////////////////////////////////////////////////////
	// * Perform the appropriate action
}

func UpdateNewProRes(c *gin.Context) {
	// Parsing request information
	appWA := resourceApplication.ApplicationWA{}
	if err := c.ShouldBindJSON(&appWA); err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"error": err,
		}).Error("ApplicationWA BindJSON error.")

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
