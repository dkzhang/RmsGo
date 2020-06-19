package appTempCRUD

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/model/appTemp"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/sirupsen/logrus"
)

type appTempAuthority struct {
	RelationShipBetween func(user.UserInfo, appTemp.AppTemp) bool
	Operation           int
	Permission          bool
	Description         string
}

const (
	OPS_CREATE = 1 << iota
	OPS_RETRIEVE
	OPS_UPDATE
	OPS_DELETE
)

var theAppTempAuthorityTable []appTempAuthority

func init() {
	theAppTempAuthorityTable = make([]appTempAuthority, 0)

	theAppTempAuthorityTable = append(theAppTempAuthorityTable, appTempAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, appTempAccessed appTemp.AppTemp) bool {
			if userLoginInfo.Role == user.RoleProjectChief &&
				appTempAccessed.UserID == userLoginInfo.UserID {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_CREATE | OPS_RETRIEVE | OPS_UPDATE | OPS_DELETE,
		Permission:  true,
		Description: "Allow RoleProjectChief CRUD appTemp owned by himself",
	})
}

func AppTempAuthorityCheck(theLogMap logMap.LogMap,
	userLoginInfo user.UserInfo, appTempAccessed appTemp.AppTemp, ops int) (permission bool) {
	for _, rule := range theAppTempAuthorityTable {
		if (ops&rule.Operation != 0) && rule.RelationShipBetween(userLoginInfo, appTempAccessed) == true {
			// There are two priority options: allow priority, do not allow priority
			if rule.Permission == true {
				theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo":   userLoginInfo,
					"appTempAccessed": appTempAccessed,
					"Description":     rule.Description,
				}).Info("UserAuthorityCheck match permission allow.")
				return true
			} else {
				theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo":   userLoginInfo,
					"appTempAccessed": appTempAccessed,
					"Description":     rule.Description,
				}).Info("UserAuthorityCheck match permission  not allowed.")
				return false
			}
		}
	}

	// There are two default options here: default allowed, default not allowed
	// We choose default not allowed
	theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"UserLoginInfo":   userLoginInfo,
		"appTempAccessed": appTempAccessed,
	}).Info("no permission allow matched")
	return false
}
