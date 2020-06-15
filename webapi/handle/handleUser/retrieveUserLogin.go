package handleUser

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RetrieveUserLogin(c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo(c)
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
