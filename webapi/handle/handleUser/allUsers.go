package handleUser

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AllUsers(infra *infrastructure.Infrastructure, c *gin.Context) {
	userID := c.GetInt("userID")
	if userID < 0 {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userID,
		}).Error("get userID from gin.Context failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	userInfo, err := infra.TheUserDM.QueryUserByID(userID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userID,
		}).Error("theUserDM.QueryUserByID (using userID from gin.Context) failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	switch userInfo.Role {
	case user.RoleProjectChief:
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userID,
		}).Error("RoleProjectChief has no right to access interface <AllUsers>.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "当前用户无权访问该接口",
		})
		return
	case user.RoleController:
		users := infra.TheUserDM.QueryUserByFilter(func(user.UserInfo) bool { return true })
		c.JSON(http.StatusOK, gin.H{
			"msg":   fmt.Sprintf("查询到%d个用户信息", len(users)),
			"users": users,
		})
		return
	case user.RoleApprover:
		usersd := infra.TheUserDM.QueryUserByDepartmentCode(userInfo.DepartmentCode)
		c.JSON(http.StatusOK, gin.H{
			"msg":   fmt.Sprintf("查询到%d个用户信息", len(usersd)),
			"users": usersd,
		})
		return
	}
}
