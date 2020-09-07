package projectResDB_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/projectResDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProjectResDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProjectResDB Suite")
}

var prdb projectResDB.ProjectResDB

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	prdb = projectResDB.NewProjectResPg(db, projectRes.TableName)
})

var _ = AfterSuite(func() {
	prdb.Close()
})
