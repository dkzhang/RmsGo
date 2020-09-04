package applicationDB_test

import (
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"github.com/dkzhang/RmsGo/webapi/model/application"
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
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	adb = applicationDB.NewApplicationPg(db, application.TableApp, application.TableAppOps)
})

var _ = AfterSuite(func() {
	adb.Close()
})
