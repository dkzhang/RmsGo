package workflow

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/webapiError"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/generalForm"
	"github.com/dkzhang/RmsGo/webapi/model/user"
)

//type GeneralWorkflow interface {
//	Apply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err)
//	Process(form generalForm.GeneralForm, userInfo user.UserInfo) (waErr webapiError.Err)
//}
type ApplyFunc func(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err)
type ProcessFunc func(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err)

type GeneralWorkflow struct {
	applyMap   map[KeySRA]ApplyFunc
	processMap map[KeySRA]ProcessFunc
}

func NewGeneralWorkflow(am map[KeySRA]ApplyFunc, pm map[KeySRA]ProcessFunc) GeneralWorkflow {
	return GeneralWorkflow{
		applyMap:   am,
		processMap: pm,
	}
}

func (gwf GeneralWorkflow) Apply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err) {
	k := KeySRA{
		AppType:   form.Type,
		AppStatus: 0,
		UserRole:  userInfo.Role,
		Action:    form.Action,
	}
	execFunc, ok := gwf.applyMap[k]
	if !ok {
		return -1, webapiError.WaErr(webapiError.TypeAuthorityError,
			fmt.Sprintf("apply application (key: %v) does not allowed", k),
			"该用户无权对该申请单进行该操作")
	}
	return execFunc(form, userInfo)
}

func (gwf GeneralWorkflow) Process(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	k := KeySRA{
		AppType:   form.Type,
		AppStatus: app.Status,
		UserRole:  userInfo.Role,
		Action:    form.Action,
	}
	execFunc, ok := gwf.processMap[k]
	if !ok {
		return webapiError.WaErr(webapiError.TypeAuthorityError,
			fmt.Sprintf("apply application (key: %v) does not allowed", k),
			"该用户无权对该申请单进行该操作")
	}
	return execFunc(form, app, userInfo)
}

type KeySRA struct {
	AppType   int
	AppStatus int
	UserRole  int
	Action    int
}
