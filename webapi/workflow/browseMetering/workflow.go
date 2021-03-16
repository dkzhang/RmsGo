package browseMetering

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/gRpcService/client"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
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
	mc  client.MeteringClient
}

func NewWorkflow(adm applicationDM.ApplicationDM, pdm projectDM.ProjectDM, mc client.MeteringClient) workflow.GeneralWorkflow {
	wf := Workflow{
		adm: adm,
		pdm: pdm,
		mc:  mc,
	}
	applyMap := make(map[workflow.KeyTSRA]workflow.ApplyFunc)
	processMap := make(map[workflow.KeyTSRA]workflow.ProcessFunc)

	// 系统自动发起
	applyMap[workflow.KeyTSRA{
		AppType:   application.AppTypeBrowseMetering,
		AppStatus: 0,
		UserRole:  user.RoleSystem,
		Action:    application.AppActionSubmit,
	}] = wf.SystemApply

	// 项目长通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeBrowseMetering,
		AppStatus: application.AppStatusProjectChief,
		UserRole:  user.RoleProjectChief,
		Action:    application.AppActionPass,
	}] = wf.ProjectChiefProcessPass

	// 审批人通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeBrowseMetering,
		AppStatus: application.AppStatusApprover,
		UserRole:  user.RoleApprover,
		Action:    application.AppActionPass,
	}] = wf.ApproverProcessPass

	// 2级审批人通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeBrowseMetering,
		AppStatus: application.AppStatusApprover2,
		UserRole:  user.RoleApprover2,
		Action:    application.AppActionPass,
	}] = wf.Approver2ProcessPass

	// 调度员通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeBrowseMetering,
		AppStatus: application.AppStatusController,
		UserRole:  user.RoleController,
		Action:    application.AppActionPass,
	}] = wf.ControllerProcessPass

	return workflow.NewGeneralWorkflow(applyMap, processMap)
}

func (wf Workflow) SystemApply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err) {
	projectID := form.ProjectID

	pi, err := wf.pdm.QueryByID(projectID)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}
	piJson, _ := json.Marshal(pi)

	// (1) Query Project Metering Info from gRpc client
	ms, err := wf.mc.QueryMetering(projectID, metering.TypeSettlement, "")
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeNotFound,
			fmt.Sprintf("QueryMetering (id=%d)  error: %v", projectID, err),
			"查询项目结算信息失败")
	}

	msJson, err := json.Marshal(ms)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeNotFound,
			fmt.Sprintf("json.Marshal Metering error: ms=%v, err=%v", ms, err),
			"json转换项目结算信息失败")
	}

	// (2) Insert New Application
	theApplication := application.Application{
		//ApplicationID:            0,
		ProjectID:                projectID,
		Type:                     application.AppTypeBrowseMetering,
		Status:                   application.AppStatusProjectChief,
		ApplicantUserID:          userInfo.UserID,
		ApplicantUserChineseName: userInfo.ChineseName,
		DepartmentCode:           userInfo.DepartmentCode,
		BasicContent:             string(msJson),
		ExtraContent:             string(piJson),
	}

	appID, err = wf.adm.Insert(theApplication)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation Insert error: %v", err),
			"在数据库中新建结算单传阅记录失败")
	}

	// (3) Insert New ApplicationOps
	theAppOpsRecord := application.AppOpsRecord{
		//RecordID:           0,
		ProjectID:          projectID,
		ApplicationID:      appID,
		OpsUserID:          userInfo.UserID,
		OpsUserChineseName: userInfo.ChineseName,
		Action:             application.AppActionSubmit,
		ActionStr:          "系统自动生成并发起",
		BasicInfo:          string(msJson),
		ExtraInfo:          "",
	}

	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps error: %v", err),
			"在数据库中新建申请单操作记录失败")
	}
	return appID, nil
}

func (wf Workflow) ProjectChiefProcessPass(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {

	theProject, err := wf.pdm.QueryByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}

	var approvalInfo gfApplication.ApprovalInfo
	err = json.Unmarshal([]byte(form.BasicContent), &approvalInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to ApprovalInfo error: %v", err),
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
		ActionStr:          "是",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps for Approver error: %v", err),
			"无法为项目长在数据库中新建申请单操作记录")
	}

	// Update Application
	app.Status = application.AppStatusApprover
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Approver error: %v", err),
			"无法为项目长在数据库中更新Application")
	}

	return nil
}

func (wf Workflow) ApproverProcessPass(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {

	theProject, err := wf.pdm.QueryByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}

	var approvalInfo gfApplication.ApprovalInfo
	err = json.Unmarshal([]byte(form.BasicContent), &approvalInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to ApprovalInfo error: %v", err),
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
		ActionStr:          "是",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps for Approver error: %v", err),
			"无法为审批人在数据库中新建申请单操作记录")
	}

	// Update Application
	app.Status = application.AppStatusApprover2
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Approver error: %v", err),
			"无法为审批人在数据库中更新Application")
	}

	return nil
}

func (wf Workflow) Approver2ProcessPass(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {

	theProject, err := wf.pdm.QueryByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}

	var approvalInfo gfApplication.ApprovalInfo
	err = json.Unmarshal([]byte(form.BasicContent), &approvalInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to ApprovalInfo error: %v", err),
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
		ActionStr:          "是",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
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

	return nil
}

func (wf Workflow) ControllerProcessPass(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {

	theProject, err := wf.pdm.QueryByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目信息失败")
	}

	var approvalInfo gfApplication.ApprovalInfo
	err = json.Unmarshal([]byte(form.BasicContent), &approvalInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to ApprovalInfo error: %v", err),
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
		ActionStr:          "是",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps for Approver error: %v", err),
			"无法为调度员在数据库中新建申请单操作记录")
	}

	err = wf.pdm.UpdateStatusInfo(project.StatusInfo{
		ProjectID:   theProject.ProjectID,
		BasicStatus: project.BasicStatusArchived,
	})
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation UpdateStatusInfo project.BasicStatusArchived error: %v", err),
			"无法在数据库中将项目状态置为已归档")
	}

	// Update Application
	app.Status = application.AppStatusArchived
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Approver error: %v", err),
			"无法为调度员在数据库中更新Application")
	}

	return nil
}
