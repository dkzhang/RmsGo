package extractLoginUserInfo

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Extract(infra *webapi.Infrastructure, c *gin.Context) (userLoginInfo user.UserInfo, err error) {
	userLoginID := c.GetInt("userID")
	if userLoginID < 0 {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginID,
		}).Error("get userLoginID from gin.Context failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return user.UserInfo{}, fmt.Errorf("get userLoginID from gin.Context failed: %v", err)
	}

	userLoginInfo, err = infra.TheUserDM.QueryUserByID(userLoginID)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginID,
		}).Error("TheUserDM.QueryUserByID (using userLoginID from gin.Context) failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return user.UserInfo{}, fmt.Errorf("TheUserDM.QueryUserByID (using userLoginID from gin.Context) error: %v", err)
	}
	return userLoginInfo, nil
}
