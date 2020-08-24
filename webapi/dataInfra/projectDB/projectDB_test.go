package projectDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("ProjectDB", func() {

	Describe("InsertAndRetrieve", func() {
		//Query All Project Info from db
		It("QueryAllInfo from db", func() {
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
					BasicStatus:             rand.Intn(100),
					ComputingAllocStatus:    rand.Intn(100),
					StorageAllocStatus:      rand.Intn(100),
					StartDate:               time.Now(),
					EndDate:                 time.Now().AddDate(0, 0, 10),
					TotalDaysApply:          10,
					EndReminderAt:           time.Now().AddDate(0, 0, 10),
					AppInProgressNum:        rand.Intn(100),
					AppAccomplishedNum:      rand.Intn(100),
					MeteringInProgressNum:   rand.Intn(100),
					MeteringAccomplishedNum: rand.Intn(100),
					ResAllocNum:             rand.Intn(100),
					CpuNodesExpected:        rand.Intn(100),
					GpuNodesExpected:        rand.Intn(100),
					StorageSizeExpected:     rand.Intn(100),
					CpuNodesAcquired:        rand.Intn(100),
					GpuNodesAcquired:        rand.Intn(100),
					StorageSizeAcquired:     rand.Intn(100),
					CreatedAt:               time.Now(),
					UpdatedAt:               time.Now(),
				}

				projectID, err := pdb.InsertAllInfo(proSs[i], proDs[i])
				Expect(err).ShouldNot(HaveOccurred(), "InsertAllInfo %d error: %v", i, err)
				By(fmt.Sprintf("InsertAllInfo %d success, projectID=%d", i, projectID))
				proSs[i].ProjectID = projectID
				proDs[i].ProjectID = projectID
			}

			for j := 1; j < 4; j++ {
				psi, err := pdb.QueryStaticInfoByID(j)
				Expect(err).ShouldNot(HaveOccurred(), "QueryStaticInfoByID %d error: %v", j, err)

				Expect(psi.CreatedAt).Should(BeTemporally("~", proSs[j-1].CreatedAt, time.Second))
				Expect(psi.UpdatedAt).Should(BeTemporally("~", proSs[j-1].UpdatedAt, time.Second))
				proSs[j-1].CreatedAt = psi.CreatedAt
				proSs[j-1].UpdatedAt = psi.UpdatedAt

				Expect(psi).Should(Equal(proSs[j-1]))
				By(fmt.Sprintf("project static info: %v", psi))
			}

			for j := 1; j < 4; j++ {
				pdi, err := pdb.QueryDynamicInfoByID(j)
				Expect(err).ShouldNot(HaveOccurred(), "QueryDynamicInfoByID %d error: %v", j, err)

				Expect(pdi.CreatedAt).Should(BeTemporally("~", proDs[j-1].CreatedAt, time.Second))
				Expect(pdi.UpdatedAt).Should(BeTemporally("~", proDs[j-1].UpdatedAt, time.Second))

				Expect(pdi.StartDate).Should(BeTemporally("~", proDs[j-1].StartDate, time.Second))
				Expect(pdi.EndDate).Should(BeTemporally("~", proDs[j-1].EndDate, time.Second))
				Expect(pdi.EndReminderAt).Should(BeTemporally("~", proDs[j-1].EndReminderAt, time.Second))

				proDs[j-1].CreatedAt = pdi.CreatedAt
				proDs[j-1].UpdatedAt = pdi.UpdatedAt
				proDs[j-1].StartDate = pdi.StartDate
				proDs[j-1].EndDate = pdi.EndDate
				proDs[j-1].EndReminderAt = pdi.EndReminderAt

				Expect(pdi).Should(Equal(proDs[j-1]))
				By(fmt.Sprintf("project dynamic info: %v", pdi))
			}
		})
	})

	Describe("Update from db", func() {
		//Query All Project Info from db
		It("UpdateStaticInfo", func() {
			id := 2

			psi, err := pdb.QueryStaticInfoByID(id)
			Expect(err).ShouldNot(HaveOccurred(), "QueryStaticInfoByID %d error: %v", id, err)

			psi.ProjectName = fmt.Sprintf("ProjectNameUpdated%d", id)
			psi.ProjectCode = fmt.Sprintf("ProjectNameUpdated%d", id)
			psi.ExtraInfo = fmt.Sprintf("ProjectNameUpdated%d", id)
			err = pdb.UpdateStaticInfo(psi)
			Expect(err).ShouldNot(HaveOccurred(), "UpdateStaticInfo %d error: %v", id, err)

			psiU, err := pdb.QueryStaticInfoByID(id)
			Expect(err).ShouldNot(HaveOccurred(), "QueryStaticInfoByID %d error: %v", id, err)
			Expect(psiU).Should(Equal(psi))
		})

		It("UpdateStaticInfo", func() {
			id := 3

			pdi, err := pdb.QueryDynamicInfoByID(id)
			Expect(err).ShouldNot(HaveOccurred(), "QueryDynamicInfoByID %d error: %v", id, err)

			pdi.BasicStatus = rand.Intn(100)
			pdi.CpuNodesExpected = rand.Intn(100)

			err = pdb.UpdateDynamicInfo(pdi)
			Expect(err).ShouldNot(HaveOccurred(), "UpdateDynamicInfo %d error: %v", id, err)

			pdiU, err := pdb.QueryDynamicInfoByID(id)
			Expect(err).ShouldNot(HaveOccurred(), "QueryDynamicInfoByID %d error: %v", id, err)
			Expect(pdiU).Should(Equal(pdi))
		})
	})

	Describe("QueryAllInfo from db", func() {
		//Query All Project Info from db
		It("QueryAllInfo from db", func() {
			_, _, err := pdb.QueryAllInfo()
			Expect(err).ShouldNot(HaveOccurred(), "QueryAllInfo error: %v")
		})
	})

})
