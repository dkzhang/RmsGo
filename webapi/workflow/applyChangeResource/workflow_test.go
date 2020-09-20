package applyChangeResource_test

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

	Describe("ProjectChief launch a new Resource  Change application", func() {
		It("ProjectChief Apply 3 Application No. 4,5,6", func() {
			projectID := 1

			arc := gfApplication.AppResChange{
				Resource: resource.Resource{
					CpuNodes:    projectID * 21,
					GpuNodes:    projectID * 11,
					StorageSize: projectID * 51,
				},
				EndDateStr: "2020-10-01",
			}
			arcB, _ := json.Marshal(arc)
			arcJson := string(arcB)

			form := generalForm.GeneralForm{
				ProjectID:    projectID,
				FormID:       0,
				Type:         application.AppTypeChange,
				Action:       application.AppActionSubmit,
				BasicContent: arcJson,
				ExtraContent: fmt.Sprintf("extra content %d", projectID),
			}

			appID, waErr := gwf.Apply(form, userPC)
			Expect(waErr).ShouldNot(HaveOccurred(), fmt.Sprintf("apply change for project %d error", projectID))
			By(fmt.Sprintf("Apply New ProRes Application %d success, got appID = %d", projectID, appID))

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
			By(fmt.Sprintf("QueryProjectInfoByID %d = %v", appAfter.ProjectID, proInfoAfter))

			appsAfter, err := adm.QueryAll(-1, -1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(appsAfter)).Should(Equal(4))
		})
	})

	Describe("Approver examine the new Project&Resource application, agree app1, reject app2", func() {
		It("Approver query the application, agree project1~app4", func() {
			appID := 4
			projectID := 1
			apps, err := adm.QueryByDepartmentCode(userApp.DepartmentCode, application.AppTypeALL, application.AppStatusApprover)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("QueryByDepartmentCode len(apps) = %d", len(apps)))

			app4, err := adm.QueryByID(appID)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("QueryByID %d = %v", appID, app4))

			waErr := gwf.Process(generalForm.GeneralForm{
				ProjectID:    projectID,
				FormID:       appID,
				Type:         application.AppTypeChange,
				Action:       1,
				BasicContent: "",
				ExtraContent: "",
			}, app4, userApp)
			Expect(waErr).Should(BeNil())
		})

		//It("Approver query the application, reject project2~app5", func() {
		//	appID := 5
		//	projectID := 2
		//	apps, err := adm.QueryByDepartmentCode(userApp.DepartmentCode, application.AppTypeALL, application.AppStatusALL)
		//	Expect(err).ShouldNot(HaveOccurred())
		//	By(fmt.Sprintf("QueryByDepartmentCode len(apps) = %d", len(apps)))
		//
		//	app, err := adm.QueryByID(appID)
		//	Expect(err).ShouldNot(HaveOccurred())
		//	By(fmt.Sprintf("QueryByID %d = %v", appID, app))
		//
		//	waErr := gwf.Process(generalForm.GeneralForm{
		//		ProjectID:    projectID,
		//		FormID:       appID,
		//		Type:         application.AppTypeChange,
		//		Action:       1,
		//		BasicContent: "",
		//		ExtraContent: "",
		//	}, app, userApp)
		//	Expect(waErr).Should(BeNil())
		//})
	})

	//Describe("Controller examine the new Project&Resource application, agree app1", func() {
	//	It("Controller query the application, agree app1", func() {
	//		appID := 1
	//
	//		apps, err := adm.QueryByDepartmentCode(userApp.DepartmentCode, application.AppTypeALL, application.AppStatusApprover)
	//		Expect(err).ShouldNot(HaveOccurred())
	//		By(fmt.Sprintf("QueryByDepartmentCode len(apps) = %d", len(apps)))
	//
	//		app1, err := adm.QueryByID(appID)
	//		Expect(err).ShouldNot(HaveOccurred())
	//		By(fmt.Sprintf("QueryByID 1 = %v", app1))
	//
	//		bcs := gfApplication.CtrlApprovalInfo{ProjectCode: fmt.Sprintf("ProjectCode%d", app1.ProjectID)}
	//		bcb, _ := json.Marshal(bcs)
	//
	//		waErr := gwf.Process(generalForm.GeneralForm{
	//			ProjectID:    app1.ProjectID,
	//			FormID:       app1.ApplicationID,
	//			Type:         application.AppTypeNew,
	//			Action:       1,
	//			BasicContent: string(bcb),
	//			ExtraContent: "",
	//		}, app1, userCtrl)
	//		Expect(waErr).Should(BeNil())
	//
	//		// Check After
	//		app1After, err := adm.QueryByID(appID)
	//		Expect(err).ShouldNot(HaveOccurred())
	//		Expect(app1After.Status).Should(Equal(application.AppStatusArchived))
	//		By(fmt.Sprintf("QueryByID 1 = %v", app1After))
	//
	//		proInfo1After, err := pdm.QueryByID(app1.ProjectID)
	//		Expect(err).ShouldNot(HaveOccurred())
	//		Expect(proInfo1After.ProjectCode).Should(Equal(bcs.ProjectCode))
	//	})
	//})

	Describe("ProjectChief launch a new Resource  Change application", func() {
		It("ProjectChief Apply 3 Application No. 4,5,6", func() {
			By(fmt.Sprintf("%v%v", userApp, userCtrl))
		})
	})
})
