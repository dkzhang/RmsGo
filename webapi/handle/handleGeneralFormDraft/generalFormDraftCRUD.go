package handleGeneralFormDraft

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/authority/authGeneralFormDraftCRUD"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/dkzhang/RmsGo/webapi/model/generalFormDraft"
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

	// Load GeneralFormDraft CreatedInfo from Request
	GeneralFormDraftCreated := generalFormDraft.GeneralFormDraft{}
	err = c.BindJSON(&GeneralFormDraftCreated)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&GeneralFormDraftCreated) error.")
		return
	}

	// fill attribute
	GeneralFormDraftCreated.UserID = userLoginInfo.UserID
	GeneralFormDraftCreated.FormID = -1

	// authentication
	permission := authGeneralFormDraftCRUD.AuthorityCheck(infra.TheLogMap,
		userLoginInfo, GeneralFormDraftCreated, authGeneralFormDraftCRUD.OPS_CREATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":             userLoginInfo.UserID,
			"GeneralFormDraftCreated": GeneralFormDraftCreated,
		}).Error("Create GeneralFormDraft failed, since AuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	// Insert into DB
	id, err := infra.TheGeneralFormDraftDB.InsertGeneralFormDraft(GeneralFormDraftCreated)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":             userLoginInfo.UserID,
			"GeneralFormDraftCreated": GeneralFormDraftCreated,
			"error":                   err,
		}).Error("TheGeneralFormDraftDB.InsertGeneralFormDraft error.")
		return
	}

	// success
	GeneralFormDraftCreated.FormID = id
	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":           userLoginInfo,
		"GeneralFormDraftCreated": GeneralFormDraftCreated,
	}).Info("Create GeneralFormDraft success.")
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("创建申请表草稿成功: %v", GeneralFormDraftCreated),
		"id":  id,
	})
	return
}

func RetrieveByOwner(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	// skip authentication

	theGeneralFormDrafts, err := infra.TheGeneralFormDraftDB.QueryGeneralFormDraftByOwner(userLoginInfo.UserID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginInfo": userLoginInfo,
			"error":         err,
		}).Error("TheGeneralFormDraftDB.QueryGeneralFormDraftByOwner error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	// all success
	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":        userLoginInfo,
		"theGeneralFormDrafts": theGeneralFormDrafts,
	}).Info("Retrieve GeneralFormDraft by Owner success.")
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("查询到当前登录用户(id = %d)的所有申请表草稿信息成功，共计%d份",
			userLoginInfo.UserID, len(theGeneralFormDrafts)),
		"GeneralFormDraft": theGeneralFormDrafts,
	})
	return
}

func RetrieveByID(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	GeneralFormDraftAccessed, err := extractAccessedGeneralFormDraftInfo(infra, c)
	if err != nil {
		return
	}

	// authentication
	permission := authGeneralFormDraftCRUD.AuthorityCheck(infra.TheLogMap,
		userLoginInfo, GeneralFormDraftAccessed, authGeneralFormDraftCRUD.OPS_UPDATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":              userLoginInfo.UserID,
			"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
		}).Error("Update GeneralFormDraft failed, since AuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	// all success
	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":            userLoginInfo,
		"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
	}).Info("Retrieve GeneralFormDraft by ID success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":                      fmt.Sprintf("查询申请表草稿(id=%d)信息成功", GeneralFormDraftAccessed.FormID),
		"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
	})
	return
}

func Update(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	GeneralFormDraftAccessed, err := extractAccessedGeneralFormDraftInfo(infra, c)
	if err != nil {
		return
	}

	// Load GeneralFormDraft Update Info from Request
	GeneralFormDraftUpdated := generalFormDraft.GeneralFormDraft{}
	err = c.BindJSON(&GeneralFormDraftUpdated)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&GeneralFormDraftUpdated) error.")
		return
	}

	// fill attribute
	GeneralFormDraftUpdated.UserID = userLoginInfo.UserID
	GeneralFormDraftUpdated.FormID = GeneralFormDraftAccessed.FormID

	// authentication
	permission := authGeneralFormDraftCRUD.AuthorityCheck(infra.TheLogMap,
		userLoginInfo, GeneralFormDraftAccessed, authGeneralFormDraftCRUD.OPS_UPDATE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":              userLoginInfo.UserID,
			"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
		}).Error("Update GeneralFormDraft failed, since AuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	// Update to the DB
	err = infra.TheGeneralFormDraftDB.UpdateGeneralFormDraft(GeneralFormDraftUpdated)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":             userLoginInfo.UserID,
			"GeneralFormDraftUpdated": GeneralFormDraftUpdated,
			"error":                   err,
		}).Error("TheGeneralFormDraftDB.UpdateGeneralFormDraft error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	// all success
	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":           userLoginInfo,
		"GeneralFormDraftUpdated": GeneralFormDraftUpdated,
	}).Info("Update GeneralFormDraft success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":                     fmt.Sprintf("更新申请表草稿(id=%d)信息成功", GeneralFormDraftUpdated.FormID),
		"GeneralFormDraftUpdated": GeneralFormDraftUpdated,
	})
	return
}

func Delete(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := extractLoginUserInfo.Extract(infra, c)
	if err != nil {
		return
	}

	GeneralFormDraftAccessed, err := extractAccessedGeneralFormDraftInfo(infra, c)
	if err != nil {
		return
	}

	permission := authGeneralFormDraftCRUD.AuthorityCheck(infra.TheLogMap,
		userLoginInfo, GeneralFormDraftAccessed, authGeneralFormDraftCRUD.OPS_DELETE)

	if permission == false {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":              userLoginInfo.UserID,
			"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
		}).Error("Delete GeneralFormDraft failed, since GeneralFormDraft AuthorityCheck permission not allowed.")
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "此用户无权访问该数据",
		})
		return
	}

	err = infra.TheGeneralFormDraftDB.DeleteGeneralFormDraft(GeneralFormDraftAccessed.FormID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID":              userLoginInfo.UserID,
			"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
			"error":                    err,
		}).Error("TheGeneralFormDraftDB.DeleteGeneralFormDraft error.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}

	infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"userLoginInfo":            userLoginInfo,
		"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
	}).Info("Delete GeneralFormDraft success.")
	c.JSON(http.StatusOK, gin.H{
		"msg":              fmt.Sprintf("删除申请表草稿(id=%d)信息成功", GeneralFormDraftAccessed.FormID),
		"GeneralFormDraft": GeneralFormDraftAccessed,
	})
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func extractAccessedGeneralFormDraftInfo(infra *infrastructure.Infrastructure, c *gin.Context) (GeneralFormDraftAccessed generalFormDraft.GeneralFormDraft, err error) {
	idStr := c.Param("id")
	GeneralFormDraftAccessedID, err := strconv.Atoi(idStr)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get GeneralFormDraft ID from gin.Context failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟操作的GeneralFormDraftID无效",
		})
		return generalFormDraft.GeneralFormDraft{}, fmt.Errorf("get GeneralFormDraft ID from gin.Context failed: %v", err)
	}

	GeneralFormDraftAccessed, err = infra.TheGeneralFormDraftDB.QueryGeneralFormDraftByID(GeneralFormDraftAccessedID)
	if err != nil {
		infra.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"GeneralFormDraftAccessedID": GeneralFormDraftAccessedID,
		}).Error("TheGeneralFormDraftDB.QueryGeneralFormDraftByID (using GeneralFormDraftAccessedID from gin.Context) failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该GeneralFormDraft",
		})
		return generalFormDraft.GeneralFormDraft{}, fmt.Errorf("theUserDM.QueryUserByID (using userAccessedID from gin.Context) error: %v", err)
	}
	return GeneralFormDraftAccessed, nil
}
