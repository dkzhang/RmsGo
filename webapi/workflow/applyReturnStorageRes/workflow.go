package applyReturnStorageRes

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/myUtils/webapiError"
	"github.com/dkzhang/RmsGo/webapi/authority/authApplication"
	"github.com/dkzhang/RmsGo/webapi/authority/authProject"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectResDM"
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
	prdm      projectResDM.ProjectResDM
	bmwf      workflow.GeneralWorkflow
	theLogMap logMap.LogMap
}

func NewWorkflow(adm applicationDM.ApplicationDM, pdm projectDM.ProjectDM,
	prdm projectResDM.ProjectResDM, bmwf workflow.GeneralWorkflow,
	lm logMap.LogMap) workflow.GeneralWorkflow {
	wf := Workflow{
		adm:       adm,
		pdm:       pdm,
		prdm:      prdm,
		bmwf:      bmwf,
		theLogMap: lm,
	}
	applyMap := make(map[workflow.KeyTSRA]workflow.ApplyFunc)
	processMap := make(map[workflow.KeyTSRA]workflow.ProcessFunc)

	// 项目长首次提交
	applyMap[workflow.KeyTSRA{
		AppType:   application.AppTypeReturnStorage,
		AppStatus: 0,
		UserRole:  user.RoleProjectChief,
		Action:    application.AppActionSubmit,
	}] = wf.ProjectChiefApply

	// 项目长重新提交
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeReturnStorage,
		AppStatus: application.AppStatusProjectChief,
		UserRole:  user.RoleProjectChief,
		Action:    application.AppActionSubmit,
	}] = wf.ProjectChiefProcessResubmit

	// 调度员通过
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeReturnStorage,
		AppStatus: application.AppStatusController,
		UserRole:  user.RoleController,
		Action:    application.AppActionPass,
	}] = wf.ControllerProcessPass

	// 调度员拒绝
	processMap[workflow.KeyTSRA{
		AppType:   application.AppTypeReturnStorage,
		AppStatus: application.AppStatusController,
		UserRole:  user.RoleController,
		Action:    application.AppActionReject,
	}] = wf.ControllerProcessReject

	return workflow.NewGeneralWorkflow(applyMap, processMap)
}

func (wf Workflow) ProjectChiefApply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err) {
	returnResInfo, err := gfApplication.JsonUnmarshalAppResStoReturn(form.BasicContent)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppResStoReturn error: %v", err),
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

	// check project basic status
	if pi.BasicStatus != project.BasicStatusRunning {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("project BasicStatus check failed: expect to be running but got %d", pi.BasicStatus),
			"当前项目未处于running状态，无法回收资源")
	}

	// check storage and cgpu remained
	if returnResInfo.SizeAfter == 0 {
		if pi.CpuNodesAcquired != 0 || pi.GpuNodesAcquired != 0 {
			return -1, webapiError.WaErr(webapiError.TypeBadRequest,
				fmt.Sprintf("project CPU/GPU Acquired must be all ZERO before return all storage"),
				"归还全部存储资源前，必须先归还全部计算资源")
		}
	}
	if returnResInfo.SizeAfter > pi.StorageSizeAcquired {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("storage size after return must be smaller than currently acquired"),
			"归还后剩余的计算资源必须小于当前已获得的计算资源")
	}

	// (2) Insert New Application
	theApplication := application.Application{
		//ApplicationID:            0,
		ProjectID:                form.ProjectID,
		Type:                     form.Type,
		Status:                   application.AppStatusController,
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

func (wf Workflow) ProjectChiefProcessResubmit(form generalForm.GeneralForm, app application.Application, userInfo user.UserInfo) (waErr webapiError.Err) {
	returnResInfo, err := gfApplication.JsonUnmarshalAppResStoReturn(form.BasicContent)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppResStoReturn error: %v", err),
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

	// check project basic status
	if pi.BasicStatus != project.BasicStatusRunning {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("project BasicStatus check failed: expect to be running but got %d", pi.BasicStatus),
			"当前项目未处于running状态，无法回收资源")
	}

	// check storage and cgpu remained
	if returnResInfo.SizeAfter == 0 {
		if pi.CpuNodesAcquired != 0 || pi.GpuNodesAcquired != 0 {
			return webapiError.WaErr(webapiError.TypeBadRequest,
				fmt.Sprintf("project CPU/GPU Acquired must be all ZERO before return all storage"),
				"归还全部存储资源前，必须先归还全部计算资源")
		}
	}
	if returnResInfo.SizeAfter > pi.StorageSizeAcquired {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("storage size after return must be smaller than currently acquired"),
			"归还后剩余的计算资源必须小于当前已获得的计算资源")
	}

	// (2) Insert New ApplicationOps
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
			fmt.Sprintf("database operation InsertApplicationOps error: %v", err),
			"在数据库中新建申请单操作记录失败")
	}

	// (3) Update Application
	app.Status = application.AppStatusController
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

	var appCtrlProjectInfo gfApplication.ApprovalInfo
	err = json.Unmarshal([]byte(form.BasicContent), &appCtrlProjectInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to CtrlApprovalInfoWithProjectCode error: %v", err),
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

	var appCtrlProjectInfo gfApplication.ApprovalInfo
	err = json.Unmarshal([]byte(form.BasicContent), &appCtrlProjectInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to CtrlApprovalInfoWithProjectCode error: %v", err),
			"无法解析form.BasicContent的结构")
	}

	//extract applyReturnStoputeRes info
	returnResInfo, err := gfApplication.JsonUnmarshalAppResStoReturn(app.BasicContent)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppResStoReturn error: %v", err),
			"无法解析申请表单的BasicContent的json结构")
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
			fmt.Sprintf("database operation InsertApplicationOps for Controller error: %v", err),
			"无法为调度员在数据库中新建申请单操作记录")
	}

	// Schedule Project Res
	err = wf.prdm.SchedulingStorage(theProject.ProjectID, returnResInfo.SizeAfter,
		"Controller approve ProjectChief return storage", userInfo)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("ProjectResDM SchedulingStorage error: %v", err),
			"归还存储资源错误")
	}

	// update project apply info
	err = wf.pdm.UpdateApplyInfo(project.ApplyInfo{
		ProjectID:           theProject.ProjectID,
		StartDate:           theProject.StartDate,
		TotalDaysApply:      theProject.TotalDaysApply,
		EndReminderAt:       theProject.EndReminderAt,
		CpuNodesExpected:    theProject.CpuNodesExpected,
		GpuNodesExpected:    theProject.GpuNodesExpected,
		StorageSizeExpected: theProject.StorageSizeExpected + (returnResInfo.SizeAfter - theProject.StorageSizeAcquired),
	})
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation pdm.UpdateApplyInfo error: %v", err),
			"在数据库中更新项目ApplyInfo失败")
	}

	// check project storage remained
	// archive project if storage remained is ZERO
	if returnResInfo.SizeAfter == 0 {
		err = wf.pdm.UpdateStatusInfo(project.StatusInfo{
			ProjectID:   theProject.ProjectID,
			BasicStatus: project.BasicStatusSettlement,
		})
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("database operation UpdateStatusInfo project.BasicStatusArchived error: %v", err),
				"无法在数据库中将项目状态置为待结算")
		}

		//发起计量单传阅流程
		// fix a bug: remove systemUserAuto. System can only user ProjectChief to apply
		_, err = wf.bmwf.Apply(generalForm.GeneralForm{
			ProjectID:    theProject.ProjectID,
			FormID:       0,
			Type:         application.AppTypeBrowseMetering,
			Action:       application.AppActionSubmit,
			BasicContent: "",
			ExtraContent: "",
		},
			user.UserInfo{
				UserID:         theProject.ChiefID,
				ChineseName:    theProject.ChiefChineseName,
				DepartmentCode: theProject.DepartmentCode,
				Role:           user.RoleSystem,
			})

		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("System Apply BrowseMeteringWorkflow error: %v", err),
				"无法启动计量单传阅工作流")
		}
	}

	// Update Application
	app.Status = application.AppStatusArchived
	err = wf.adm.Update(app)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("Update for Controller error: %v", err),
			"无法为调度员在数据库中更新Application")
	}

	return nil
}
