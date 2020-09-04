package resAllocDM_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResAllocDM(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResAllocDM Suite")
}

var radb resAllocDB.ResAllocDB
var radm resAllocDM.ResAllocDM

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	radb = resAllocDB.NewResAllocPg(db, resAlloc.TableNameCPU)
	radm = resAllocDM.NewResAllocDirectDB(radb)
})

var _ = AfterSuite(func() {
	radb.Close()
})
