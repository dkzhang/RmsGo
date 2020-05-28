package userDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	"github.com/dkzhang/RmsGo/dbManage/pgManage"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userDB"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("UserDB", func() {
	Describe("insert new user with no error", func() {
		Context("", func() {

			os.Setenv("DbConf", "./../../../Configuration/Security/database.yaml")
			pgManage.CreateAllTable()

			//GinkgoWriter.Write([]byte(fmt.Sprintf("config.TheDbConfig = %v \n", config.TheDbConfig)))
			By(fmt.Sprintf("config.TheDbConfig = %v \n", config.TheDbConfig))

			var (
				udb userDB.UserDB
			)

			BeforeEach(func() {
				db, _ := postgreOps.ConnectToDatabase(config.TheDbConfig.ThePgConfig)
				udb = userDB.NewUserInPostgre(db)
			})

			AfterEach(func() {
				udb.Close()
			})

			It("err should be nil", func() {
				err := udb.InsertUser(user.UserInfo{
					UserName:       "zhang001",
					ChineseName:    "张三1",
					Department:     "计服中心",
					DepartmentCode: "JF",
					Section:        "信息技术室",
					Mobile:         "15383021234",
					Role:           1,
					Status:         2,
					Remarks:        "haha",
				})
				Expect(err).Should(BeNil())
			})

			It("err should be nil", func() {
				err := udb.InsertUser(user.UserInfo{
					UserName:       "zhang001",
					ChineseName:    "张三1",
					Department:     "计服中心",
					DepartmentCode: "JF",
					Section:        "信息技术室",
					Mobile:         "15383021234",
					Role:           1,
					Status:         2,
					Remarks:        "haha",
				})
				Expect(err != nil).Should(BeTrue())
			})
		})
	})
})
