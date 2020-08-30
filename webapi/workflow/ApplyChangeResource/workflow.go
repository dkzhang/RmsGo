package ApplyChangeResource

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/myUtils/webapiError"
	"github.com/dkzhang/RmsGo/webapi/authority/authApplication"
	"github.com/dkzhang/RmsGo/webapi/authority/authProject"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/generalForm"
	"github.com/dkzhang/RmsGo/webapi/model/gfApplication"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/dkzhang/RmsGo/webapi/workflow"
	"github.com/sirupsen/logrus"
)

type Workflow struct {
	adm       applicationDM.ApplicationDM
	pdm       projectDM.ProjectDM
	theLogMap logMap.LogMap
}

func NewWorkflow(adm applicationDM.ApplicationDM, pdm projectDM.ProjectDM, lm logMap.LogMap) workflow.GeneralWorkflow {
	wf := Workflow{
		adm:       adm,
		pdm:       pdm,
		theLogMap: lm,
	}
	applyMap := make(map[workflow.KeyTSRA]workflow.ApplyFunc)
	processMap := make(map[workflow.KeyTSRA]workflow.ProcessFunc)

	// 项目长首次提交
	applyMap[workflow.KeyTSRA{
		AppType:   application.AppTypeChange,
		AppStatus: 0,
		UserRole:  user.RoleProjectChief,
		Action:    application.AppActionSubmit,
	}] = wf.ProjectChiefApply

	// 审批人通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeChange,
		AppStatus: application.AppStatusApprover,
		UserRole:  user.RoleApprover,
		Action:    application.AppActionPass,
	}] = wf.ApproverProcessPassOrReject

	// 审批人拒绝
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeChange,
		AppStatus: application.AppStatusApprover,
		UserRole:  user.RoleApprover,
		Action:    application.AppActionReject,
	}] = wf.ApproverProcessPassOrReject

	// 项目长重新提交
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeChange,
		AppStatus: application.AppStatusProjectChief,
		UserRole:  user.RoleProjectChief,
		Action:    application.AppActionSubmit,
	}] = wf.ProjectChiefProcessResubmit

	// 调度员通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeChange,
		AppStatus: application.AppStatusController,
		UserRole:  user.RoleController,
		Action:    application.AppActionPass,
	}] = wf.ControllerProcessPass

	// 调度员拒绝
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeChange,
		AppStatus: application.AppStatusController,
		UserRole:  user.RoleController,
		Action:    application.AppActionReject,
	}] = wf.ControllerProcessReject

	return workflow.NewGeneralWorkflow(applyMap, processMap)
}

func (wf Workflow) ProjectChiefApply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err) {

	var app gfApplication.AppResChange
	err := json.Unmarshal(([]byte)(form.BasicContent), &app)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
			"无法解析申请表单的BasicContent的json结构")
	}
	// (1) Query Project Info from DM and Check auth permission
	pi, err := wf.pdm.QueryByID(form.ProjectID)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeNotFound,
			fmt.Sprintf("Query Project Info ByID (id=%d)  error: %v", form.ProjectID, err),
			"查询项目信息失败")
	}

	// ask permission for update project.
	permission := authProject.AuthorityCheck(wf.theLogMap, userInfo, pi, authApplication.OPS_UPDATE)
	//test
	logrus.Infof("permission = %v, basic status = %d", permission, pi.BasicStatus)
	logrus.Infof("project info = %v", pi)

	if permission == false {
		return -1, webapiError.WaErr(webapiError.TypeAuthorityError,
			fmt.Sprintf("AuthorityCheck reject"),
			"当前用户无权访问该项目")
	}
	//compare res alloc
	if app.CpuNodes < pi.CpuNodesAcquired || app.GpuNodes < pi.GpuNodesAcquired || app.StorageSize < pi.StorageSizeAcquired {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("Res after change master greater than acquired: Expect CPU %d>%d, GPU %d>%d, Storage %d>%d",
				app.CpuNodes, pi.CpuNodesAcquired, app.GpuNodes, pi.GpuNodesAcquired, app.StorageSize, pi.StorageSizeAcquired),
			"变更后的资源数量应不少于已获得的资源数量")
	}

	// (2) Insert New Application
	theApplication := application.Application{
		//ApplicationID:            0,
		ProjectID:                form.ProjectID,
		Type:                     form.Type,
		Status:                   application.AppStatusApprover,
		ApplicantUserID:          userInfo.UserID,
		ApplicantUserChineseName: userInfo.ChineseName,
		DepartmentCode:           userInfo.DepartmentCode,
		BasicContent:             form.BasicContent,
		ExtraContent:             form.ExtraContent,
	}

	appID, err = wf.adm.Insert(theApplication)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation Insert error: %v", err),
			"在数据库中新建申请单记录失败")
	}

	// (3) Insert New ApplicationOps
	theAppOpsRecord := application.AppOpsRecord{
		//RecordID:           0,
		ProjectID:          form.ProjectID,
		ApplicationID:      appID,
		OpsUserID:          userInfo.UserID,
		OpsUserChineseName: userInfo.ChineseName,
		Action:             form.Action,
		ActionStr:          "首次提交",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
	}

	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps error: %v", err),
			"在数据库中新建申请单操作记录失败")
	}
	return appID, nil
}

func (wf Workflow) ApproverProcessPassOrReject(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	theProject, err := wf.pdm.QueryByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}

	// Check Action value
	switch form.Action {
	case 1:
		// Insert New ApplicationOps
		theAppOpsRecord := application.AppOpsRecord{
			//RecordID:           0,
			ProjectID:          theProject.ProjectID,
			ApplicationID:      app.ApplicationID,
			OpsUserID:          userInfo.UserID,
			OpsUserChineseName: userInfo.ChineseName,
			Action:             form.Action,
			ActionStr:          "是",
			BasicInfo:          form.BasicContent,
			ExtraInfo:          form.ExtraContent,
		}
		_, err := wf.adm.InsertAppOps(theAppOpsRecord)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("database operation InsertApplicationOps for Approver error: %v", err),
				"无法为审批人在数据库中新建申请单操作记录")
		}

		// Update Application
		app.Status = application.AppStatusController
		err = wf.adm.Update(app)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("Update for Approver error: %v", err),
				"无法为审批人在数据库中更新Application")
		}
	case -1:
		// Insert New ApplicationOps
		theAppOpsRecord := application.AppOpsRecord{
			//RecordID:           0,
			ProjectID:          theProject.ProjectID,
			ApplicationID:      app.ApplicationID,
			OpsUserID:          userInfo.UserID,
			OpsUserChineseName: userInfo.ChineseName,
			Action:             form.Action,
			ActionStr:          "否",
			BasicInfo:          form.BasicContent,
			ExtraInfo:          form.ExtraContent,
		}
		_, err := wf.adm.InsertAppOps(theAppOpsRecord)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("database operation InsertApplicationOps for Approver error: %v", err),
				"无法为审批人在数据库中新建申请单操作记录")
		}

		// Update Application
		app.Status = application.AppStatusProjectChief
		err = wf.adm.Update(app)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("Update for Approver error: %v", err),
				"无法为审批人在数据库中更新Application")
		}

	default:
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("unsupported action value: %d", form.Action),
			"不支持此action值")
	}

	return nil
}

func (wf Workflow) ProjectChiefProcessResubmit(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	var appRC gfApplication.AppResChange
	err := json.Unmarshal(([]byte)(form.BasicContent), &app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
			"无法解析申请表单的BasicContent的json结构")
	}
	// (1) Query Project Info from DM and Check auth permission
	pi, err := wf.pdm.QueryByID(form.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeNotFound,
			fmt.Sprintf("Query Project Info ByID (id=%d)  error: %v", form.ProjectID, err),
			"查询项目信息失败")
	}

	// ask permission for update project.
	permission := authProject.AuthorityCheck(wf.theLogMap, userInfo, pi, authApplication.OPS_UPDATE)
	//test
	logrus.Infof("permission = %v, basic status = %d", permission, pi.BasicStatus)
	logrus.Infof("project info = %v", pi)

	if permission == false {
		return webapiError.WaErr(webapiError.TypeAuthorityError,
			fmt.Sprintf("AuthorityCheck reject"),
			"当前用户无权访问该项目")
	}
	//compare res alloc
	if appRC.CpuNodes < pi.CpuNodesAcquired || appRC.GpuNodes < pi.GpuNodesAcquired || appRC.StorageSize < pi.StorageSizeAcquired {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("Res after change master greater than acquired: Expect CPU %d>%d, GPU %d>%d, Storage %d>%d",
				appRC.CpuNodes, pi.CpuNodesAcquired, appRC.GpuNodes, pi.GpuNodesAcquired, appRC.StorageSize, pi.StorageSizeAcquired),
			"变更后的资源数量应不少于已获得的资源数量")
	}

	// Insert New ApplicationOps
	theAppOpsRecord := application.AppOpsRecord{
		//RecordID:           0,
		ProjectID:          pi.ProjectID,
		ApplicationID:      app.ApplicationID,
		OpsUserID:          userInfo.UserID,
		OpsUserChineseName: userInfo.ChineseName,
		Action:             form.Action,
		ActionStr:          "重新提交",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps for ProjectChief error: %v", err),
			"无法为项目长在数据库中新建申请单操作记录")
	}

	// Update Application
	app.Status = application.AppStatusApprover
	app.BasicContent = form.BasicContent
	app.ExtraContent = form.ExtraContent

	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for ProjectChief error: %v", err),
			"无法为项目长在数据库中更新Application")
	}

	return nil
}

func (wf Workflow) ControllerProcessReject(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	theProject, err := wf.pdm.QueryByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}

	var appCtrlProjectInfo gfApplication.AppCtrlProjectInfo
	err = json.Unmarshal([]byte(form.BasicContent), &appCtrlProjectInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppCtrlProjectInfo error: %v", err),
			"无法解析form.BasicContent的结构")
	}

	// Check Action value
	// Reject

	// Insert New ApplicationOps
	theAppOpsRecord := application.AppOpsRecord{
		//RecordID:           0,
		ProjectID:          theProject.ProjectID,
		ApplicationID:      app.ApplicationID,
		OpsUserID:          userInfo.UserID,
		OpsUserChineseName: userInfo.ChineseName,
		Action:             form.Action,
		ActionStr:          "否",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps for Controller error: %v", err),
			"无法为调度员在数据库中新建申请单操作记录")
	}

	// Update Application
	app.Status = application.AppStatusProjectChief
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Application")
	}

	return nil
}

func (wf Workflow) ControllerProcessPass(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	theProject, err := wf.pdm.QueryByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}

	//var appCtrlProjectInfo gfApplication.AppCtrlProjectInfo
	//err = json.Unmarshal([]byte(form.BasicContent), &appCtrlProjectInfo)
	//if err != nil {
	//	return webapiError.WaErr(webapiError.TypeBadRequest,
	//		fmt.Sprintf("json Unmarshal to AppCtrlProjectInfo error: %v", err),
	//		"无法解析form.BasicContent的结构")
	//}

	// Insert New ApplicationOps
	theAppOpsRecord := application.AppOpsRecord{
		//RecordID:           0,
		ProjectID:          theProject.ProjectID,
		ApplicationID:      app.ApplicationID,
		OpsUserID:          userInfo.UserID,
		OpsUserChineseName: userInfo.ChineseName,
		Action:             form.Action,
		ActionStr:          "是",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps for Controller error: %v", err),
			"无法为调度员在数据库中新建申请单操作记录")
	}

	// Update Application
	app.Status = application.AppStatusArchived
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Application")
	}

	// Update Project
	var appRC gfApplication.AppResChange
	err = json.Unmarshal(([]byte)(app.BasicContent), &app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
			"无法解析申请表单的BasicContent的json结构")
	}

	ai := project.ApplyInfo{
		ProjectID:           theProject.ProjectID,
		StartDate:           theProject.StartDate,
		TotalDaysApply:      theProject.TotalDaysApply + int(appRC.EndDate.Sub(theProject.EndReminderAt).Hours()/24),
		EndReminderAt:       appRC.EndDate,
		CpuNodesExpected:    appRC.CpuNodes,
		GpuNodesExpected:    appRC.GpuNodes,
		StorageSizeExpected: appRC.StorageSize,
	}

	err = wf.pdm.UpdateApplyInfo(ai)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Project项目变更信息")
	}

	return nil
}
