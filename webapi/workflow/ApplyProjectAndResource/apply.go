package ApplyProjectAndResource

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/webapiError"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDB"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/generalForm"
	"github.com/dkzhang/RmsGo/webapi/model/gfApplication"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"gopkg.in/yaml.v2"
	"time"
)

type Workflow struct {
	adb applicationDB.ApplicationDB
	pdb projectDB.ProjectDB
}

func (wf Workflow) Apply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, waErr webapiError.Err) {
	var app gfApplication.AppNewProRes
	err := yaml.Unmarshal([]byte(form.BasicContent), &app)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
			"无法解析申请表单的BasicContent的json结构")
	}

	// check
	if form.Action != 1 {
		return -1, webapiError.WaErr(webapiError.TypeBadRequest,
			fmt.Sprintf("the Action in Apply must be equal 1"),
			"项目长首次提交申请的action必须为1")
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
		BasicStatus:             project.BasicStatusApplying,
		ComputingAllocStatus:    project.ResNotYetAssigned,
		StorageAllocStatus:      project.ResNotYetAssigned,
		StartDate:               app.StartDate,
		DaysOfUse:               app.DaysOfUse,
		EndDate:                 app.EndDate,
		AppInProgressNum:        1,
		AppAccomplishedNum:      0,
		MeteringInProgressNum:   0,
		MeteringAccomplishedNum: 0,
		ResAllocNum:             0,
		CpuNodesExpected:        0,
		GpuNodesExpected:        0,
		StorageSizeExpected:     0,
		CpuNodesAcquired:        0,
		GpuNodesAcquired:        0,
		StorageSizeAcquired:     0,
		TotalDaysApply:          0,
		EndReminderAt:           time.Now().AddDate(100, 0, 0),
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}

	projectID, err := wf.pdb.InsertAllInfo(theProjectS, theProjectD)
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

	appID, err = wf.adb.InsertApplication(theApplication)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplication error: %v", err),
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

	_, err = wf.adb.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return -1, webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation InsertApplicationOps error: %v", err),
			"在数据库中新建申请单操作记录失败")
	}
	return appID, nil
}

func (wf Workflow) Process(form generalForm.GeneralForm, userInfo user.UserInfo) (waErr webapiError.Err) {

	// Query Application and Project static & dynamic
	theApplication, err := wf.adb.QueryApplicationByID(form.FormID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryApplicationByID error: %v", err),
			"在数据库中查询申请单失败")
	}

	theProjectS, err := wf.pdb.QueryStaticInfoByID(theApplication.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryApplicationByID error: %v", err),
			"在数据库中查询项目静态信息失败")
	}

	theProjectD, err := wf.pdb.QueryDynamicInfoByID(theApplication.ProjectID)
	if err != nil {
		return webapiError.WaErr(webapiError.TypeDatabaseError,
			fmt.Sprintf("database operation QueryApplicationByID error: %v", err),
			"在数据库中查询项目动态信息失败")
	}

	switch theApplication.Status {
	case application.AppStatusProjectChief:
		if userInfo.Role != user.RoleProjectChief {
			return webapiError.WaErr(webapiError.TypeAuthorityError,
				fmt.Sprintf("application status = %d but current user is %d", theApplication.Status, userInfo.Role),
				"当前用户无权操作该申请单，不符合设定流程")
		}
		if userInfo.UserID != theApplication.ApplicantUserID {
			return webapiError.WaErr(webapiError.TypeAuthorityError,
				fmt.Sprintf("application owner id = %d but current user is %d", theApplication.ApplicantUserID, userInfo.UserID),
				"当前用户无权操作该申请单，不是创建该表单的项目长")
		}

		//项目长重新提交表单
		if form.Action != 1 {
			return webapiError.WaErr(webapiError.TypeBadRequest,
				fmt.Sprintf("the Action for ProjectChief ReApply must be equal 1"),
				"项目长重新提交申请单时form.Action必须为1")
		}

		var app gfApplication.AppNewProRes
		err = yaml.Unmarshal([]byte(form.BasicContent), &app)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeBadRequest,
				fmt.Sprintf("json Unmarshal to AppNewProRes error: %v", err),
				"无法解析form.BasicContent的结构")
		}

		// Insert New ApplicationOps
		theAppOpsRecord := application.AppOpsRecord{
			//RecordID:           0,
			ProjectID:          theProjectS.ProjectID,
			ApplicationID:      theApplication.ApplicationID,
			OpsUserID:          userInfo.UserID,
			OpsUserChineseName: userInfo.ChineseName,
			Action:             form.Action,
			ActionStr:          "重新提交",
			BasicInfo:          form.BasicContent,
			ExtraInfo:          form.ExtraContent,
			CreatedAt:          time.Now(),
		}
		_, err := wf.adb.InsertAppOps(theAppOpsRecord)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("database operation InsertApplicationOps error: %v", err),
				"无法为项目长在数据库中新建申请单操作记录")
		}

		// Update Application
		theApplication.Status = application.AppStatusApprover
		theApplication.BasicContent = form.BasicContent
		theApplication.ExtraContent = form.ExtraContent
		theApplication.UpdatedAt = time.Now()
		err = wf.adb.UpdateApplication(theApplication)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("UpdateApplication for ProjectChief error: %v", err),
				"无法为项目长在数据库中更新Application")
		}

		// Update Project
		theProjectS.ProjectName = app.ProjectName
		theProjectS.ExtraInfo = form.ExtraContent
		theProjectS.UpdatedAt = time.Now()
		err = wf.pdb.UpdateStaticInfo(theProjectS)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("UpdateApplication for ProjectChief error: %v", err),
				"无法为项目长在数据库中更新Project静态信息")
		}

		theProjectD.StartDate = app.StartDate
		theProjectD.DaysOfUse = app.DaysOfUse
		theProjectD.EndDate = app.EndDate
		theProjectD.UpdatedAt = time.Now()
		err = wf.pdb.UpdateDynamicInfo(theProjectD)
		if err != nil {
			return webapiError.WaErr(webapiError.TypeDatabaseError,
				fmt.Sprintf("UpdateApplication for ProjectChief error: %v", err),
				"无法为项目长在数据库中更新Project动态信息")
		}

		return nil

	case application.AppStatusApprover:
		if userInfo.Role != user.RoleApprover {
			return webapiError.WaErr(webapiError.TypeAuthorityError,
				fmt.Sprintf("application status = %d but current user is %d", theApplication.Status, userInfo.Role),
				"当前用户无权操作该申请单，不符合设定流程")
		}
		if userInfo.UserID != theApplication.ApplicantUserID {
			return webapiError.WaErr(webapiError.TypeAuthorityError,
				fmt.Sprintf("application department code = %s but current user's department code is %s", theApplication.DepartmentCode, userInfo.DepartmentCode),
				"当前用户无权操作该申请单，不是该表单所属单位的审批人")
		}
	case application.AppStatusController:
		if userInfo.Role != user.RoleController {
			return webapiError.WaErr(webapiError.TypeAuthorityError,
				fmt.Sprintf("application status = %d but current user is %d", theApplication.Status, userInfo.Role),
				"当前用户无权操作该申请单，不符合设定流程")
		}

	}

	// (1) Update Project

	// (2) Update Application

	// (3) Update ApplicationOps

	return webapiError.WaErr(webapiError.TypeNotYetImplemented, "not yet implemented", "尚未实现")
}
