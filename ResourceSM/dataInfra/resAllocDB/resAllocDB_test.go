package resAllocDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("ResAllocDB", func() {
	BeforeEach(func() {

	})

	Describe("Insert ResAllocRecord", func() {
		It("Insert ResNodes", func() {
			num := 100
			resNodes := make([]resAlloc.Record, num)
			for i := 0; i < num; i++ {
				resNodes[i] = resAlloc.Record{
					ProjectID:          1 + rand.Intn(10),
					NumBefore:          rand.Intn(100),
					AllocInfoBefore:    []int64{1, 2, 3, 4, 5},
					AllocInfoBeforeStr: "[1, 2, 3, 4, 5]",
					NumAfter:           rand.Intn(100),
					AllocInfoAfter:     []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					AllocInfoAfterStr:  "[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]",
					NumChange:          rand.Intn(100),
					AllocInfoChange:    []int64{6, 7, 8, 9, 10},
					AllocInfoChangeStr: "[6, 7, 8, 9, 10]",
					CtrlID:             1 + rand.Intn(5),
					CtrlChineseName:    fmt.Sprintf("CtrlChineseName%d", rand.Intn(10)),
					CreatedAt:          time.Time{},
				}
				err := radb.Insert(resNodes[i])
				Expect(err).ShouldNot(HaveOccurred())
			}

			for j := 0; j < 10; j++ {
				recordID := rand.Intn(num) + 1
				ra, err := radb.QueryByID(recordID)
				Expect(err).ShouldNot(HaveOccurred())
				By(fmt.Sprintf("record (ID=%d) = %v", recordID, ra))
			}

			for j := 0; j < 10; j++ {
				projectID := j + 1
				rs, err := radb.QueryByProjectID(projectID)
				Expect(err).ShouldNot(HaveOccurred())
				By(fmt.Sprintf("len(records) (projectID=%d) = %d", projectID, len(rs)))
			}

			rsALL, err := radb.QueryAll()
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("len(records) (ALL) = %d", len(rsALL)))
		})
	})
})
