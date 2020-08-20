package extractLoginUserInfo

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDM"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Extractor struct {
	theLogMap logMap.LogMap
	theUserDM userDM.UserDM
}

func NewExtractor(lm logMap.LogMap, udm userDM.UserDM) Extractor {
	return Extractor{
		theLogMap: lm,
		theUserDM: udm,
	}
}

func (ext Extractor) Extract(c *gin.Context) (userLoginInfo user.UserInfo, err error) {
	userLoginID := c.GetInt("userID")
	if userLoginID < 0 {
		ext.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginID,
		}).Error("get userLoginID from gin.Context failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return user.UserInfo{}, fmt.Errorf("get userLoginID from gin.Context failed: %v", err)
	}

	userLoginInfo, err = ext.theUserDM.QueryUserByID(userLoginID)
	if err != nil {
		ext.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginID,
		}).Error("theUserDM.QueryUserByID (using userLoginID from gin.Context) failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return user.UserInfo{}, fmt.Errorf("theUserDM.QueryUserByID (using userLoginID from gin.Context) error: %v", err)
	}
	return userLoginInfo, nil
}

func Extract(infra *infrastructure.Infrastructure, c *gin.Context) (userLoginInfo user.UserInfo, err error) {
	userLoginID := c.GetInt("userID")
	if userLoginID < 0 {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginID,
		}).Error("get userLoginID from gin.Context failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return user.UserInfo{}, fmt.Errorf("get userLoginID from gin.Context failed: %v", err)
	}

	userLoginInfo, err = infra.TheUserDM.QueryUserByID(userLoginID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginID,
		}).Error("theUserDM.QueryUserByID (using userLoginID from gin.Context) failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return user.UserInfo{}, fmt.Errorf("theUserDM.QueryUserByID (using userLoginID from gin.Context) error: %v", err)
	}
	return userLoginInfo, nil
}
