package projectResDM_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("ProjectResDM", func() {
	BeforeEach(func() {

	})

	Describe("Insert ProjectRes Info", func() {
		It("Insert ProjectRes", func() {
			num := 100
			prs := make([]projectRes.ResInfo, num)
			for i := 0; i < num; i++ {
				projectID := i + 1
				prs[i] = projectRes.ResInfo{
					ProjectID:           projectID,
					CpuNodesAcquired:    rand.Intn(1000),
					GpuNodesAcquired:    rand.Intn(1000),
					StorageSizeAcquired: rand.Intn(1000),
					CpuNodesArray:       []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					CpuNodesStr:         fmt.Sprintf("CpuNodesStr %d", rand.Intn(1000)),
					GpuNodesArray:       []int64{1, 2, 3, 4, 5},
					GpuNodesStr:         fmt.Sprintf("GpuNodesStr %d", rand.Intn(1000)),
					StorageAllocInfo:    fmt.Sprintf("StorageAllocInfo %d", rand.Intn(1000)),
					CreatedAt:           time.Now(),
					UpdatedAt:           time.Now(),
				}

				err := prdm.Insert(prs[i])
				Expect(err).ShouldNot(HaveOccurred())
			}

			// insert a projectRes which Project already existed.
			for j := 0; j < 10; j++ {
				err := prdm.Insert(prs[rand.Intn(num)])
				Expect(err).Should(HaveOccurred())
			}

			// query by id and update
			for j := 0; j < 10; j++ {
				projectID := rand.Intn(num) + 1

				ra := projectRes.ResInfo{
					ProjectID:           projectID,
					CpuNodesAcquired:    rand.Intn(1000),
					GpuNodesAcquired:    rand.Intn(1000),
					StorageSizeAcquired: rand.Intn(1000),
					CpuNodesArray:       []int64{2, 4, 6, 8, 10},
					CpuNodesStr:         fmt.Sprintf("CpuNodesStrUpdated %d", rand.Intn(1000)),
					GpuNodesArray:       []int64{1, 3, 5, 7, 9},
					GpuNodesStr:         fmt.Sprintf("GpuNodesStrUpdated %d", rand.Intn(1000)),
					StorageAllocInfo:    fmt.Sprintf("StorageAllocInfoUpdated %d", rand.Intn(1000)),
					UpdatedAt:           time.Now(),
				}
				err := prdm.Update(ra)
				Expect(err).ShouldNot(HaveOccurred())

				raUpdated, err := prdm.QueryByID(projectID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(ra.ProjectID).Should(Equal(raUpdated.ProjectID))
				Expect(ra.CpuNodesAcquired).Should(Equal(raUpdated.CpuNodesAcquired))
				Expect(ra.GpuNodesAcquired).Should(Equal(raUpdated.GpuNodesAcquired))
				Expect(ra.StorageSizeAcquired).Should(Equal(raUpdated.StorageSizeAcquired))
				Expect(ra.CpuNodesArray).Should(Equal(raUpdated.CpuNodesArray))
				Expect(ra.CpuNodesStr).Should(Equal(raUpdated.CpuNodesStr))
				Expect(ra.GpuNodesArray).Should(Equal(raUpdated.GpuNodesArray))
				Expect(ra.GpuNodesStr).Should(Equal(raUpdated.GpuNodesStr))
				Expect(ra.StorageAllocInfo).Should(Equal(raUpdated.StorageAllocInfo))
				Expect(ra.UpdatedAt).Should(BeTemporally("~", raUpdated.UpdatedAt, time.Second))
			}

			// query all
			rsALL, err := prdm.QueryAll()
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("len(records) (ALL) = %d", len(rsALL)))
		})
	})
})
