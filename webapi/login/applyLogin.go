package login

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userDM"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Login(c *gin.Context) {

}

func ApplyLogin(c *gin.Context) {
	//获取用户名信息
	userName := c.Query("username")

	//验证用户名合法
	if user.CheckUserName(userName) == false {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无效的用户名",
		})
		return
	}

	//
	udm, err := userDM.GetInstance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	user, err := udm.QueryUserByName(userName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无效的用户名",
		})
		return
	}

	//从数据库中查找该用户(TODO 改为通过userDM中检索)
	user, err := userDM.QueryUserByName(userName, webapi.TheContext.TheDb)
	if err != nil {
		logMap.GetLog(logMap.NORMAL).WithFields(logrus.Fields{
			"UserName": userName,
			"error":    err,
		}).Errorf("Query user from database error.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该用户",
		})
		return
	}

	//从redis中查询，是否有短信锁，（是否有账户锁）
	if webapi.TheContext.TheRedis.IsExist(
		fmt.Sprintf("user_smslock_%d", user.UserID)) == true {

		logMap.GetLog(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": user.UserID,
			"error":  err,
		}).Errorf("user sms lock exist.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "检测到短时间内频繁申请登录，请稍后再试",
		})
		return
	}

	//检查该用户是否已被停用
	//status = 1 可用，-1停用，-9标记删除

	///////////////////////////////////////////////////////////////////////////////////////////////
	//用户验证全通过

	//将该用户信息置于redis中(作废：改用内存DataManagement)

	//userJson, err := json.Marshal(user)
	//if err != nil {
	//	logMap.GetLog(logMap.NORMAL).WithFields(logrus.Fields{
	//		"UserName":  userName,
	//		"User Info": user,
	//		"error":     err,
	//	}).Errorf("json.Marshal user data error.")
	//
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"msg": "服务器处理用户信息出错",
	//	})
	//	return
	//}
	//
	//webapi.TheContext.TheRedis.Set(
	//	fmt.Sprintf("user_inf_%d", user.UserID),
	//	userJson,
	//	time.Second,
	//)

	//生成临时密码，加密后置于redis中并短信发送

}
