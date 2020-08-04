package workflow

import (
	"github.com/dkzhang/RmsGo/myUtils/webapiError"
	"github.com/dkzhang/RmsGo/webapi/model/generalForm"
	"github.com/dkzhang/RmsGo/webapi/model/user"
)

type GeneralWorkflow interface {
	Apply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err)
	Process(form generalForm.GeneralForm, userInfo user.UserInfo) (waErr webapiError.Err)
}
