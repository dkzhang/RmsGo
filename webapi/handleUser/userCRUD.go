package handleUser

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/authority"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {

}

func Retrieve(c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo(c)
	if err != nil {
		return
	}

	userAccessedInfo, err := extractAccessedUserInfo(c)
	if err != nil {
		return
	}

	permission := authority.UserAuthorityCheck(userLoginInfo, userAccessedInfo, authority.OPS_RETRIEVE)

	if permission == false {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
			"error":          err,
		}).Error("get userInfo failed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":    userLoginInfo,
		"userAccessedInfo": userAccessedInfo,
	}).Info("get userAccessedID from gin.Context failed.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("查询用户(id=%d)信息成功", userAccessedInfo.UserID),
		"user": userAccessedInfo,
	})
	return
}

func extractLoginUserInfo(c *gin.Context) (userLoginInfo user.UserInfo, err error) {
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

	userLoginInfo, err = webapi.TheInfras.TheUserDM.QueryUserByID(userLoginID)
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

func extractAccessedUserInfo(c *gin.Context) (userAccessedInfo user.UserInfo, err error) {
	idStr := c.Param("id")
	userAccessedID, err := strconv.Atoi(idStr)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get userAccessedID from gin.Context failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟操作的userID无效",
		})
		return user.UserInfo{}, fmt.Errorf("get userAccessedID from gin.Context failed: %v", err)
	}

	userAccessedInfo, err = webapi.TheInfras.TheUserDM.QueryUserByID(userAccessedID)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userAccessedID": userAccessedID,
		}).Error("TheUserDM.QueryUserByID (using userAccessedID from gin.Context) failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该用户",
		})
		return user.UserInfo{}, fmt.Errorf("TheUserDM.QueryUserByID (using userAccessedID from gin.Context) error: %v", err)
	}
	return userAccessedInfo, nil
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
