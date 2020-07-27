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

func (wf Workflow) Apply(form generalForm.GeneralForm, userInfo user.UserInfo) (appID int, err error) {
	var app gfApplication.AppNewProRes
	err = yaml.Unmarshal([]byte(form.BasicContent), &app)
	if err != nil {
		return -1, fmt.Errorf("json Unmarshal to AppNewProRes error: %v", err)
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
		ProjectID:               0,
		BasicStatus:             0,
		ComputingAllocStatus:    0,
		StorageAllocStatus:      0,
		AppInProgressNum:        0,
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
		EndReminderAt:           time.Time{},
		CreatedAt:               time.Time{},
		UpdatedAt:               time.Time{},
	}

	wf.pdb.InsertAllInfo(theProjectS, theProjectD)

	// (2) Insert New Application
	theApplication := application.Application{
		//ApplicationID:            0,
		ProjectID:                -1,
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
		return -1, fmt.Errorf("database operation InsertApplication error: %v", err)
	}

	// (3) Insert New ApplicationOps

	return appID, nil
}

func (wf Workflow) Process(form generalForm.GeneralForm) (err error) {

}
