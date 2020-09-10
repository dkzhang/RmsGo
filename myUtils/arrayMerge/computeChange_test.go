package arrayMerge_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/arrayMerge"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ComputeChange", func() {
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
})
