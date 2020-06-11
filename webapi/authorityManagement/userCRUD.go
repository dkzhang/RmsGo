package authorityManagement

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/sirupsen/logrus"
)

type UserAuthority struct {
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

var TheUserAuthorityTable []UserAuthority

func init() {
	TheUserAuthorityTable = make([]UserAuthority, 0)

	TheUserAuthorityTable = append(TheUserAuthorityTable, UserAuthority{
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

	TheUserAuthorityTable = append(TheUserAuthorityTable, UserAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, userAccessedInfo user.UserInfo) bool {
			if userLoginInfo.Role == user.RoleApprover &&
				userLoginInfo.DepartmentCode == userAccessedInfo.DepartmentCode {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_CREATE | OPS_RETRIEVE | OPS_UPDATE | OPS_DELETE,
		Permission:  true,
		Description: "Allow RoleRoleApprover to do CRUD ops to all users in the same department",
	})

	TheUserAuthorityTable = append(TheUserAuthorityTable, UserAuthority{
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

func UserAuthorityCheck(userLoginID, userAccessedID int, ops int) (userLoginInfo user.UserInfo, userAccessedInfo user.UserInfo,
	permission bool, err error) {
	userLoginInfo, err = webapi.TheInfras.TheUserDM.QueryUserByID(userLoginID)
	if err != nil {
		return user.UserInfo{}, user.UserInfo{},
			false, fmt.Errorf("TheUserDM.QueryUserByID userLoginID error: %v", err)
	}

	userAccessedInfo, err = webapi.TheInfras.TheUserDM.QueryUserByID(userAccessedID)
	if err != nil {
		return user.UserInfo{}, user.UserInfo{},
			false, fmt.Errorf("TheUserDM.QueryUserByID userLoginID error: %v", err)
	}

	for _, rule := range TheUserAuthorityTable {
		if (ops&rule.Operation != 0) && rule.RelationShipBetween(userLoginInfo, userAccessedInfo) == true {
			// There are two priority options: allow priority, do not allow priority
			// We choose do allow priority
			if rule.Permission == true {
				logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo":    userLoginInfo,
					"UserAccessedInfo": userAccessedInfo,
					"Description":      rule.Description,
				}).Info("UserAuthorityCheck match permission.")
				return userLoginInfo, userAccessedInfo, true, nil
			}
		}
	}

	// There are two default options here: default allowed, default not allowed
	// We choose default not allowed, since we choose do allow priority
	logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"UserLoginInfo":    userLoginInfo,
		"UserAccessedInfo": userAccessedInfo,
	}).Info("no permission allow matched")
	return userLoginInfo, userAccessedInfo, false, nil
}
