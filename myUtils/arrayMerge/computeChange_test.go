package arrayMerge_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/arrayMerge"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array Merge", func() {
	Describe("ComputeChange", func() {
		It("normal", func() {
			before := []int64{1, 2, 5, 8, 9, 3}
			after := []int64{2, 4, 6, 8, 10, 7}
			change, increased, reduced, err := arrayMerge.ComputeChange(before, after)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("change=%v, increased=%d, reduced=%d, err=%v", change, increased, reduced, err))
		})

		It("before is empty", func() {
			before := []int64{}
			after := []int64{2, 4, 6, 8, 10}
			change, increased, reduced, err := arrayMerge.ComputeChange(before, after)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("change=%v, increased=%d, reduced=%d, err=%v", change, increased, reduced, err))
		})

		It("after is empty", func() {
			before := []int64{1, 2, 5, 8, 9}
			after := []int64{}
			change, increased, reduced, err := arrayMerge.ComputeChange(before, after)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("change=%v, increased=%d, reduced=%d, err=%v", change, increased, reduced, err))
		})

		It("duplicated in before", func() {
			before := []int64{1, 2, 5, 8, 9, 9}
			after := []int64{2, 4, 6, 8, 10}
			change, increased, reduced, err := arrayMerge.ComputeChange(before, after)

			Expect(err).Should(HaveOccurred())
			By(fmt.Sprintf("change=%v, increased=%d, reduced=%d, err=%v", change, increased, reduced, err))
		})

		It("duplicated in after", func() {
			before := []int64{1, 2, 5, 8, 9}
			after := []int64{2, 4, 6, 8, 10, 4}
			change, increased, reduced, err := arrayMerge.ComputeChange(before, after)

			Expect(err).Should(HaveOccurred())
			By(fmt.Sprintf("change=%v, increased=%d, reduced=%d, err=%v", change, increased, reduced, err))
		})

		It("before is nil", func() {
			var before []int64
			before = nil
			after := []int64{1, 2, 5, 8, 9}
			change, increased, reduced, err := arrayMerge.ComputeChange(before, after)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("change=%v, increased=%d, reduced=%d, err=%v", change, increased, reduced, err))
		})

		It("after is nil", func() {
			before := []int64{1, 2, 5, 8, 9}
			var after []int64
			after = nil

			change, increased, reduced, err := arrayMerge.ComputeChange(before, after)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("change=%v, increased=%d, reduced=%d, err=%v", change, increased, reduced, err))
		})

		It("before & after is all nil", func() {
			var before, after []int64
			before, after = nil, nil
			change, increased, reduced, err := arrayMerge.ComputeChange(before, after)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("change=%v, increased=%d, reduced=%d, err=%v", change, increased, reduced, err))
		})
	})

	Describe("ApplyBeforeChange", func() {
		It("normal", func() {
			var before, change, after []int64
			var err error

			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			change = []int64{1, 3, -2, -16}
			after, err = arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))

			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			change = []int64{-12, -16}
			after, err = arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))
		})

		It("before is nil normal", func() {
			var before, change []int64
			before = nil
			change = []int64{12, 16}
			after, err := arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))
		})

		It("before is nil error", func() {
			var before, change []int64
			before = nil
			change = []int64{-12, -16}
			after, err := arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).Should(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))
		})

		It("change is nil normal", func() {
			var before, change []int64
			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			change = nil
			after, err := arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))
		})

		It("before & change is nil ", func() {
			var before, change []int64
			before = nil
			change = nil
			after, err := arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))
		})

		It("duplicate value ", func() {
			var before, change []int64
			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			change = []int64{1, 3, -2, -16, 10}
			after, err := arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).Should(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))
		})

		It("zero value ", func() {
			var before, change, after []int64
			var err error

			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			change = []int64{1, 3, -2, -16, 0}
			after, err = arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).Should(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))

			before = []int64{2, 4, 6, 8, 0, 10, 12, 14, 16}
			change = []int64{1, 3, -2, -16}
			after, err = arrayMerge.ApplyBeforeChange(before, change)

			Expect(err).Should(HaveOccurred())
			By(fmt.Sprintf("before=%v, change=%v, after=%v, err=%v", before, change, after, err))
		})
	})

	Describe("ApplyBeforeSubtract", func() {
		It("normal", func() {
			var before, subtract, after []int64
			var err error

			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			subtract = []int64{14, 8, 4, 6}
			after, err = arrayMerge.ApplyBeforeSubtract(before, subtract)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("before=%v, subtract=%v, after=%v, err=%v", before, subtract, after, err))

			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			subtract = []int64{16, 4}
			after, err = arrayMerge.ApplyBeforeSubtract(before, subtract)

			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("before=%v, subtract=%v, after=%v, err=%v", before, subtract, after, err))
		})

		It("negative value in subtract", func() {
			var before, subtract, after []int64
			var err error

			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			subtract = []int64{14, 8, 2, -10, 6}
			after, err = arrayMerge.ApplyBeforeSubtract(before, subtract)

			Expect(err).Should(HaveOccurred())
			By(fmt.Sprintf("before=%v, subtract=%v, after=%v, err=%v", before, subtract, after, err))
		})

		It("subtract value not exist in before", func() {
			var before, subtract, after []int64
			var err error

			before = []int64{2, 4, 6, 8, 10, 12, 14, 16}
			subtract = []int64{14, 8, 2, 11, 6}
			after, err = arrayMerge.ApplyBeforeSubtract(before, subtract)

			Expect(err).Should(HaveOccurred())
			By(fmt.Sprintf("before=%v, subtract=%v, after=%v, err=%v", before, subtract, after, err))
		})
	})
})
