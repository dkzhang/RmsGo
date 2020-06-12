package userCRUD

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/sirupsen/logrus"
)

type userAuthority struct {
	RelationShipBetween func(user.UserInfo, user.UserInfo) bool
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

var theUserAuthorityTable []userAuthority

func init() {
	theUserAuthorityTable = make([]userAuthority, 0)

	theUserAuthorityTable = append(theUserAuthorityTable, userAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, userAccessedInfo user.UserInfo) bool {
			if userLoginInfo.Role == user.RoleController {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_CREATE | OPS_RETRIEVE | OPS_UPDATE | OPS_DELETE,
		Permission:  true,
		Description: "Allow RoleController to do CRUD ops to all users",
	})

	theUserAuthorityTable = append(theUserAuthorityTable, userAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, userAccessedInfo user.UserInfo) bool {
			if userLoginInfo.Role == user.RoleApprover &&
				userLoginInfo.DepartmentCode == userAccessedInfo.DepartmentCode {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_RETRIEVE | OPS_UPDATE,
		Permission:  true,
		Description: "Allow RoleRoleApprover to do Retrieve & Update ops to all users in the same department",
	})

	theUserAuthorityTable = append(theUserAuthorityTable, userAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, userAccessedInfo user.UserInfo) bool {
			if userLoginInfo.Role == user.RoleApprover &&
				userAccessedInfo.Role == user.RoleProjectChief &&
				userLoginInfo.DepartmentCode == userAccessedInfo.DepartmentCode {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_CREATE | OPS_DELETE,
		Permission:  true,
		Description: "Allow RoleRoleApprover to do Create & Delete ops to all RoleProjectChief in the same department",
	})

	theUserAuthorityTable = append(theUserAuthorityTable, userAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, userAccessedInfo user.UserInfo) bool {
			if userLoginInfo.Role == user.RoleProjectChief &&
				userLoginInfo.UserID == userAccessedInfo.UserID {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_RETRIEVE,
		Permission:  true,
		Description: "Allow RoleProjectChief to do retrieve his own info",
	})
}

func UserAuthorityCheck(userLoginInfo, userAccessedInfo user.UserInfo, ops int) (permission bool) {
	for _, rule := range theUserAuthorityTable {
		if (ops&rule.Operation != 0) && rule.RelationShipBetween(userLoginInfo, userAccessedInfo) == true {
			// There are two priority options: allow priority, do not allow priority
			if rule.Permission == true {
				logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo":    userLoginInfo,
					"UserAccessedInfo": userAccessedInfo,
					"Description":      rule.Description,
				}).Info("UserAuthorityCheck match permission allow.")
				return true
			} else {
				logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo":    userLoginInfo,
					"UserAccessedInfo": userAccessedInfo,
					"Description":      rule.Description,
				}).Info("UserAuthorityCheck match permission  not allowed.")
				return false
			}
		}
	}

	// There are two default options here: default allowed, default not allowed
	// We choose default not allowed
	logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"UserLoginInfo":    userLoginInfo,
		"UserAccessedInfo": userAccessedInfo,
	}).Info("no permission allow matched")
	return false
}
