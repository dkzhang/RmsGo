package handleUser

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RetrieveUserLogin(infra *webapi.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo": userLoginInfo,
	}).Info("Retrieve userInfo success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("获取当前登录用户(id=%d)信息成功", userLoginInfo.UserID),
		"user": userLoginInfo,
	})
	return
}
