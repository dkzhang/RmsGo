package userDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/DbManage/PgManage"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	"github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDB"
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

	//GinkgoWriter.Write([]byte(fmt.Sprintf("config.TheDbSecurity = %v \n", config.TheDbSecurity)))
	By(fmt.Sprintf("config.TheDbSecurity = %v \n", security.TheDbSecurity))

	db, err := postgreOps.ConnectToDatabase(security.TheDbConfig.ThePgConfig)
	Expect(err).ShouldNot(HaveOccurred(), "postgreOps.ConnectToDatabase error: %v", err)
	udb = userDB.NewUserInPostgre(db)
})

var _ = AfterSuite(func() {
	udb.Close()
})
