package handleLogin

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Login(c *gin.Context) {
	// Get UserName from gin.Context
	userName := c.Query("username")
	passwd := c.Query("passwd")

	// Validate UserName
	if user.CheckUserName(userName) == false {
		logMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
			"UserName": userName,
		}).Error("user login attempt failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无效的用户名",
		})
		return
	}

	// Query User from the UserDM
	userInfo, err := webapi.TheInfras.TheUserDM.QueryUserByName(userName)
	if err != nil {
		logMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
			"UserName": userName,
			"error":    err,
		}).Error("user login attempt failed: Query userInfo from database error.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "该用户不存在",
		})
		return
	}

	// Check if the account status is normal
	if userInfo.Status != user.StatusNormal {
		logMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
			"UserID":   userInfo.UserID,
			"UserName": userName,
			"error":    err,
		}).Error("user login attempt failed: userInfo account status is not normal.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "该用户账号已被停用或删除",
		})
		return
	}

	///////////////////////////////////////////////////////////////////////////////////////////////
	// ValidatePassword
	if webapi.TheInfras.TheUserTempDM.ValidatePassword(userInfo.UserID, passwd) == false {
		logMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
			"UserID":   userInfo.UserID,
			"UserName": userName,
		}).Error("user login attempt failed: user ValidatePassword return false.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "密码错误或密码已过期",
		})
		return
	}

	// Delete Password
	webapi.TheInfras.TheUserTempDM.DelPassword(userInfo.UserID)

	///////////////////////////////////////////////////////////////////////////////////////////////
	// CreateToken
	token, err := webapi.TheInfras.TheUserTempDM.CreateToken(userInfo.UserID)
	if err != nil {
		logMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
			"UserID": userInfo.UserID,
			"error":  err,
		}).Error("user login attempt failed: TheUserTempDM.CreateToken error.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	///////////////////////////////////////////////////////////////////////////////////////////////
	// All pass
	logMap.Log(logMap.NORMAL, logMap.LOGIN).WithFields(logrus.Fields{
		"UserID":   userInfo.UserID,
		"UserName": userName,
	}).Info("user login success.")

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  userInfo,
		"msg":   "用户名密码验证通过，用户登录成功",
	})
	return

}
