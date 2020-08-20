package authApplication

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/user"
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
}

func AuthorityCheck(theLogMap logMap.LogMap,
	userLoginInfo user.UserInfo,
	app application.Application,
	ops int) (permission bool) {

	return false
}
