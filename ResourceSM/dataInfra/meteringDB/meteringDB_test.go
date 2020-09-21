package meteringDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"github.com/dkzhang/RmsGo/myUtils/timeParse"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("MeteringDB", func() {
	BeforeEach(func() {

	})

	Describe("MeteringDB", func() {
		It("Insert ", func() {
			mss := make([]metering.Statement, 6)
			mss[0] = metering.Statement{
				ProjectID:    1,
				MeteringType: metering.TypeSettlement,
			}

			mss[1] = metering.Statement{
				ProjectID:    2,
				MeteringType: metering.TypeSettlement,
			}

			mss[2] = metering.Statement{
				//MeteringID:           0,
				ProjectID:            3,
				MeteringType:         metering.TypeSettlement,
				MeteringTypeInfo:     "",
				CpuAmountInDays:      0,
				GpuAmountInDays:      0,
				StorageAmountInDays:  0,
				CpuAmountInHours:     rand.Intn(1000),
				GpuAmountInHours:     rand.Intn(1000),
				StorageAmountInHours: rand.Intn(1000),
				CpuNodeMeteringJson:  "",
				GpuNodeMeteringJson:  "",
				StorageMeteringJson:  "",
				CreatedAt:            time.Now(),
			}
			mss[2].CpuAmountInDays = timeParse.HoursToDays(mss[2].CpuAmountInHours)
			mss[2].GpuAmountInDays = timeParse.HoursToDays(mss[2].GpuAmountInHours)
			mss[2].StorageAmountInDays = timeParse.HoursToDays(mss[2].StorageAmountInHours)

			mss[3] = metering.Statement{
				ProjectID:        1,
				MeteringType:     metering.TypeMonthly,
				MeteringTypeInfo: "202009",
			}

			mss[4] = metering.Statement{
				ProjectID:        1,
				MeteringType:     metering.TypeAnnual,
				MeteringTypeInfo: "2020",
			}

			mss[5] = metering.Statement{
				ProjectID:        2,
				MeteringType:     metering.TypeQuarterly,
				MeteringTypeInfo: "2020Q2",
			}

			for i := 0; i < 6; i++ {
				mid, err := mdb.Insert(mss[i])
				Expect(err).ShouldNot(HaveOccurred())
				By(fmt.Sprintf("Insert ms (id=%d) success", mid))
			}

			ms2, err := mdb.Query(3, metering.TypeSettlement, "")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ms2.CreatedAt).Should(BeTemporally("~", mss[2].CreatedAt, time.Second))
			mss[2].CreatedAt = ms2.CreatedAt
			mss[2].MeteringID = ms2.MeteringID

			Expect(ms2).Should(Equal(mss[2]))

			mssPro1, err := mdb.QueryAll(1, -1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(mssPro1)).Should(Equal(3))

			mssPro1S, err := mdb.QueryAll(1, metering.TypeSettlement)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(mssPro1S)).Should(Equal(1))
		})
	})
})
