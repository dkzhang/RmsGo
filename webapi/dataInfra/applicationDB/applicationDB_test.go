package applicationDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("ApplicationDB", func() {
	var (
		userPC   user.UserInfo
		userApp  user.UserInfo
		userCtrl user.UserInfo
	)
	BeforeEach(func() {
		userPC = user.UserInfo{
			UserID:         1,
			UserName:       "zhj001",
			ChineseName:    "张俊001",
			Department:     "计服中心",
			DepartmentCode: "jf",
			Section:        "",
			Mobile:         "",
			Role:           user.RoleProjectChief,
			Status:         user.StatusNormal,
			Remarks:        "",
		}
		userApp = user.UserInfo{
			UserID:         2,
			UserName:       "zhj002",
			ChineseName:    "张俊002",
			Department:     "计服中心",
			DepartmentCode: "jf",
			Section:        "",
			Mobile:         "",
			Role:           user.RoleApprover,
			Status:         user.StatusNormal,
			Remarks:        "",
		}
		userCtrl = user.UserInfo{
			UserID:         7,
			UserName:       "zhj007",
			ChineseName:    "张俊007",
			Department:     "调度小组",
			DepartmentCode: "ctrl",
			Section:        "",
			Mobile:         "",
			Role:           user.RoleController,
			Status:         user.StatusNormal,
			Remarks:        "",
		}
	})
	Describe("ProjectChief launch a new Project&Resource application", func() {
		// Insert new project
		// Assumed completed
		projectID := 1

		//Insert new application
		It("", func() {
			appID, err := adb.InsertApplication(application.Application{
				ProjectID:                projectID,
				Type:                     application.AppTypeNew,
				Status:                   application.AppStatusProjectChief,
				ApplicantUserID:          userPC.UserID,
				ApplicantUserChineseName: userPC.ChineseName,
				DepartmentCode:           userPC.DepartmentCode,
				BasicContent:             "Application BasicContent",
				ExtraContent:             "Application ExtraContent",
				CreatedAt:                time.Now(),
			})
			Expect(err).ShouldNot(HaveOccurred(), "InsertApplication error: %v", err)
			By(fmt.Sprintf("InsertApplication success, got application ID = %d", appID))

			adb.InsertAppOps(application.AppOpsRecord{
				ProjectID:          projectID,
				ApplicationID:      appID,
				OpsUserID:          0,
				OpsUserChineseName: "",
				Action:             0,
				ActionStr:          "",
				BasicInfo:          "",
				ExtraInfo:          "",
				CreatedAt:          time.Time{},
			})
		})
	})

})
