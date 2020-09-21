package meteringDB_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/meteringDB"
	"github.com/dkzhang/RmsGo/ResourceSM/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMeteringDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MeteringDB Suite")
}

var mdb meteringDB.MeteringDB

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	mdb = meteringDB.NewMeteringPg(db, metering.TableName)
})

var _ = AfterSuite(func() {
	mdb.Close()
})
