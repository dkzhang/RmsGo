package projectDM_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("ProjectDM", func() {

	Describe("InsertAndRetrieve", func() {
		//Query All Project Info from dm
		It("QueryAllInfo from dm", func() {
			proSs := make([]project.StaticInfo, 3)
			proDs := make([]project.DynamicInfo, 3)

			for i := 0; i < 3; i++ {
				proSs[i] = project.StaticInfo{
					ProjectName:      fmt.Sprintf("ProjectName%d", i),
					ProjectCode:      fmt.Sprintf("ProjectCode%d", i),
					DepartmentCode:   "jf",
					Department:       "计服中心",
					ChiefID:          1,
					ChiefChineseName: "项目长001",
					ExtraInfo:        fmt.Sprintf("ExtraInfo%d", i),
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}

				proDs[i] = project.DynamicInfo{
					BasicStatus:          rand.Intn(100),
					ComputingAllocStatus: rand.Intn(100),
					StorageAllocStatus:   rand.Intn(100),
					StartDate:            time.Now(),
					TotalDaysApply:       10,
					EndReminderAt:        time.Now().AddDate(0, 0, 10),
					CpuNodesExpected:     rand.Intn(100),
					GpuNodesExpected:     rand.Intn(100),
					StorageSizeExpected:  rand.Intn(100),
					CpuNodesAcquired:     rand.Intn(100),
					GpuNodesAcquired:     rand.Intn(100),
					StorageSizeAcquired:  rand.Intn(100),
					CreatedAt:            time.Now(),
					UpdatedAt:            time.Now(),
				}

				projectID, err := pdm.InsertAllInfo(project.ProjectInfo{
					ProjectID:      0,
					TheStaticInfo:  proSs[i],
					TheDynamicInfo: proDs[i],
				})

				Expect(err).ShouldNot(HaveOccurred(), "InsertAllInfo %d error: %v", i, err)
				By(fmt.Sprintf("InsertAllInfo %d success, projectID=%d", i, projectID))
				proSs[i].ProjectID = projectID
				proDs[i].ProjectID = projectID
			}

			for j := 1; j < 4; j++ {
				psi, err := pdm.QueryStaticInfoByID(j)
				Expect(err).ShouldNot(HaveOccurred(), "QueryStaticInfoByID %d error: %v", j, err)

				Expect(psi.CreatedAt).Should(BeTemporally("~", proSs[j-1].CreatedAt, time.Second))
				Expect(psi.UpdatedAt).Should(BeTemporally("~", proSs[j-1].UpdatedAt, time.Second))
				proSs[j-1].CreatedAt = psi.CreatedAt
				proSs[j-1].UpdatedAt = psi.UpdatedAt

				Expect(psi).Should(Equal(proSs[j-1]))
				By(fmt.Sprintf("project static info: %v", psi))
			}

			for j := 1; j < 4; j++ {
				pdi, err := pdm.QueryDynamicInfoByID(j)
				Expect(err).ShouldNot(HaveOccurred(), "QueryDynamicInfoByID %d error: %v", j, err)

				Expect(pdi.CreatedAt).Should(BeTemporally("~", proDs[j-1].CreatedAt, time.Second))
				Expect(pdi.UpdatedAt).Should(BeTemporally("~", proDs[j-1].UpdatedAt, time.Second))

				Expect(pdi.StartDate).Should(BeTemporally("~", proDs[j-1].StartDate, time.Second))
				Expect(pdi.EndReminderAt).Should(BeTemporally("~", proDs[j-1].EndReminderAt, time.Second))

				proDs[j-1].CreatedAt = pdi.CreatedAt
				proDs[j-1].UpdatedAt = pdi.UpdatedAt
				proDs[j-1].StartDate = pdi.StartDate
				proDs[j-1].EndReminderAt = pdi.EndReminderAt

				Expect(pdi).Should(Equal(proDs[j-1]))
				By(fmt.Sprintf("project dynamic info: %v", pdi))
			}
		})
	})

	Describe("Update from dm", func() {
		//Query All Project Info from dm
		It("UpdateStaticInfo", func() {
			id := 2

			psi, err := pdm.QueryStaticInfoByID(id)
			Expect(err).ShouldNot(HaveOccurred(), "QueryStaticInfoByID %d error: %v", id, err)

			psi.ProjectName = fmt.Sprintf("ProjectNameUpdated%d", id)
			psi.ProjectCode = fmt.Sprintf("ProjectNameUpdated%d", id)
			psi.ExtraInfo = fmt.Sprintf("ProjectNameUpdated%d", id)
			err = pdm.UpdateStaticInfo(psi)
			Expect(err).ShouldNot(HaveOccurred(), "UpdateStaticInfo %d error: %v", id, err)

			psiU, err := pdm.QueryStaticInfoByID(id)
			Expect(err).ShouldNot(HaveOccurred(), "QueryStaticInfoByID %d error: %v", id, err)
			psi.UpdatedAt = psiU.UpdatedAt
			Expect(psiU).Should(Equal(psi))
		})

		It("UpdateStaticInfo", func() {
			id := 3

			pdi, err := pdm.QueryDynamicInfoByID(id)
			Expect(err).ShouldNot(HaveOccurred(), "QueryDynamicInfoByID %d error: %v", id, err)

			pdi.BasicStatus = rand.Intn(100)
			pdi.CpuNodesExpected = rand.Intn(100)

			err = pdm.UpdateDynamicInfo(pdi)
			Expect(err).ShouldNot(HaveOccurred(), "UpdateDynamicInfo %d error: %v", id, err)

			pdiU, err := pdm.QueryDynamicInfoByID(id)
			Expect(err).ShouldNot(HaveOccurred(), "QueryDynamicInfoByID %d error: %v", id, err)
			pdi.UpdatedAt = pdiU.UpdatedAt
			Expect(pdiU).Should(Equal(pdi))
		})
	})

	Describe("QueryAllInfo from dm", func() {
		//Query All Project Info from dm
		It("QueryAllInfo from dm", func() {
			_, err := pdm.QueryAllInfo()
			Expect(err).ShouldNot(HaveOccurred(), "QueryAllInfo error: %v")
		})
	})

})
