package login

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Login(c *gin.Context) {

}

func ApplyLogin(c *gin.Context) {
	//获取用户名信息
	userName := ""

	//验证用户名合法
	if user.CheckUserName(userName) == false {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "用户名错误",
		})
		return
	}

	//从redis中查询，是否有短信锁，（是否有账户锁）

	//从数据库中查找该用户
	user, err := user.QueryUserByName(userName, webapi.TheContext.TheDb)
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

	//将该用户信息置于redis中
	userJson, err := json.Marshal(user)
	if err != nil {
		logMap.GetLog(logMap.NORMAL).WithFields(logrus.Fields{
			"UserName":  userName,
			"User Info": user,
			"error":     err,
		}).Errorf("json.Marshal user data error.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器处理用户信息出错",
		})
		return
	}

	webapi.TheContext.TheRedis.Set(
		fmt.Sprintf("user_%d", user.UserID),
		userJson,
		time.Second,
	)

	//生成临时密码，加密后置于redis中并短信发送

}
