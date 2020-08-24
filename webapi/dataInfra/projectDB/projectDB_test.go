package projectDB_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectDB", func() {

	Describe("QueryAllInfo from db", func() {

		//Query All Project Info from db
		It("QueryAllInfo from db", func() {
			_, _, err := pdb.QueryAllInfo()
			Expect(err).ShouldNot(HaveOccurred(), "QueryAllInfo error: %v")
		})
	})

	//Describe("QueryDynamicInfoByID", func() {
	//
	//	//Query All Project Info from db
	//	It("QueryDynamicInfoByID", func() {
	//		for i:=1; i<4; i++ {
	//			psi, err := pdb.QueryDynamicInfoByID(i)
	//			Expect(err).ShouldNot(HaveOccurred(), "QueryDynamicInfoByID %d error: %v", i, err)
	//			By(fmt.Sprintf("project dynamic info: %v", psi))
	//		}
	//	})
	//})
	//
	//Describe("QueryStaticInfoByID", func() {
	//
	//	//Query All Project Info from db
	//	It("QueryStaticInfoByID", func() {
	//		for i:=1; i<4; i++ {
	//			psi, err := pdb.QueryStaticInfoByID(i)
	//			Expect(err).ShouldNot(HaveOccurred(), "QueryStaticInfoByID %d error: %v", i, err)
	//			By(fmt.Sprintf("project static info: %v", psi))
	//		}
	//	})
	//})
})
