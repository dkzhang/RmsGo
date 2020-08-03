package ApplyProjectAndResource

import (
	"fmt"
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

func (wf Workflow) Apply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, err error, msg string) {
	var app gfApplication.AppNewProRes
	err = yaml.Unmarshal([]byte(form.BasicContent), &app)
	if err != nil {
		return -1, fmt.Errorf("json Unmarshal to AppNewProRes error: %v", err),
			fmt.Sprintf("无法解析form.BasicContent的结构")
	}

	// check
	if form.Action != 1 {
		return -1, fmt.Errorf("the Action in Apply must be equal 1"),
			fmt.Sprintf("首次提交申请单时form.Action必须为1")
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
		StartBillingAt:          time.Time{},
		TotalDaysApply:          0,
		EndReminderAt:           time.Now().AddDate(100, 0, 0),
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}

	projectID, err := wf.pdb.InsertAllInfo(theProjectS, theProjectD)
	if err != nil {
		return -1, fmt.Errorf("insert project info error: %v", err),
			fmt.Sprintf("在数据库中新建项目记录失败")
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
		return -1, fmt.Errorf("database operation InsertApplication error: %v", err),
			fmt.Sprintf("在数据库中新建申请单记录失败")
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

	recordID, err := wf.adb.InsertAppOps(theAppOpsRecord)
	if err != nil {
		return -1, fmt.Errorf("database operation InsertApplicationOps error: %v", err),
			fmt.Sprintf("在数据库中新建申请单操作记录失败")
	}

	return appID, nil,
		fmt.Sprintf("首次提交申请单成功,ProjectID=%d, ApplicationID=%d, RecordID=%d", projectID, appID, recordID)
}

func (wf Workflow) Process(form generalForm.GeneralForm, userInfo user.UserInfo) (err error, msg string) {

	return fmt.Errorf("NEED TODO"), ""
}
