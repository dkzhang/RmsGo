package projectDM_test

import (
	"github.com/dkzhang/RmsGo/databaseInit"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProjectDM(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProjectDM Suite")
}

var pdm projectDM.ProjectDM

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := databaseInit.ConnectToDatabase()
	pgOps.CreateAllTable(db)

	theLogMap := logMap.NewLogMap(`C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Parameter\logmap.yaml`)

	pdb := projectDB.NewProjectPg(db, project.TableName)
	var err error
	pdm, err = projectDM.NewMemoryMap(pdb, theLogMap)
	if err != nil {
		panic(err)
	}
})

var _ = AfterSuite(func() {
	//pdb.Close()
})
