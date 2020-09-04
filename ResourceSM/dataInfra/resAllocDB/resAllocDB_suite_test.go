package resAllocDB_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResAllocDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResAllocDB Suite")
}

var radb resAllocDB.ResAllocDB

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	radb = resAllocDB.NewResAllocPg(db, resAlloc.TableNameCPU)
})

var _ = AfterSuite(func() {
	radb.Close()
})
