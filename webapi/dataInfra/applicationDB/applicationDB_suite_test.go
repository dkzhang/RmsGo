package applicationDB_test

import (
	"github.com/dkzhang/RmsGo/databaseInit"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApplicationDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApplicationDB Suite")
}

var adb applicationDB.ApplicationDB

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := databaseInit.ConnectToDatabase()
	pgOps.CreateAllTable(db)

	adb = applicationDB.NewApplicationPg(db, "application", "app_ops_record")
})

var _ = AfterSuite(func() {
	adb.Close()
})
