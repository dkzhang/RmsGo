package projectDB_test

import (
	"github.com/dkzhang/RmsGo/databaseInit"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDB"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProjectDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProjectDB Suite")
}

var pdb projectDB.ProjectDB

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := databaseInit.ConnectToDatabase()
	pgOps.CreateAllTable(db)

	pdb = projectDB.NewProjectPg(db, project.TableStatic, project.TableDynamic)
})

var _ = AfterSuite(func() {
	pdb.Close()
})
