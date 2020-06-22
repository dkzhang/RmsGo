package authGeneralFormDraftCRUD

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/model/generalFormDraft"

	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/sirupsen/logrus"
)

type generalFormDraftAuthority struct {
	RelationShipBetween func(user.UserInfo, generalFormDraft.GeneralFormDraft) bool
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

var theGeneralFormDraftAuthorityTable []generalFormDraftAuthority

func init() {
	theGeneralFormDraftAuthorityTable = make([]generalFormDraftAuthority, 0)

	theGeneralFormDraftAuthorityTable = append(theGeneralFormDraftAuthorityTable, generalFormDraftAuthority{
		RelationShipBetween: func(userLoginInfo user.UserInfo, GeneralFormDraftAccessed generalFormDraft.GeneralFormDraft) bool {
			if userLoginInfo.Role == user.RoleProjectChief &&
				GeneralFormDraftAccessed.UserID == userLoginInfo.UserID {
				return true
			} else {
				return false
			}
		},
		Operation:   OPS_CREATE | OPS_RETRIEVE | OPS_UPDATE | OPS_DELETE,
		Permission:  true,
		Description: "Allow RoleProjectChief CRUD GeneralFormDraft owned by himself",
	})
}

func GeneralFormDraftAuthorityCheck(theLogMap logMap.LogMap,
	userLoginInfo user.UserInfo, GeneralFormDraftAccessed generalFormDraft.GeneralFormDraft, ops int) (permission bool) {
	for _, rule := range theGeneralFormDraftAuthorityTable {
		if (ops&rule.Operation != 0) && rule.RelationShipBetween(userLoginInfo, GeneralFormDraftAccessed) == true {
			// There are two priority options: allow priority, do not allow priority
			if rule.Permission == true {
				theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo":            userLoginInfo,
					"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
					"Description":              rule.Description,
				}).Info("UserAuthorityCheck match permission allow.")
				return true
			} else {
				theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
					"UserLoginInfo":            userLoginInfo,
					"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
					"Description":              rule.Description,
				}).Info("UserAuthorityCheck match permission  not allowed.")
				return false
			}
		}
	}

	// There are two default options here: default allowed, default not allowed
	// We choose default not allowed
	theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"UserLoginInfo":            userLoginInfo,
		"GeneralFormDraftAccessed": GeneralFormDraftAccessed,
	}).Info("no permission allow matched")
	return false
}
