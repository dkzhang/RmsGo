package applicationDM_test

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/gfApplication"
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("ApplicationDM", func() {
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

		It("ProjectChief insert 3 Application and 3 AppOpsRecord", func() {
			apps := make([]application.Application, 3)
			appOpss := make([]application.AppOpsRecord, 3)
			for i := 1; i <= 3; i++ {
				// Insert Application
				bcs := gfApplication.AppNewProRes{
					ProjectName: fmt.Sprintf("ProjectName%d", i),
					Resource: resource.Resource{
						CpuNodes:    i * 10,
						GpuNodes:    i * 20,
						StorageSize: i * 30,
					},
					StartDateStr:   "2020-01-03",
					TotalDaysApply: 10,
					EndDateStr:     "2020-01-13",
				}
				bcb, _ := json.Marshal(bcs)
				apps[i-1] = application.Application{
					ProjectID:                i,
					Type:                     application.AppTypeNew,
					Status:                   application.AppStatusApprover,
					ApplicantUserID:          userPC.UserID,
					ApplicantUserChineseName: userPC.ChineseName,
					DepartmentCode:           userPC.DepartmentCode,
					BasicContent:             string(bcb),
					ExtraContent:             fmt.Sprintf("ExtraContent%d", i),
					CreatedAt:                time.Now(),
					UpdatedAt:                time.Now(),
				}

				appID, err := adm.Insert(apps[i-1])
				Expect(err).ShouldNot(HaveOccurred(), "Insert %d error: %v", i, err)
				By(fmt.Sprintf("Insert %d success, got application ID = %d", i, appID))

				// Insert AppOpsRecord
				appOpss[i-1] = application.AppOpsRecord{
					ProjectID:          i,
					ApplicationID:      appID,
					OpsUserID:          userPC.UserID,
					OpsUserChineseName: userPC.ChineseName,
					Action:             1,
					ActionStr:          "提交",
					BasicInfo:          string(bcb),
					ExtraInfo:          fmt.Sprintf("ExtraContent%d", i),
					CreatedAt:          time.Now(),
				}
				recordID, err := adm.InsertAppOps(appOpss[i-1])
				Expect(err).ShouldNot(HaveOccurred(), "InsertAppOps %d error: %v", i, err)
				By(fmt.Sprintf("InsertAppOps %d success, got ops record ID = %d", i, recordID))
			}

			// Retrieve Application and AppOpsRecord by id
			for i := 1; i <= 3; i++ {
				app, err := adm.QueryByID(i)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(app.ApplicationID).Should(Equal(i))
				Expect(app.CreatedAt).Should(BeTemporally("~", apps[i-1].CreatedAt, time.Second))
				Expect(app.UpdatedAt).Should(BeTemporally("~", apps[i-1].UpdatedAt, time.Second))
				apps[i-1].CreatedAt = app.CreatedAt
				apps[i-1].UpdatedAt = app.UpdatedAt
				apps[i-1].ApplicationID = app.ApplicationID
				Expect(app).Should(Equal(apps[i-1]))
				By(fmt.Sprintf("application info %d : %v", i, app))

				appOpsArray, err := adm.QueryAppOpsByAppId(i)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(appOpsArray)).Should(Equal(1))
				appOps := appOpsArray[0]
				Expect(appOps.CreatedAt).Should(BeTemporally("~", appOpss[i-1].CreatedAt, time.Second))
				appOpss[i-1].CreatedAt = appOps.CreatedAt
				appOpss[i-1].RecordID = appOps.RecordID
				Expect(appOps).Should(Equal(appOpss[i-1]))
				By(fmt.Sprintf("appOps record %d : %v", i, appOps))
			}

		})
		It("ProjectChief insert 3 Application and 3 AppOpsRecord", func() {
			By(fmt.Sprintf("%v, %v", userApp, userCtrl))
		})

	})

	Describe("Approver examine the new Project&Resource application, agree app1, reject app2", func() {
		It("Approver agree app1, query the application,insert an AppOpsRecord, update the application", func() {
			projectID := 1
			appID := 1

			app, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryByID error: %v", err)
			Expect(app.ApplicationID).Should(Equal(appID))
			Expect(app.Status).Should(Equal(application.AppStatusApprover))
			By(fmt.Sprintf("QueryByID success, got application = %v", app))

			appOps := application.AppOpsRecord{
				ProjectID:          projectID,
				ApplicationID:      appID,
				OpsUserID:          userApp.UserID,
				OpsUserChineseName: userApp.ChineseName,
				Action:             1,
				ActionStr:          "是",
				BasicInfo:          "BasicInfo: 1",
				ExtraInfo:          "ExtraInfo:同意",
				CreatedAt:          time.Now(),
			}
			recordID, err := adm.InsertAppOps(appOps)
			Expect(err).ShouldNot(HaveOccurred(), "InsertAppOps error: %v", err)
			By(fmt.Sprintf("InsertAppOps success, got ops record ID = %d", recordID))

			app.Status = application.AppStatusController
			err = adm.Update(app)
			Expect(err).ShouldNot(HaveOccurred(), "Update error: %v", err)
			By(fmt.Sprintf("Update success"))

			// check result
			appAfter, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryByID error: %v", err)
			Expect(appAfter.Status).Should(Equal(application.AppStatusController))
			By(fmt.Sprintf("QueryByID success, got application = %v", appAfter))

			appOpsArrayAfter, err := adm.QueryAppOpsByAppId(appID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(appOpsArrayAfter)).Should(Equal(2))
			appOpsAfter := appOpsArrayAfter[1]
			Expect(appOpsAfter.CreatedAt).Should(BeTemporally("~", appOps.CreatedAt, time.Second))
			appOps.CreatedAt = appOpsAfter.CreatedAt
			appOps.RecordID = appOpsAfter.RecordID
			Expect(appOpsAfter).Should(Equal(appOps))
			By(fmt.Sprintf("appOpsAfter record : %v", appOpsAfter))
		})

		It("Approver reject app2, query the application,insert an AppOpsRecord, update the application", func() {

			projectID := 2
			appID := 2

			app, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryByID error: %v", err)
			By(fmt.Sprintf("QueryByID success, got application = %v", app))

			appOps := application.AppOpsRecord{
				ProjectID:          projectID,
				ApplicationID:      appID,
				OpsUserID:          userApp.UserID,
				OpsUserChineseName: userApp.ChineseName,
				Action:             1,
				ActionStr:          "否",
				BasicInfo:          "BasicInfo: 1",
				ExtraInfo:          "ExtraInfo:拒绝",
				CreatedAt:          time.Now(),
			}
			recordID, err := adm.InsertAppOps(appOps)
			Expect(err).ShouldNot(HaveOccurred(), "InsertAppOps error: %v", err)
			By(fmt.Sprintf("InsertAppOps success, got ops record ID = %d", recordID))

			app.Status = application.AppStatusProjectChief
			err = adm.Update(app)
			Expect(err).ShouldNot(HaveOccurred(), "Update error: %v", err)
			By(fmt.Sprintf("Update success"))

			// check result
			appAfter, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryByID error: %v", err)
			Expect(appAfter.Status).Should(Equal(application.AppStatusProjectChief))
			By(fmt.Sprintf("QueryByID success, got application = %v", appAfter))

			appOpsArrayAfter, err := adm.QueryAppOpsByAppId(appID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(appOpsArrayAfter)).Should(Equal(2))
			appOpsAfter := appOpsArrayAfter[1]
			Expect(appOpsAfter.CreatedAt).Should(BeTemporally("~", appOps.CreatedAt, time.Second))
			appOps.CreatedAt = appOpsAfter.CreatedAt
			appOps.RecordID = appOpsAfter.RecordID
			Expect(appOpsAfter).Should(Equal(appOps))
			By(fmt.Sprintf("appOpsAfter record : %v", appOpsAfter))
		})
	})

	Describe("Controller check the new Project&Resource application, agree app1", func() {
		It("Controller agree app1", func() {

			projectID := 1
			appID := 1

			app, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryByID error: %v", err)
			Expect(app.Status).Should(Equal(application.AppStatusController))
			By(fmt.Sprintf("QueryByID success, got application = %v", app))

			bcs := gfApplication.CtrlApprovalInfoWithProjectCode{ProjectCode: fmt.Sprintf("ProjectCode%d", projectID)}
			bcb, _ := json.Marshal(bcs)
			appOps := application.AppOpsRecord{
				ProjectID:          projectID,
				ApplicationID:      appID,
				OpsUserID:          userApp.UserID,
				OpsUserChineseName: userApp.ChineseName,
				Action:             1,
				ActionStr:          "是",
				BasicInfo:          string(bcb),
				ExtraInfo:          "ExtraInfo:同意",
				CreatedAt:          time.Now(),
			}
			recordID, err := adm.InsertAppOps(appOps)
			Expect(err).ShouldNot(HaveOccurred(), "InsertAppOps error: %v", err)
			By(fmt.Sprintf("InsertAppOps success, got ops record ID = %d", recordID))

			appNewProRes, err := gfApplication.JsonUnmarshalAppNewProRes(app.BasicContent)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("Unmarshal AppNewProRes: %v", appNewProRes))

			app.Status = application.AppStatusArchived
			err = adm.Update(app)
			Expect(err).ShouldNot(HaveOccurred(), "Update error: %v", err)
			By(fmt.Sprintf("Update success"))

			// check result
			appAfter, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryByID error: %v", err)
			Expect(appAfter.Status).Should(Equal(application.AppStatusArchived))
			By(fmt.Sprintf("QueryByID success, got application = %v", appAfter))

			appOpsArrayAfter, err := adm.QueryAppOpsByAppId(appID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(appOpsArrayAfter)).Should(Equal(3))
			appOpsAfter := appOpsArrayAfter[2]
			Expect(appOpsAfter.CreatedAt).Should(BeTemporally("~", appOps.CreatedAt, time.Second))
			appOps.CreatedAt = appOpsAfter.CreatedAt
			appOps.RecordID = appOpsAfter.RecordID
			Expect(appOpsAfter).Should(Equal(appOps))
			By(fmt.Sprintf("appOpsAfter record : %v", appOpsAfter))
		})
	})
})
