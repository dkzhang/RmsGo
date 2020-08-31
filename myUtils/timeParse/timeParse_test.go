package timeParse_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/timeParse"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TimeParse", func() {
	It("ParseDateShangHai single", func() {

		timeStr := "2020-8-7"

		t, err := timeParse.ParseDateShangHaiS(timeStr)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(t.Year()).Should(BeNumerically("==", 2020))
		Expect(t.Month()).Should(BeNumerically("==", 8))
		Expect(t.Day()).Should(BeNumerically("==", 7))
	})

	It("ParseDateShangHai double", func() {

		timeStr := "2020-08-27"

		t, err := timeParse.ParseDateShangHai(timeStr)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(t.Year()).Should(BeNumerically("==", 2020))
		Expect(t.Month()).Should(BeNumerically("==", 8))
		Expect(t.Day()).Should(BeNumerically("==", 27))
		By(fmt.Sprintf("time = %v", t))
	})

	It("ParseDateShangHai double error", func() {

		timeStr := "2014-02-29"

		_, err := timeParse.ParseDateShangHai(timeStr)
		Expect(err).Should(HaveOccurred())
		By(fmt.Sprintf("err = %v", err))
	})
})
