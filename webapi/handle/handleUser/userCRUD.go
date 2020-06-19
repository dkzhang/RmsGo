package handleUser

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/authority/userCRUD"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Create(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	// Load userCreatedInfo from Request
	userCreatedInfo := user.UserInfo{}
	err = c.BindJSON(&userCreatedInfo)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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

	//userCreatedInfo.UserName = user.StandardizedUserName(userCreatedInfo.UserName, userCreatedInfo.DepartmentCode)

	// Check permission
	permission := userCRUD.UserAuthorityCheck(infra.TheLogMap, userLoginInfo, userCreatedInfo, userCRUD.OPS_CREATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":     userLoginInfo.UserID,
			"userCreatedInfo": userCreatedInfo.UserID,
		}).Error("Create user failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	// Insert pre-check
	msg, err := infra.TheUserDM.InsertUserPreCheck(userCreatedInfo)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginInfo":   userLoginInfo,
			"userCreatedInfo": userCreatedInfo,
			"error":           err,
		}).Error("TheUserDM.InsertUserPreCheck error.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": msg,
		})
		return
	}

	// Insert into userDM
	err = infra.TheUserDM.InsertUser(userCreatedInfo)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":     userLoginInfo.UserID,
			"userCreatedInfo": userCreatedInfo.UserID,
			"error":           err,
		}).Error("TheUserDM.InsertUser error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":   userLoginInfo,
		"userCreatedInfo": userCreatedInfo,
	}).Info("Create user success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("创建用户(name=%s)信息成功", userCreatedInfo.UserName),
		"user": userCreatedInfo,
	})
	return

}

func Retrieve(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	userAccessedInfo, err := extractAccessedUserInfo(infra, c)
	if err != nil {
		return
	}

	permission := userCRUD.UserAuthorityCheck(infra.TheLogMap, userLoginInfo, userAccessedInfo, userCRUD.OPS_RETRIEVE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
		}).Error("Retrieve userInfo failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":    userLoginInfo,
		"userAccessedInfo": userAccessedInfo,
	}).Info("Retrieve userInfo success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("查询用户(id=%d)信息成功", userAccessedInfo.UserID),
		"user": userAccessedInfo,
	})
	return
}

func Update(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	userAccessedInfo, err := extractAccessedUserInfo(infra, c)
	if err != nil {
		return
	}

	permission := userCRUD.UserAuthorityCheck(infra.TheLogMap, userLoginInfo, userAccessedInfo, userCRUD.OPS_UPDATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
		}).Error("Update userInfo failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	// load userUpdatedInfo from request json
	userUpdatedInfo := user.UserInfo{}
	err = c.BindJSON(&userUpdatedInfo)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
			"error":          err,
		}).Error("c.BindJSON(&userUpdatedInfo) error.")
		return
	}

	// set updated user's userID, role, userName
	userUpdatedInfo.UserID = userAccessedInfo.UserID
	userUpdatedInfo.Role = userAccessedInfo.Role
	//userUpdatedInfo.UserName = user.StandardizedUserName(userUpdatedInfo.UserName, userUpdatedInfo.DepartmentCode)

	// update pre-check
	msg, err := infra.TheUserDM.UpdateUserPreCheck(userUpdatedInfo)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":     userLoginInfo.UserID,
			"userAccessedID":  userAccessedInfo.UserID,
			"userUpdatedInfo": userUpdatedInfo,
			"error":           err,
		}).Error("TheUserDM.UpdateUserPreCheck error.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": msg,
		})
		return
	}

	// update in userDM
	err = infra.TheUserDM.UpdateUser(userUpdatedInfo)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":     userLoginInfo.UserID,
			"userAccessedID":  userAccessedInfo.UserID,
			"userUpdatedInfo": userUpdatedInfo,
			"error":           err,
		}).Error("TheUserDM.UpdateUser error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	// all success
	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":    userLoginInfo,
		"userAccessedInfo": userAccessedInfo,
		"userUpdatedInfo":  userUpdatedInfo,
	}).Info("Update user success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("更新用户(id=%d)信息成功", userUpdatedInfo.UserID),
		"user": userUpdatedInfo,
	})
	return
}

func Delete(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	userAccessedInfo, err := extractAccessedUserInfo(infra, c)
	if err != nil {
		return
	}

	permission := userCRUD.UserAuthorityCheck(infra.TheLogMap, userLoginInfo, userAccessedInfo, userCRUD.OPS_DELETE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
		}).Error("Delete user failed, since UserAuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	err = infra.TheUserDM.DeleteUser(userAccessedInfo.UserID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":    userLoginInfo.UserID,
			"userAccessedID": userAccessedInfo.UserID,
			"error":          err,
		}).Error("TheUserDM.DeleteUser error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
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

func extractAccessedUserInfo(infra *infrastructure.Infrastructure, c *gin.Context) (userAccessedInfo user.UserInfo, err error) {
	idStr := c.Param("id")
	userAccessedID, err := strconv.Atoi(idStr)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get userAccessedID from gin.Context failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟操作的userID无效",
		})
		return user.UserInfo{}, fmt.Errorf("get userAccessedID from gin.Context failed: %v", err)
	}

	userAccessedInfo, err = infra.TheUserDM.QueryUserByID(userAccessedID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userAccessedID": userAccessedID,
		}).Error("TheUserDM.QueryUserByID (using userAccessedID from gin.Context) failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该用户",
		})
		return user.UserInfo{}, fmt.Errorf("TheUserDM.QueryUserByID (using userAccessedID from gin.Context) error: %v", err)
	}
	return userAccessedInfo, nil
}
