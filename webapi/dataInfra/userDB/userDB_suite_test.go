package userDB_test

import (
	"github.com/dkzhang/RmsGo/databaseInit"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
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
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := databaseInit.ConnectToDatabase()
	pgOps.CreateAllTable(db)

	udb = userDB.NewUserInPostgre(db)
})

var _ = AfterSuite(func() {
	udb.Close()
})
