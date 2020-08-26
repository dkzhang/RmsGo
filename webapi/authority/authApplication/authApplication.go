package authApplication

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/sirupsen/logrus"
)

type applicationAuthority struct {
	RelationShipBetween func(user.UserInfo, application.Application) bool
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

var theApplicationAuthorityTable []applicationAuthority

func init() {
	theApplicationAuthorityTable = make([]applicationAuthority, 0)

	theApplicationAuthorityTable = append(theApplicationAuthorityTable, applicationAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, applicationAccessed application.Application) bool {
			if userLoginInfo.Role == user.RoleProjectChief {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_CREATE,
		Permission:  true,
		Description: "Allow RoleProjectChief CREATE Application",
	})

	theApplicationAuthorityTable = append(theApplicationAuthorityTable, applicationAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, applicationAccessed application.Application) bool {
			if userLoginInfo.Role == user.RoleProjectChief && applicationAccessed.ApplicantUserID == userLoginInfo.UserID {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_RETRIEVE | OPS_UPDATE,
		Permission:  true,
		Description: "Allow RoleProjectChief RETRIEVE and UPDATE Application owned by himself",
	})

	theApplicationAuthorityTable = append(theApplicationAuthorityTable, applicationAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, applicationAccessed application.Application) bool {
			if userLoginInfo.Role == user.RoleApprover && applicationAccessed.DepartmentCode == userLoginInfo.DepartmentCode {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_RETRIEVE | OPS_UPDATE,
		Permission:  true,
		Description: "Allow RoleProjectChief RETRIEVE and UPDATE Application",
	})

	theApplicationAuthorityTable = append(theApplicationAuthorityTable, applicationAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, applicationAccessed application.Application) bool {
			if userLoginInfo.Role == user.RoleController {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_RETRIEVE | OPS_UPDATE,
		Permission:  true,
		Description: "Allow RoleProjectChief RETRIEVE and UPDATE Application",
	})
}

func AuthorityCheck(theLogMap logMap.LogMap,
	userLoginInfo user.UserInfo,
	app application.Application,
	ops int) (permission bool) {

	for _, rule := range theApplicationAuthorityTable {
		if (ops&rule.Operation != 0) && rule.RelationShipBetween(userLoginInfo, app) == true {
			// There are two priority options: allow priority, do not allow priority
			if rule.Permission == true {
				theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo": userLoginInfo,
					"Application":   app,
					"Description":   rule.Description,
				}).Info("UserAuthorityCheck match permission allow.")
				return true
			} else {
				theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo": userLoginInfo,
					"Application":   app,
					"Description":   rule.Description,
				}).Info("UserAuthorityCheck match permission  not allowed.")
				return false
			}
		}
	}

	// There are two default options here: default allowed, default not allowed
	// We choose default not allowed
	theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"UserLoginInfo": userLoginInfo,
		"Application":   app,
	}).Info("no permission allow matched")
	return false
}
