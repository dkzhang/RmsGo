package meteringDM_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("MeteringDM", func() {
	BeforeEach(func() {

	})

	Describe("ComputeSettlement", func() {
		It("Compute CPU ", func() {
			projectID := 1

			cpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  10,
				CreatedAt: time.Date(2020, 1, 2, 0, 0, 0, 0, time.Now().Location()),
			})

			cpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  20,
				CreatedAt: time.Date(2020, 1, 4, 0, 0, 0, 0, time.Now().Location()),
			})

			cpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  40,
				CreatedAt: time.Date(2020, 1, 6, 0, 0, 0, 0, time.Now().Location()),
			})

			cpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  20,
				CreatedAt: time.Date(2020, 1, 8, 0, 0, 0, 0, time.Now().Location()),
			})

			cpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  0,
				CreatedAt: time.Date(2020, 2, 1, 0, 0, 0, 0, time.Now().Location()),
			})

			ms, err := mdm.QueryWithCreate(projectID, metering.TypeSettlement, "")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ms.CpuAmountInHours).Should(Equal(14880))
			Expect(ms.CpuAmountInDays).Should(Equal(620))
			By(fmt.Sprintf("CpuAmountInHours=%d, CpuAmountInDays=%d", ms.CpuAmountInHours, ms.CpuAmountInDays))
			By(fmt.Sprintf("metering.Statement=%v", ms))
		})

		It("Compute GPU ", func() {
			projectID := 2

			gpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  10,
				CreatedAt: time.Date(2019, 11, 1, 0, 0, 0, 0, time.Now().Location()),
			})

			gpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  20,
				CreatedAt: time.Date(2020, 1, 4, 0, 0, 0, 0, time.Now().Location()),
			})

			gpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  0,
				CreatedAt: time.Date(2020, 2, 6, 0, 0, 0, 0, time.Now().Location()),
			})

			gpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  30,
				CreatedAt: time.Date(2020, 8, 8, 0, 0, 0, 0, time.Now().Location()),
			})

			gpuDB.Insert(resAlloc.Record{
				ProjectID: projectID,
				NumAfter:  0,
				CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.Now().Location()),
			})

			ms, err := mdm.QueryWithCreate(projectID, metering.TypeSettlement, "")
			Expect(err).ShouldNot(HaveOccurred())
			//Expect(ms.GpuAmountInHours).Should(Equal(158640))
			//Expect(ms.GpuAmountInDays).Should(Equal(6610))
			By(fmt.Sprintf("GpuAmountInHours=%d, GpuAmountInDays=%d", ms.GpuAmountInHours, ms.GpuAmountInDays))
			By(fmt.Sprintf("metering.Statement=%v", ms))
		})
	})
})
