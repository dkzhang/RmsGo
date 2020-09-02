package resNodeDB_test

import (
	"github.com/dkzhang/RmsGo/databaseInit"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/resNodeDB"
	"github.com/dkzhang/RmsGo/webapi/model/resource/resNode"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResNodeDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResNodeDB Suite")
}

var rndb resNodeDB.ResNodeDB

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := databaseInit.ConnectToDatabase()
	pgOps.CreateAllTable(db)

	rndb = resNodeDB.NewResNodePg(db, resNode.TableName)
})

var _ = AfterSuite(func() {
	rndb.Close()
})
