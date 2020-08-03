package workflow

import (
	"github.com/dkzhang/RmsGo/webapi/model/generalForm"
	"github.com/dkzhang/RmsGo/webapi/model/user"
)

type GeneralWorkflow interface {
	Apply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, err error, msg string)
	Process(form generalForm.GeneralForm, userInfo user.UserInfo) (err error, msg string)
}
