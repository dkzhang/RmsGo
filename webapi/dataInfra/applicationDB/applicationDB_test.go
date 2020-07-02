package applicationDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
			ChineseName:    "项目长001",
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
			ChineseName:    "审批人002",
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
			ChineseName:    "调度员007",
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

		basicStr := "Application %d BasicContent"
		extraStr := "Application %d ExtraContent"

		//Insert new application
		It("ProjectChief insert 3 Application and 3 AppOpsRecord", func() {
			for i := 1; i <= 3; i++ {
				appID, err := adb.InsertApplication(application.Application{
					ProjectID:                projectID,
					Type:                     application.AppTypeNew,
					Status:                   application.AppStatusProjectChief,
					ApplicantUserID:          userPC.UserID,
					ApplicantUserChineseName: userPC.ChineseName,
					DepartmentCode:           userPC.DepartmentCode,
					BasicContent:             fmt.Sprintf(basicStr, i),
					ExtraContent:             fmt.Sprintf(extraStr, i),
				})
				Expect(err).ShouldNot(HaveOccurred(), "InsertApplication %d error: %v", i, err)
				By(fmt.Sprintf("InsertApplication %d success, got application ID = %d", i, appID))

				recordID, err := adb.InsertAppOps(application.AppOpsRecord{
					ProjectID:          projectID,
					ApplicationID:      appID,
					OpsUserID:          userPC.UserID,
					OpsUserChineseName: userPC.ChineseName,
					Action:             1,
					ActionStr:          "提交",
					BasicInfo:          fmt.Sprintf(basicStr, i),
					ExtraInfo:          fmt.Sprintf(extraStr, i),
				})
				Expect(err).ShouldNot(HaveOccurred(), "InsertAppOps %d error: %v", i, err)
				By(fmt.Sprintf("InsertAppOps %d success, got ops record ID = %d", i, recordID))
			}
		})
	})

	Describe("Approver examine the new Project&Resource application", func() {
		projectID := 1
		appID := 1

		It("Approver query the application,insert an AppOpsRecord, update the application", func() {

			app, err := adb.QueryApplicationByID(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryApplicationByID error: %v", err)
			By(fmt.Sprintf("QueryApplicationByID success, got application = %v", app))

			recordID, err := adb.InsertAppOps(application.AppOpsRecord{
				ProjectID:          projectID,
				ApplicationID:      appID,
				OpsUserID:          userApp.UserID,
				OpsUserChineseName: userApp.ChineseName,
				Action:             1,
				ActionStr:          "是",
				BasicInfo:          "",
				ExtraInfo:          "同意",
			})
			Expect(err).ShouldNot(HaveOccurred(), "InsertAppOps error: %v", err)
			By(fmt.Sprintf("InsertAppOps success, got ops record ID = %d", recordID))

			app.Status = application.AppStatusController
			err = adb.UpdateApplication(app)
			Expect(err).ShouldNot(HaveOccurred(), "UpdateApplication error: %v", err)
			By(fmt.Sprintf("UpdateApplication success"))
		})

	})

	Describe("Controller check the new Project&Resource application", func() {
		projectID := 1
		appID := 1

		It("Controller query the application,insert an AppOpsRecord, update the application", func() {

			app, err := adb.QueryApplicationByID(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryApplicationByID error: %v", err)
			By(fmt.Sprintf("QueryApplicationByID success, got application = %v", app))

			recordID, err := adb.InsertAppOps(application.AppOpsRecord{
				ProjectID:          projectID,
				ApplicationID:      appID,
				OpsUserID:          userCtrl.UserID,
				OpsUserChineseName: userCtrl.ChineseName,
				Action:             1,
				ActionStr:          "是",
				BasicInfo:          "",
				ExtraInfo:          "同意",
			})
			Expect(err).ShouldNot(HaveOccurred(), "InsertAppOps error: %v", err)
			By(fmt.Sprintf("InsertAppOps success, got ops record ID = %d", recordID))

			app.Status = application.AppStatusArchived
			err = adb.UpdateApplication(app)
			Expect(err).ShouldNot(HaveOccurred(), "UpdateApplication error: %v", err)
			By(fmt.Sprintf("UpdateApplication success"))
		})
	})

	Describe("QueryAppOpsByAppId", func() {
		It("Query application 1 ops", func() {
			appID := 1
			records, err := adb.QueryAppOpsByAppId(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryAppOpsByAppId error: %v", err)
			By(fmt.Sprintf("QueryAppOpsByAppId success: %v", records))
		})
	})
})
