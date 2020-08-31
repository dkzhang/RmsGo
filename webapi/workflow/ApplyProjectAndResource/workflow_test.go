package ApplyProjectAndResource_test

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/generalForm"
	"github.com/dkzhang/RmsGo/webapi/model/gfApplication"
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Workflow", func() {
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
		It("ProjectChief Apply 3 Application", func() {
			for i := 1; i < 4; i++ {

				anpr := gfApplication.AppNewProRes{
					ProjectName: fmt.Sprintf("Project%d", i),
					Resource: resource.Resource{
						CpuNodes:    i * 20,
						GpuNodes:    i * 10,
						StorageSize: i * 50,
					},
					StartDateStr:   "2020-08-17",
					TotalDaysApply: 10,
					EndDateStr:     "2020-08-27",
				}
				anprB, _ := json.Marshal(anpr)
				anprJson := string(anprB)

				form := generalForm.GeneralForm{
					ProjectID:    0,
					FormID:       0,
					Type:         application.AppTypeNew,
					Action:       application.AppActionSubmit,
					BasicContent: anprJson,
					ExtraContent: fmt.Sprintf("extra content %d", i),
				}

				appID, waErr := gwf.Apply(form, userPC)
				Expect(waErr).ShouldNot(HaveOccurred())
				By(fmt.Sprintf("Apply New ProRes Application %d success, got appID = %d", i, appID))

				/////////////////////////////////
				formB, _ := json.Marshal(form)
				By(string(formB))
				/////////////////////////////////

				//Check After
				appAfter, err := adm.QueryByID(appID)
				Expect(err).ShouldNot(HaveOccurred())
				By(fmt.Sprintf("QueryByID %d = %v", appID, appAfter))

				proInfoAfter, err := pdm.QueryByID(appAfter.ProjectID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(proInfoAfter.ProjectName).Should(Equal(anpr.ProjectName))
				By(fmt.Sprintf("QueryStaticInfoByID %d = %v", appAfter.ProjectID, proInfoAfter))
			}

			appsAfter, err := adm.QueryAll(-1, -1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(appsAfter)).Should(Equal(3))
		})
	})

	Describe("Approver examine the new Project&Resource application, agree app1, reject app2", func() {
		It("Approver query the application, agree app1", func() {
			apps, err := adm.QueryByDepartmentCode(userApp.DepartmentCode, application.AppTypeALL, application.AppStatusApprover)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("QueryByDepartmentCode len(apps) = %d", len(apps)))

			app1, err := adm.QueryByID(1)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("QueryByID 1 = %v", app1))

			waErr := gwf.Process(generalForm.GeneralForm{
				ProjectID:    1,
				FormID:       1,
				Type:         application.AppTypeNew,
				Action:       1,
				BasicContent: "",
				ExtraContent: "",
			}, app1, userApp)
			Expect(waErr).Should(BeNil())
		})

		It("Approver query the application, reject app2", func() {
			apps, err := adm.QueryByDepartmentCode(userApp.DepartmentCode, application.AppTypeALL, application.AppStatusALL)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("QueryByDepartmentCode len(apps) = %d", len(apps)))

			app2, err := adm.QueryByID(2)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("QueryByID 2 = %v", app2))

			waErr := gwf.Process(generalForm.GeneralForm{
				ProjectID:    2,
				FormID:       2,
				Type:         application.AppTypeNew,
				Action:       1,
				BasicContent: "",
				ExtraContent: "",
			}, app2, userApp)
			Expect(waErr).Should(BeNil())
		})
	})

	Describe("Controller examine the new Project&Resource application, agree app1", func() {
		It("Controller query the application, agree app1", func() {
			appID := 1

			apps, err := adm.QueryByDepartmentCode(userApp.DepartmentCode, application.AppTypeALL, application.AppStatusApprover)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("QueryByDepartmentCode len(apps) = %d", len(apps)))

			app1, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("QueryByID 1 = %v", app1))

			bcs := gfApplication.CtrlApprovalInfo{ProjectCode: fmt.Sprintf("ProjectCode%d", app1.ProjectID)}
			bcb, _ := json.Marshal(bcs)

			waErr := gwf.Process(generalForm.GeneralForm{
				ProjectID:    app1.ProjectID,
				FormID:       app1.ApplicationID,
				Type:         application.AppTypeNew,
				Action:       1,
				BasicContent: string(bcb),
				ExtraContent: "",
			}, app1, userCtrl)
			Expect(waErr).Should(BeNil())

			// Check After
			app1After, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(app1After.Status).Should(Equal(application.AppStatusArchived))
			By(fmt.Sprintf("QueryByID 1 = %v", app1After))

			proInfo1After, err := pdm.QueryByID(app1.ProjectID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(proInfo1After.ProjectCode).Should(Equal(bcs.ProjectCode))
		})
	})
})
