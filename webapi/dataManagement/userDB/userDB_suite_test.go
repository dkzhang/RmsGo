package userDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/DbManage/PgManage"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userDB"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserDB Suite")
}

var udb userDB.UserInPostgre

var _ = BeforeSuite(func() {
	os.Setenv("DbConf", "./../../../Configuration/Security/database.yaml")
	PgManage.CreateAllTable()

	//GinkgoWriter.Write([]byte(fmt.Sprintf("config.TheDbConfig = %v \n", config.TheDbConfig)))
	By(fmt.Sprintf("config.TheDbConfig = %v \n", security.TheDbConfig))

	db, err := postgreOps.ConnectToDatabase(security.TheDbConfig.ThePgConfig)
	Expect(err).ShouldNot(HaveOccurred(), "postgreOps.ConnectToDatabase error: %v", err)
	udb = userDB.NewUserInPostgre(db)
})

var _ = AfterSuite(func() {
	udb.Close()
})
