package resNodeDB_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
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
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	rndb = resNodeDB.NewResNodePg(db, resNode.TableName)
})

var _ = AfterSuite(func() {
	rndb.Close()
})
