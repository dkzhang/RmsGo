package ApplyProjectAndResource

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/webapiError"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/generalForm"
	"github.com/dkzhang/RmsGo/webapi/model/gfApplication"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/dkzhang/RmsGo/webapi/workflow"
)

type Workflow struct {
	adm applicationDM.ApplicationDM
	pdm projectDM.ProjectDM
}

func NewWorkflow(adm applicationDM.ApplicationDM, pdm projectDM.ProjectDM) workflow.GeneralWorkflow {
	wf := Workflow{
		adm: adm,
		pdm: pdm,
	}
	applyMap := make(map[workflow.KeyTSRA]workflow.ApplyFunc)
	processMap := make(map[workflow.KeyTSRA]workflow.ProcessFunc)

	// 项目长首次提交
	applyMap[workflow.KeyTSRA{
		AppType:   application.AppTypeNew,
		AppStatus: 0,
		UserRole:  user.RoleProjectChief,
		Action:    application.AppActionSubmit,
	}] = wf.ProjectChiefApply

	// 审批人通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeNew,
		AppStatus: application.AppStatusApprover,
		UserRole:  user.RoleApprover,
		Action:    application.AppActionPass,
	}] = wf.ApproverProcessPassOrReject

	// 审批人拒绝
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeNew,
		AppStatus: application.AppStatusApprover,
		UserRole:  user.RoleApprover,
		Action:    application.AppActionReject,
	}] = wf.ApproverProcessPassOrReject

	// 项目长重新提交
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeNew,
		AppStatus: application.AppStatusProjectChief,
		UserRole:  user.RoleProjectChief,
		Action:    application.AppActionSubmit,
	}] = wf.ProjectChiefProcessResubmit

	// 调度员通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeNew,
		AppStatus: application.AppStatusController,
		UserRole:  user.RoleController,
		Action:    application.AppActionPass,
	}] = wf.ControllerProcessPass

	// 调度员拒绝
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeNew,
		AppStatus: application.AppStatusController,
		UserRole:  user.RoleController,
		Action:    application.AppActionReject,
	}] = wf.ControllerProcessReject

	return workflow.NewGeneralWorkflow(applyMap, processMap)
}

func (wf Workflow) ProjectChiefApply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err) {

	var app gfApplication.AppNewProRes
	err := json.Unmarshal(([]byte)(form.BasicContent), &app)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
			"无法解析申请表单的BasicContent的json结构")
	}

	// (1) Insert New Project
	theProject := project.Info{
		//ProjectID:            0,
		ProjectName: app.ProjectName,
		//ProjectCode:          "",
		DepartmentCode:   userInfo.DepartmentCode,
		Department:       userInfo.Department,
		ChiefID:          userInfo.UserID,
		ChiefChineseName: userInfo.ChineseName,
		ExtraInfo:        form.ExtraContent,

		BasicStatus:          project.BasicStatusApplying,
		ComputingAllocStatus: project.ResNotYetAssigned,
		StorageAllocStatus:   project.ResNotYetAssigned,

		StartDate:           app.StartDate,
		TotalDaysApply:      app.TotalDaysApply,
		EndReminderAt:       app.EndDate,
		CpuNodesExpected:    0,
		GpuNodesExpected:    0,
		StorageSizeExpected: 0,

		CpuNodesAcquired:    0,
		GpuNodesAcquired:    0,
		StorageSizeAcquired: 0,
		CpuNodesMap:         "",
		GpuNodesMap:         "",
		StorageAllocInfo:    "",
	}

	projectID, err := wf.pdm.Insert(theProject)

	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("insert project info error: %v", err),
			"在数据库中新建项目记录失败")
	}

	// (2) Insert New Application
	theApplication := application.Application{
		//ApplicationID:            0,
		ProjectID:                projectID,
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
		ProjectID:          projectID,
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
	theProject, err := wf.pdm.QueryByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}

	var appNewProRes gfApplication.AppNewProRes
	err = json.Unmarshal([]byte(form.BasicContent), &appNewProRes)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
			"无法解析form.BasicContent的结构")
	}

	// Insert New ApplicationOps
	theAppOpsRecord := application.AppOpsRecord{
		//RecordID:           0,
		ProjectID:          theProject.ProjectID,
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

	// Update Project
	bi := project.BasicInfo{
		ProjectID:   theProject.ProjectID,
		ProjectName: appNewProRes.ProjectName,
		ExtraInfo:   form.ExtraContent,
	}
	err = wf.pdm.UpdateBasicInfo(bi)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for ProjectChief error: %v", err),
			"无法为项目长在数据库中更新Project基本信息")
	}

	ai := project.ApplyInfo{
		ProjectID:           theProject.ProjectID,
		StartDate:           appNewProRes.StartDate,
		TotalDaysApply:      appNewProRes.TotalDaysApply,
		EndReminderAt:       appNewProRes.EndDate,
		CpuNodesExpected:    appNewProRes.CpuNodes,
		GpuNodesExpected:    appNewProRes.GpuNodes,
		StorageSizeExpected: appNewProRes.StorageSize,
	}

	err = wf.pdm.UpdateApplyInfo(ai)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for ProjectChief error: %v", err),
			"无法为项目长在数据库中更新Project申请信息")
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

	var appCtrlProjectInfo gfApplication.CtrlApprovalInfo
	err = json.Unmarshal([]byte(form.BasicContent), &appCtrlProjectInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to CtrlApprovalInfo error: %v", err),
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

	var appCtrlProjectInfo gfApplication.CtrlApprovalInfo
	err = json.Unmarshal([]byte(form.BasicContent), &appCtrlProjectInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to CtrlApprovalInfo error: %v", err),
			"无法解析form.BasicContent的结构")
	}

	// Check Action value
	// Pass

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
	ci := project.CodeInfo{
		ProjectID:   theProject.ProjectID,
		ProjectCode: appCtrlProjectInfo.ProjectCode,
	}
	err = wf.pdm.UpdateCodeInfo(ci)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Project项目编码信息")
	}

	si := project.StatusInfo{
		ProjectID:            theProject.ProjectID,
		BasicStatus:          project.BasicStatusEstablished,
		ComputingAllocStatus: project.ResNotYetAssigned,
		StorageAllocStatus:   project.ResNotYetAssigned,
	}
	err = wf.pdm.UpdateStatusInfo(si)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Project状态信息")
	}
	return nil
}
