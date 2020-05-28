package userDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserDB", func() {

	var (
		user1 user.UserInfo
		user2 user.UserInfo
		user3 user.UserInfo
		user4 user.UserInfo
	)

	BeforeEach(func() {
		user1 = user.UserInfo{
			UserName:       "zhang001",
			ChineseName:    "张一一",
			Department:     "部门一",
			DepartmentCode: "DEP1",
			Section:        "某科室一",
			Mobile:         "12300001111",
			Role:           user.RoleProjectChief,
			Status:         user.StatusNormal,
			Remarks:        "user1",
		}
		user2 = user.UserInfo{
			UserName:       "zhang001",
			ChineseName:    "张二二",
			Department:     "部门一",
			DepartmentCode: "DEP1",
			Section:        "某科室一",
			Mobile:         "12300002222",
			Role:           user.RoleProjectChief,
			Status:         user.StatusNormal,
			Remarks:        "user2",
		}
		user3 = user.UserInfo{
			UserName:       "zhang003",
			ChineseName:    "张三三",
			Department:     "部门一",
			DepartmentCode: "DEP1",
			Section:        "某科室三",
			Mobile:         "12300003333",
			Role:           user.RoleApprover,
			Status:         user.StatusNormal,
			Remarks:        "user3",
		}
		user4 = user.UserInfo{
			UserName:       "zhang004",
			ChineseName:    "张四四",
			Department:     "部门四",
			DepartmentCode: "DEP4",
			Section:        "某科室四",
			Mobile:         "12300004444",
			Role:           user.RoleApprover,
			Status:         user.StatusNormal,
			Remarks:        "user4",
		}
	})

	Describe("insert new user", func() {
		Context("insert first user", func() {
			It("should success", func() {
				err := udb.InsertUser(user1)
				Expect(err).ShouldNot(HaveOccurred(), "insert first user error", err)
			})
		})
		Context("insert 2nd user", func() {
			It("insert 2nd user with same name of first user should error", func() {
				err := udb.InsertUser(user2)
				Expect(err).Should(HaveOccurred())
				By(fmt.Sprintf("error = %v", err))
			})
		})
		Context("insert 3rd user", func() {
			It("should success", func() {
				err := udb.InsertUser(user3)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("insert 4th user", func() {
			It("should success", func() {
				err := udb.InsertUser(user4)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
