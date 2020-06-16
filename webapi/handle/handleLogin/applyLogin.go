package handleLogin

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/myUtils/shortMessageService"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ApplyLogin(infra *infrastructure.Infrastructure, c *gin.Context) {
	// Get UserName from gin.Context
	userName := c.Query("username")

	// Validate UserName
	if user.CheckUserName(userName) == false {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无效的用户名",
		})
		return
	}

	// Query User from the UserDM
	userInfo, err := infra.TheUserDM.QueryUserByName(userName)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserName": userName,
			"error":    err,
		}).Error("Query userInfo from database error.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "该用户不存在",
		})
		return
	}

	// Check if the account status is normal
	if userInfo.Status != user.StatusNormal {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userInfo.UserID,
			"error":  err,
		}).Error("userInfo account status is not normal.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "该用户账号已被停用或删除",
		})
		return
	}

	// Check if the account is in sms-locked status
	if infra.TheUserTempDM.IsSmsLock(userInfo.UserID) == true {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID":   userInfo.UserID,
			"UserName": userInfo.UserName,
		}).Error("userInfo account is in sms-locked status.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "检测到短时间内频繁申请登录，请稍后再试",
		})
		return
	}

	///////////////////////////////////////////////////////////////////////////////////////////////
	// All verifications pass

	//Generate temporary password
	passwd, err := infra.TheUserTempDM.SetPassword(userInfo.UserID)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userInfo.UserID,
			"error":  err,
		}).Error("TheUserTempDM.SetPassword error.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	// Send temporary password to user's mobile phone by SMS
	resp, err := infra.TheSmsService.SendSMS(shortMessageService.MessageContent{
		PhoneNumberSet: []string{shortMessageService.ChineseMobile(userInfo.Mobile)},
		TemplateParamSet: []string{userInfo.ChineseName, passwd,
			fmt.Sprintf("%.1f", infra.TheLoginConfig.ThePasswordConfig.Expire.Minutes())},
	})
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userInfo.UserID,
			"Mobile": userInfo.Mobile,
			"error":  err,
		}).Error("TheSmsService.SendSMS error.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}
	logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
		"UserID": userInfo.UserID,
		"Mobile": userInfo.Mobile,
		"resp":   resp,
	}).Info("TheSmsService.SendSMS success.")

	// Set SMS lock
	infra.TheUserTempDM.LockSms(userInfo.UserID)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userInfo.UserID,
			"error":  err,
		}).Error("TheUserTempDM.LockSms error.")
	}

	///////////////////////////////////////////////////////////////////////////////////////////////
	// All pass
	c.JSON(http.StatusOK, gin.H{
		"msg": "用户名验证成功，密码已短信发送",
	})
	return

}
