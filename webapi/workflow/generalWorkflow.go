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

type ApplyHookFunc func(form generalForm.GeneralForm, userInfo user.UserInfo, appID int, waErr webapiError.Err)
type ProcessHookFunc func(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo, waErr webapiError.Err)

type GeneralWorkflow struct {
	applyMap   map[KeyTSRA]ApplyFunc
	processMap map[KeyTSRA]ProcessFunc

	applyHook   map[KeyTSRA]([]ApplyHookFunc)
	processHook map[KeyTSRA]([]ProcessHookFunc)
}

func NewGeneralWorkflow(am map[KeyTSRA]ApplyFunc, pm map[KeyTSRA]ProcessFunc) GeneralWorkflow {
	return GeneralWorkflow{
		applyMap:    am,
		processMap:  pm,
		applyHook:   make(map[KeyTSRA]([]ApplyHookFunc)),
		processHook: make(map[KeyTSRA]([]ProcessHookFunc)),
	}
}

func (gwf GeneralWorkflow) HookApply(k KeyTSRA, ahf ApplyHookFunc) {
	if _, ok := gwf.applyHook[k]; !ok {
		gwf.applyHook[k] = append(gwf.applyHook[k], ahf)
	}
}

func (gwf GeneralWorkflow) HookProcess(k KeyTSRA, phf ProcessHookFunc) {
	if _, ok := gwf.processHook[k]; !ok {
		gwf.processHook[k] = append(gwf.processHook[k], phf)
	}
}

func (gwf GeneralWorkflow) Apply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err) {
	k := KeyTSRA{
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

	appID, waErr = execFunc(form, userInfo)

	// Hook
	if hooks, ok := gwf.applyHook[k]; ok {
		for _, h := range hooks {
			h(form, userInfo, appID, waErr)
		}
	}

	return appID, waErr
}

func (gwf GeneralWorkflow) Process(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	k := KeyTSRA{
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

	waErr = execFunc(form, app, userInfo)

	// Hook
	if hooks, ok := gwf.processHook[k]; ok {
		for _, h := range hooks {
			h(form, app, userInfo, waErr)
		}
	}
	return waErr
}

type KeyTSRA struct {
	AppType   int
	AppStatus int
	UserRole  int
	Action    int
}
