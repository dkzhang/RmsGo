package ApplyProjectAndResource

import (
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
	"gopkg.in/yaml.v2"
	"time"
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
	err := yaml.Unmarshal([]byte(form.BasicContent), &app)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
			"无法解析申请表单的BasicContent的json结构")
	}

	// (1) Insert New Project
	theProjectS := project.StaticInfo{
		//ProjectID:        0,
		ProjectName: app.ProjectName,
		//ProjectCode:      "",
		DepartmentCode:   userInfo.DepartmentCode,
		Department:       userInfo.Department,
		ChiefID:          userInfo.UserID,
		ChiefChineseName: userInfo.ChineseName,
		ExtraInfo:        form.ExtraContent,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	theProjectD := project.DynamicInfo{
		//ProjectID:               0,
		BasicStatus:          project.BasicStatusApplying,
		ComputingAllocStatus: project.ResNotYetAssigned,
		StorageAllocStatus:   project.ResNotYetAssigned,
		StartDate:            app.StartDate,
		TotalDaysApply:       app.TotalDaysApply,
		CpuNodesExpected:     0,
		GpuNodesExpected:     0,
		StorageSizeExpected:  0,
		CpuNodesAcquired:     0,
		GpuNodesAcquired:     0,
		StorageSizeAcquired:  0,
		EndReminderAt:        app.EndDate,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	projectID, err := wf.pdm.InsertAllInfo(project.ProjectInfo{
		ProjectID:      theProjectS.ProjectID,
		TheStaticInfo:  theProjectS,
		TheDynamicInfo: theProjectD,
	})

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
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
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
		CreatedAt:          time.Now(),
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
	theProjectS, err := wf.pdm.QueryStaticInfoByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目静态信息失败")
	}

	// Check Action value
	switch form.Action {
	case 1:
		// Insert New ApplicationOps
		theAppOpsRecord := application.AppOpsRecord{
			//RecordID:           0,
			ProjectID:          theProjectS.ProjectID,
			ApplicationID:      app.ApplicationID,
			OpsUserID:          userInfo.UserID,
			OpsUserChineseName: userInfo.ChineseName,
			Action:             form.Action,
			ActionStr:          "是",
			BasicInfo:          form.BasicContent,
			ExtraInfo:          form.ExtraContent,
			CreatedAt:          time.Now(),
		}
		_, err := wf.adm.InsertAppOps(theAppOpsRecord)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("database operation InsertApplicationOps for Approver error: %v", err),
				"无法为审批人在数据库中新建申请单操作记录")
		}

		// Update Application
		app.Status = application.AppStatusController
		app.UpdatedAt = time.Now()
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
			ProjectID:          theProjectS.ProjectID,
			ApplicationID:      app.ApplicationID,
			OpsUserID:          userInfo.UserID,
			OpsUserChineseName: userInfo.ChineseName,
			Action:             form.Action,
			ActionStr:          "否",
			BasicInfo:          form.BasicContent,
			ExtraInfo:          form.ExtraContent,
			CreatedAt:          time.Now(),
		}
		_, err := wf.adm.InsertAppOps(theAppOpsRecord)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("database operation InsertApplicationOps for Approver error: %v", err),
				"无法为审批人在数据库中新建申请单操作记录")
		}

		// Update Application
		app.Status = application.AppStatusProjectChief
		app.UpdatedAt = time.Now()
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
	theProjectS, err := wf.pdm.QueryStaticInfoByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目静态信息失败")
	}

	theProjectD, err := wf.pdm.QueryDynamicInfoByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目动态信息失败")
	}

	var appNewProRes gfApplication.AppNewProRes
	err = yaml.Unmarshal([]byte(form.BasicContent), &appNewProRes)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
			"无法解析form.BasicContent的结构")
	}

	// Insert New ApplicationOps
	theAppOpsRecord := application.AppOpsRecord{
		//RecordID:           0,
		ProjectID:          theProjectS.ProjectID,
		ApplicationID:      app.ApplicationID,
		OpsUserID:          userInfo.UserID,
		OpsUserChineseName: userInfo.ChineseName,
		Action:             form.Action,
		ActionStr:          "重新提交",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
		CreatedAt:          time.Now(),
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
	app.UpdatedAt = time.Now()
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for ProjectChief error: %v", err),
			"无法为项目长在数据库中更新Application")
	}

	// Update Project
	theProjectS.ProjectName = appNewProRes.ProjectName
	theProjectS.ExtraInfo = form.ExtraContent
	err = wf.pdm.UpdateStaticInfo(theProjectS)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for ProjectChief error: %v", err),
			"无法为项目长在数据库中更新Project静态信息")
	}

	theProjectD.StartDate = appNewProRes.StartDate
	theProjectD.TotalDaysApply = appNewProRes.TotalDaysApply
	theProjectD.EndReminderAt = appNewProRes.EndDate
	err = wf.pdm.UpdateDynamicInfo(theProjectD)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for ProjectChief error: %v", err),
			"无法为项目长在数据库中更新Project动态信息")
	}

	return nil
}

func (wf Workflow) ControllerProcessReject(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	theProjectS, err := wf.pdm.QueryStaticInfoByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目静态信息失败")
	}

	var appCtrlProjectInfo gfApplication.AppCtrlProjectInfo
	err = yaml.Unmarshal([]byte(form.BasicContent), &appCtrlProjectInfo)
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
		ProjectID:          theProjectS.ProjectID,
		ApplicationID:      app.ApplicationID,
		OpsUserID:          userInfo.UserID,
		OpsUserChineseName: userInfo.ChineseName,
		Action:             form.Action,
		ActionStr:          "否",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
		CreatedAt:          time.Now(),
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps for Controller error: %v", err),
			"无法为调度员在数据库中新建申请单操作记录")
	}

	// Update Application
	app.Status = application.AppStatusProjectChief
	app.UpdatedAt = time.Now()
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Application")
	}

	return nil
}

func (wf Workflow) ControllerProcessPass(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	theProjectS, err := wf.pdm.QueryStaticInfoByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目静态信息失败")
	}

	theProjectD, err := wf.pdm.QueryDynamicInfoByID(app.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryByID error: %v", err),
			"在数据库中查询项目动态信息失败")
	}

	var appCtrlProjectInfo gfApplication.AppCtrlProjectInfo
	err = yaml.Unmarshal([]byte(form.BasicContent), &appCtrlProjectInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppCtrlProjectInfo error: %v", err),
			"无法解析form.BasicContent的结构")
	}

	// Check Action value
	// Pass

	// Insert New ApplicationOps
	theAppOpsRecord := application.AppOpsRecord{
		//RecordID:           0,
		ProjectID:          theProjectS.ProjectID,
		ApplicationID:      app.ApplicationID,
		OpsUserID:          userInfo.UserID,
		OpsUserChineseName: userInfo.ChineseName,
		Action:             form.Action,
		ActionStr:          "是",
		BasicInfo:          form.BasicContent,
		ExtraInfo:          form.ExtraContent,
		CreatedAt:          time.Now(),
	}
	_, err = wf.adm.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps for Controller error: %v", err),
			"无法为调度员在数据库中新建申请单操作记录")
	}

	// Update Application
	app.Status = application.AppStatusArchived
	app.UpdatedAt = time.Now()
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Application")
	}

	// Update Project
	theProjectS.ProjectName = appCtrlProjectInfo.ProjectCode
	err = wf.pdm.UpdateStaticInfo(theProjectS)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Project静态信息")
	}

	theProjectD.BasicStatus = project.BasicStatusEstablished

	// application passed
	var appNewProRes gfApplication.AppNewProRes
	err = yaml.Unmarshal([]byte(app.BasicContent), &appNewProRes)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeServerInternalError,
			fmt.Sprintf("app.BasicContent json Unmarshal to AppNewProRes error: %v", err),
			"无法解析form.BasicContent的结构")
	}
	theProjectD.CpuNodesExpected = appNewProRes.CpuNodes
	theProjectD.GpuNodesExpected = appNewProRes.GpuNodes
	theProjectD.StorageSizeExpected = appNewProRes.StorageSize

	err = wf.pdm.UpdateDynamicInfo(theProjectD)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Project动态信息")
	}

	return nil
}
