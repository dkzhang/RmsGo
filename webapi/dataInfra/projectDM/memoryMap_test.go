package projectDM_test

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
			pros := make([]project.Info, 3)

			for i := 0; i < 3; i++ {
				pros[i] = project.Info{
					//ProjectID:            0,
					ProjectName:      fmt.Sprintf("ProjectName%d", i+1),
					ProjectCode:      fmt.Sprintf("ProjectCode%d", i+1),
					DepartmentCode:   "jf1",
					Department:       "计服中心1",
					ChiefID:          1,
					ChiefChineseName: "项目长张俊",
					ExtraInfo:        fmt.Sprintf("ExtraInfo%d", i+1),

					BasicStatus:          rand.Intn(100),
					ComputingAllocStatus: rand.Intn(100),
					StorageAllocStatus:   rand.Intn(100),

					StartDate:           time.Now(),
					TotalDaysApply:      10,
					EndReminderAt:       time.Now().AddDate(0, 0, 10),
					CpuNodesExpected:    rand.Intn(100),
					GpuNodesExpected:    rand.Intn(100),
					StorageSizeExpected: rand.Intn(100),

					CpuNodesAcquired:    rand.Intn(100),
					GpuNodesAcquired:    rand.Intn(100),
					StorageSizeAcquired: rand.Intn(100),
					CpuNodesMap:         fmt.Sprintf("CpuNodesMap%d", i+1),
					GpuNodesMap:         fmt.Sprintf("GpuNodesMap%d", i+1),
					StorageAllocInfo:    fmt.Sprintf("StorageAllocInfo%d", i+1),

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				projectID, err := pdm.Insert(pros[i])
				Expect(err).ShouldNot(HaveOccurred(), "Insert %d error: %v", i, err)
				By(fmt.Sprintf("Insert %d success, projectID=%d", i, projectID))
				pros[i].ProjectID = projectID
			}

			for j := 1; j < 4; j++ {
				psi, err := pdm.QueryByID(j)
				Expect(err).ShouldNot(HaveOccurred(), "QueryStaticInfoByID %d error: %v", j, err)

				Expect(psi.StartDate).Should(BeTemporally("~", pros[j-1].StartDate, time.Second))
				Expect(psi.EndReminderAt).Should(BeTemporally("~", pros[j-1].EndReminderAt, time.Second))
				Expect(psi.CreatedAt).Should(BeTemporally("~", pros[j-1].CreatedAt, time.Second))
				Expect(psi.UpdatedAt).Should(BeTemporally("~", pros[j-1].UpdatedAt, time.Second))

				pros[j-1].StartDate = psi.StartDate
				pros[j-1].EndReminderAt = psi.EndReminderAt
				pros[j-1].CreatedAt = psi.CreatedAt
				pros[j-1].UpdatedAt = psi.UpdatedAt

				Expect(psi).Should(Equal(pros[j-1]))
				By(fmt.Sprintf("project info: %v", psi))
			}
		})
	})

	Describe("Update from db", func() {
		It("Update Static Info", func() {
			projectID := 2

			pi, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryInfoByID %d error: %v", projectID, err)

			bi := project.BasicInfo{
				ProjectID:   projectID,
				ProjectName: fmt.Sprintf("ProjectName%d-Updated", projectID),
				ExtraInfo:   fmt.Sprintf("ExtraInfo%d-Updated", projectID),
				UpdatedAt:   time.Now(),
			}

			err = pdm.UpdateBasicInfo(bi)
			Expect(err).ShouldNot(HaveOccurred())

			piUpdated, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred())

			// check updated info
			Expect(piUpdated.ProjectID).Should(Equal(bi.ProjectID))
			Expect(piUpdated.ProjectName).Should(Equal(bi.ProjectName))
			Expect(piUpdated.ExtraInfo).Should(Equal(bi.ExtraInfo))

			// check not updated info
			pi.ProjectName = piUpdated.ProjectName
			pi.ExtraInfo = piUpdated.ExtraInfo
			pi.UpdatedAt = piUpdated.UpdatedAt
			Expect(piUpdated).Should(Equal(pi))
		})

		It("Update Code Info", func() {
			projectID := 3

			pi, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryInfoByID %d error: %v", projectID, err)

			ci := project.CodeInfo{
				ProjectID:   projectID,
				ProjectCode: fmt.Sprintf("ProjectCode%d-Updated", projectID),
				UpdatedAt:   time.Now(),
			}

			err = pdm.UpdateCodeInfo(ci)
			Expect(err).ShouldNot(HaveOccurred())

			piUpdated, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred())

			// check updated info
			Expect(piUpdated.ProjectID).Should(Equal(ci.ProjectID))
			Expect(piUpdated.ProjectCode).Should(Equal(ci.ProjectCode))

			// check not updated info
			pi.ProjectCode = piUpdated.ProjectCode
			pi.UpdatedAt = piUpdated.UpdatedAt
			Expect(piUpdated).Should(Equal(pi))
		})

		It("Update Status Info", func() {
			projectID := 3

			pi, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryInfoByID %d error: %v", projectID, err)

			si := project.StatusInfo{
				ProjectID:            projectID,
				BasicStatus:          rand.Intn(100),
				ComputingAllocStatus: rand.Intn(100),
				StorageAllocStatus:   rand.Intn(100),
				UpdatedAt:            time.Now(),
			}

			err = pdm.UpdateStatusInfo(si)
			Expect(err).ShouldNot(HaveOccurred())

			piUpdated, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred())

			// check updated info
			Expect(piUpdated.ProjectID).Should(Equal(si.ProjectID))
			Expect(piUpdated.BasicStatus).Should(Equal(si.BasicStatus))
			Expect(piUpdated.ComputingAllocStatus).Should(Equal(si.ComputingAllocStatus))
			Expect(piUpdated.StorageAllocStatus).Should(Equal(si.StorageAllocStatus))

			// check not updated info
			pi.BasicStatus = piUpdated.BasicStatus
			pi.ComputingAllocStatus = piUpdated.ComputingAllocStatus
			pi.StorageAllocStatus = piUpdated.StorageAllocStatus
			pi.UpdatedAt = piUpdated.UpdatedAt
			Expect(piUpdated).Should(Equal(pi))
		})

		It("Update Apply Info", func() {
			projectID := 1

			pi, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryInfoByID %d error: %v", projectID, err)

			ai := project.ApplyInfo{
				ProjectID:           projectID,
				StartDate:           time.Now().AddDate(0, 1, 0),
				TotalDaysApply:      16,
				EndReminderAt:       time.Now().AddDate(0, 1, 16),
				CpuNodesExpected:    rand.Intn(100),
				GpuNodesExpected:    rand.Intn(100),
				StorageSizeExpected: rand.Intn(100),
				UpdatedAt:           time.Now(),
			}

			err = pdm.UpdateApplyInfo(ai)
			Expect(err).ShouldNot(HaveOccurred())

			piUpdated, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred())

			// check updated info
			Expect(piUpdated.StartDate).Should(BeTemporally("~", ai.StartDate, time.Second))
			Expect(piUpdated.TotalDaysApply).Should(Equal(ai.TotalDaysApply))
			Expect(piUpdated.EndReminderAt).Should(BeTemporally("~", ai.EndReminderAt, time.Second))
			Expect(piUpdated.CpuNodesExpected).Should(Equal(ai.CpuNodesExpected))
			Expect(piUpdated.GpuNodesExpected).Should(Equal(ai.GpuNodesExpected))
			Expect(piUpdated.StorageSizeExpected).Should(Equal(ai.StorageSizeExpected))

			// check not updated info
			pi.StartDate = piUpdated.StartDate
			pi.TotalDaysApply = piUpdated.TotalDaysApply
			pi.EndReminderAt = piUpdated.EndReminderAt
			pi.CpuNodesExpected = piUpdated.CpuNodesExpected
			pi.GpuNodesExpected = piUpdated.GpuNodesExpected
			pi.StorageSizeExpected = piUpdated.StorageSizeExpected
			pi.UpdatedAt = piUpdated.UpdatedAt
			Expect(piUpdated).Should(Equal(pi))
		})

		It("Update Alloc Info", func() {
			projectID := 3

			pi, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred(), "QueryInfoByID %d error: %v", projectID, err)

			ali := project.AllocInfo{
				ProjectID:           projectID,
				CpuNodesAcquired:    rand.Intn(100),
				GpuNodesAcquired:    rand.Intn(100),
				StorageSizeAcquired: rand.Intn(100),
				CpuNodesMap:         fmt.Sprintf("CpuNodesMap%dUpdated", projectID),
				GpuNodesMap:         fmt.Sprintf("GpuNodesMap%dUpdated", projectID),
				StorageAllocInfo:    fmt.Sprintf("StorageAllocInfo%dUpdated", projectID),
				UpdatedAt:           time.Now(),
			}

			err = pdm.UpdateAllocInfo(ali)
			Expect(err).ShouldNot(HaveOccurred())

			piUpdated, err := pdm.QueryByID(projectID)
			Expect(err).ShouldNot(HaveOccurred())

			// check updated info
			Expect(piUpdated.CpuNodesAcquired).Should(Equal(ali.CpuNodesAcquired))
			Expect(piUpdated.GpuNodesAcquired).Should(Equal(ali.GpuNodesAcquired))
			Expect(piUpdated.StorageSizeAcquired).Should(Equal(ali.StorageSizeAcquired))
			Expect(piUpdated.CpuNodesMap).Should(Equal(ali.CpuNodesMap))
			Expect(piUpdated.GpuNodesMap).Should(Equal(ali.GpuNodesMap))
			Expect(piUpdated.StorageAllocInfo).Should(Equal(ali.StorageAllocInfo))

			// check not updated info
			pi.CpuNodesAcquired = piUpdated.CpuNodesAcquired
			pi.GpuNodesAcquired = piUpdated.GpuNodesAcquired
			pi.StorageSizeAcquired = piUpdated.StorageSizeAcquired
			pi.CpuNodesMap = piUpdated.CpuNodesMap
			pi.GpuNodesMap = piUpdated.GpuNodesMap
			pi.StorageSizeExpected = piUpdated.StorageSizeExpected
			pi.StorageAllocInfo = piUpdated.StorageAllocInfo
			pi.UpdatedAt = piUpdated.UpdatedAt
			Expect(piUpdated).Should(Equal(pi))
		})
	})

	Describe("QueryAllInfo from db", func() {
		//Query All Project Info from db
		It("QueryAllInfo from db", func() {
			pis, err := pdm.QueryAllInfo()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(pis)).Should(Equal(3))
		})
	})

})
