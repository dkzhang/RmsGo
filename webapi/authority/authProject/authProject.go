package authProject

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/sirupsen/logrus"
)

type projectAuthority struct {
	RelationShipBetween func(user.UserInfo, project.Info) bool
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

var theProjectAuthorityTable []projectAuthority

func init() {
	theProjectAuthorityTable = make([]projectAuthority, 0)

	theProjectAuthorityTable = append(theProjectAuthorityTable, projectAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, projectAccessed project.Info) bool {
			if userLoginInfo.Role == user.RoleProjectChief &&
				projectAccessed.ChiefID == userLoginInfo.UserID &&
				(projectAccessed.BasicStatus == project.BasicStatusWaiting ||
					projectAccessed.BasicStatus == project.BasicStatusRunning) {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_UPDATE,
		Permission:  true,
		Description: "Allow RoleProjectChief UPDATE Project (in basicStatus Waiting or Running)",
	})

	theProjectAuthorityTable = append(theProjectAuthorityTable, projectAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, projectAccessed project.Info) bool {
			if userLoginInfo.Role == user.RoleProjectChief && projectAccessed.ChiefID == userLoginInfo.UserID {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_RETRIEVE,
		Permission:  true,
		Description: "Allow RoleProjectChief RETRIEVE Project owned by himself",
	})

	theProjectAuthorityTable = append(theProjectAuthorityTable, projectAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, projectAccessed project.Info) bool {
			if userLoginInfo.Role == user.RoleApprover && projectAccessed.DepartmentCode == userLoginInfo.DepartmentCode {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_RETRIEVE,
		Permission:  true,
		Description: "Allow RoleProjectChief RETRIEVE Project",
	})

	theProjectAuthorityTable = append(theProjectAuthorityTable, projectAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, projectAccessed project.Info) bool {
			if userLoginInfo.Role == user.RoleController {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_RETRIEVE | OPS_UPDATE,
		Permission:  true,
		Description: "Allow RoleProjectChief RETRIEVE and Project",
	})
}

func AuthorityCheck(theLogMap logMap.LogMap,
	userLoginInfo user.UserInfo,
	pi project.Info,
	ops int) (permission bool) {

	for _, rule := range theProjectAuthorityTable {
		if (ops&rule.Operation != 0) && rule.RelationShipBetween(userLoginInfo, pi) == true {
			// There are two priority options: allow priority, do not allow priority
			if rule.Permission == true {
				theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo": userLoginInfo,
					"ProjectInfo":   pi,
					"Description":   rule.Description,
				}).Info("UserAuthorityCheck match permission allow.")
				return true
			} else {
				theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo": userLoginInfo,
					"ProjectInfo":   pi,
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
		"ProjectInfo":   pi,
	}).Info("no permission allow matched")
	return false
}
