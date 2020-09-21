package meteringDM_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/meteringDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/meteringDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDM"
	"github.com/dkzhang/RmsGo/ResourceSM/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMeteringDM(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MeteringDM Suite")
}

var mdb meteringDB.MeteringDB
var mdm meteringDM.MeteringDM

var cpuDB resAllocDB.ResAllocDB
var gpuDB resAllocDB.ResAllocDB
var storageDB resAllocDB.ResAllocDB
var cpuDM resAllocDM.ResAllocDM
var gpuDM resAllocDM.ResAllocDM
var storageDM resAllocDM.ResAllocDM

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	mdb = meteringDB.NewMeteringPg(db, metering.TableName)

	cpuDB = resAllocDB.NewResAllocPg(db, resAlloc.TableNameCPU)
	gpuDB = resAllocDB.NewResAllocPg(db, resAlloc.TableNameGPU)
	storageDB = resAllocDB.NewResAllocPg(db, resAlloc.TableNameStorage)
	cpuDM = resAllocDM.NewResAllocDirectDB(cpuDB)
	gpuDM = resAllocDM.NewResAllocDirectDB(gpuDB)
	storageDM = resAllocDM.NewResAllocDirectDB(storageDB)

	mdm = meteringDM.NewResAllocDirectDB(mdb, cpuDM, gpuDM, storageDM)

})

var _ = AfterSuite(func() {
	mdb.Close()
})
