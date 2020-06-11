package handleUser

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/authority/userCRUD"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo(c)
	if err != nil {
		return
	}

	// Load userCreatedInfo from Request
	userCreatedInfo := user.UserInfo{}
	err = c.BindJSON(&userCreatedInfo)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&userCreatedInfo) error.")
		return
	}

	// Set userCreatedInfo default value in rule
	userCreatedInfo.UserID = -1

	switch userLoginInfo.Role {
	case user.RoleApprover:
		userCreatedInfo.Role = user.RoleProjectChief
		userCreatedInfo.Department = userLoginInfo.Department
		userCreatedInfo.DepartmentCode = userLoginInfo.DepartmentCode
	case user.RoleController:
		userCreatedInfo.Role = user.RoleApprover
	case user.RoleProjectChief:
		// Do nothing
	}

	// Check permission
	permission := userCRUD.UserAuthorityCheck(userLoginInfo, userCreatedInfo, userCRUD.OPS_CREATE)

	if permission == false {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":     userLoginInfo.UserID,
			"userCreatedInfo": userCreatedInfo.UserID,
		}).Error("Create user failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	// Insert into userDM
	err = webapi.TheInfras.TheUserDM.InsertUser(userCreatedInfo)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":     userLoginInfo.UserID,
			"userCreatedInfo": userCreatedInfo.UserID,
			"error":           err,
		}).Error("TheUserDM.InsertUser error.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法插入该用户",
		})
		return
	}

	logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":   userLoginInfo,
		"userCreatedInfo": userCreatedInfo,
	}).Info("Delete user success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("创建用户(name=%s)信息成功", userCreatedInfo.UserName),
		"user": userCreatedInfo,
	})
	return

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

	permission := userCRUD.UserAuthorityCheck(userLoginInfo, userAccessedInfo, userCRUD.OPS_RETRIEVE)

	if permission == false {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
		}).Error("Retrieve userInfo failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":    userLoginInfo,
		"userAccessedInfo": userAccessedInfo,
	}).Info("Retrieve userInfo success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("查询用户(id=%d)信息成功", userAccessedInfo.UserID),
		"user": userAccessedInfo,
	})
	return
}

func Update(c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo(c)
	if err != nil {
		return
	}

	userAccessedInfo, err := extractAccessedUserInfo(c)
	if err != nil {
		return
	}

	permission := userCRUD.UserAuthorityCheck(userLoginInfo, userAccessedInfo, userCRUD.OPS_UPDATE)

	if permission == false {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
		}).Error("Update userInfo failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	userUpdatedInfo := user.UserInfo{}
	err = c.BindJSON(&userUpdatedInfo)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
			"error":          err,
		}).Error("c.BindJSON(&userUpdatedInfo) error.")
		return
	}

	userUpdatedInfo.UserID = userAccessedInfo.UserID
	err = webapi.TheInfras.TheUserDM.UpdateUser(userUpdatedInfo)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":     userLoginInfo.UserID,
			"userAccessedID":  userAccessedInfo.UserID,
			"userUpdatedInfo": userUpdatedInfo,
			"error":           err,
		}).Error("TheUserDM.UpdateUser error.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法更新该用户",
		})
		return
	}

	logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":    userLoginInfo,
		"userAccessedInfo": userAccessedInfo,
		"userUpdatedInfo":  userUpdatedInfo,
	}).Info("Delete user success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("更新用户(id=%d)信息成功", userAccessedInfo.UserID),
		"user": userAccessedInfo,
	})
	return
}

func Delete(c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo(c)
	if err != nil {
		return
	}

	userAccessedInfo, err := extractAccessedUserInfo(c)
	if err != nil {
		return
	}

	permission := userCRUD.UserAuthorityCheck(userLoginInfo, userAccessedInfo, userCRUD.OPS_DELETE)

	if permission == false {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
		}).Error("Delete user failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	err = webapi.TheInfras.TheUserDM.DeleteUser(userAccessedInfo.UserID)
	if err != nil {
		logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
			"error":          err,
		}).Error("TheUserDM.DeleteUser error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":    userLoginInfo,
		"userAccessedInfo": userAccessedInfo,
	}).Info("Delete user success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("删除用户(id=%d)信息成功", userAccessedInfo.UserID),
		"user": userAccessedInfo,
	})
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////

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
