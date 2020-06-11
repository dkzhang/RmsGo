package handleUser

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {

}

func Retrieve(c *gin.Context) {
	idStr := c.Param("id")
	userRetrieveID, err := strconv.Atoi(idStr)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get userID from gin.Context failed.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟删除的userID无效",
		})
		return
	}

	///////////////////////////////////////////////////////

	userID := c.GetInt("userID")
	if userID < 0 {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userID,
		}).Error("get userID from gin.Context failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	userInfo, err := webapi.TheInfras.TheUserDM.QueryUserByID(userID)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userID,
		}).Error("TheUserDM.QueryUserByID (using userID from gin.Context) failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	///////////////////////////////////////////////////////
	switch userInfo.Role {
	case user.RoleProjectChief:
		if userRetrieveID != userID {
			logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"UserID":         userID,
				"UserRetrieveID": userRetrieveID,
			}).Error("RoleProjectChief has no right to access another user's info.")

			c.JSON(http.StatusForbidden, gin.H{
				"msg": "当前用户无权访问该接口",
			})
			return
		} else {
			logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"UserInfo": userInfo,
			}).Info("RoleProjectChief retrieve his own information success.")
			c.JSON(http.StatusOK, gin.H{
				"msg":  "项目长获取自己信息成功",
				"user": userInfo,
			})
			return
		}
	case user.RoleController:
	case user.RoleApprover:
	}
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	userDeleteID, err := strconv.Atoi(idStr)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get userID from gin.Context failed.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟删除的userID无效",
		})
		return
	}

	userDeleteInfo, err := webapi.TheInfras.TheUserDM.QueryUserByID(userDeleteID)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userDeleteID,
		}).Error("TheUserDM.QueryUserByID (using userID from gin.Context) failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟删除的userID不存在",
		})
		return
	}

	///////////////////////////////////////////////////////

	userID := c.GetInt("userID")
	if userID < 0 {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userID,
		}).Error("get userID from gin.Context failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	userInfo, err := webapi.TheInfras.TheUserDM.QueryUserByID(userID)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userID,
		}).Error("TheUserDM.QueryUserByID (using userID from gin.Context) failed.")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	///////////////////////////////////////////////////////

	switch userInfo.Role {
	case user.RoleProjectChief:
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": userID,
		}).Error("RoleProjectChief has no right to access interface <Delete User>.")

		c.JSON(http.StatusForbidden, gin.H{
			"msg": "当前用户无权访问该接口",
		})
		return
	case user.RoleController:
		err = webapi.TheInfras.TheUserDM.DeleteUser(userDeleteID)
		if err != nil {
			logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"UserInfo":       userInfo,
				"UserDeleteInfo": userDeleteInfo,
			}).Error("TheUserDM.DeleteUser failed.")

			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "服务器内部错误",
			})
			return
		} else {
			logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"UserInfo":       userInfo,
				"UserDeleteInfo": userDeleteInfo,
			}).Info("Delete user success.")

			c.JSON(http.StatusOK, gin.H{
				"msg": fmt.Sprintf("成功删除用户:<%s>", userDeleteInfo.UserName),
			})
			return
		}
	case user.RoleApprover:
		if userInfo.DepartmentCode != userDeleteInfo.DepartmentCode {
			logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"UserInfo":       userInfo,
				"UserDeleteInfo": userDeleteInfo,
			}).Error("delete user failed: RoleApprover cannot delete users from other departments.")

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "错误的Request：RoleApprover不能删除其他部门的用户",
			})
			return
		} else {
			err = webapi.TheInfras.TheUserDM.DeleteUser(userDeleteID)
			if err != nil {
				logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserInfo":       userInfo,
					"UserDeleteInfo": userDeleteInfo,
				}).Error("TheUserDM.DeleteUser failed.")

				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务器内部错误",
				})
				return
			} else {
				logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserInfo":       userInfo,
					"UserDeleteInfo": userDeleteInfo,
				}).Info("Delete user success.")

				c.JSON(http.StatusOK, gin.H{
					"msg": fmt.Sprintf("成功删除用户:<%s>", userDeleteInfo.UserName),
				})
				return
			}
		}
	}

}
