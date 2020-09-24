package resNodeDM_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/sirupsen/logrus"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResNodeDM(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResNodeDM Suite")
}

var cpuDB resNodeDB.ResNodeDB
var cpuDM resNodeDM.ResNodeDM

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	cpuDB = resNodeDB.NewResNodePg(db, resNode.TableNameCPU)
	var err error
	cpuDM, err = resNodeDM.NewMemoryMap(cpuDB)
	if err != nil {
		logrus.Fatalf("resNodeDM.NewMemoryMap error: %v", err)
	}
})

var _ = AfterSuite(func() {
	cpuDB.Close()
})
