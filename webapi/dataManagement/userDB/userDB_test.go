package userDB_test

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

type TestDataFile struct {
	Users []user.UserInfo `json:"users"`
}

var _ = Describe("UserDB", func() {

	var (
		testData = TestDataFile{}
	)

	BeforeEach(func() {
		filedata, err := ioutil.ReadFile("userDB_test.json")
		Expect(err).ShouldNot(HaveOccurred())

		err = json.Unmarshal(filedata, &testData)
		Expect(err).ShouldNot(HaveOccurred())

		By(fmt.Sprintf("users = %v", testData.Users))
	})

	Describe("insert new user", func() {
		Context("insert user with no name duplicate", func() {
			It("should success", func() {
				err := udb.InsertUser(testData.Users[0])
				Expect(err).ShouldNot(HaveOccurred(), "insert first user error", err)
			})
			It("should success", func() {
				err := udb.InsertUser(testData.Users[1])
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should success", func() {
				err := udb.InsertUser(testData.Users[2])
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("insert user with name duplicate", func() {
			It("insert user with same name should error", func() {
				testData.Users[3].UserName = testData.Users[0].UserName
				err := udb.InsertUser(testData.Users[3])
				Expect(err).Should(HaveOccurred())
				By(fmt.Sprintf("error = %v", err))
			})
		})

		//Context("xxx", func() {
		//	It("should success", func() {
		//		td := TestDataFile{Users: []user.UserInfo{user1,user2, user3, user4}}
		//		jsonStr, err := json.MarshalIndent(td, "", "    ")
		//		Expect(err).ShouldNot(HaveOccurred())
		//		err = ioutil.WriteFile("userDB_test.json", jsonStr, 0644)
		//		Expect(err).ShouldNot(HaveOccurred())
		//	})
		//})
	})

	Describe("query user by name", func() {
		Context("query user exist", func() {
			It("should success", func() {
				user, err := udb.QueryUserByName(testData.Users[0].UserName)
				Expect(err).ShouldNot(HaveOccurred(), "insert first user error", err)
				Expect(user.Mobile).Should(Equal(testData.Users[0].Mobile))
			})
		})
		Context("query user not exist", func() {
			It("should err", func() {
				_, err := udb.QueryUserByName(testData.Users[3].UserName)
				Expect(err).Should(HaveOccurred(), "insert first user error", err)
				By(fmt.Sprintf("error = %v", err))
			})
		})
	})
})
