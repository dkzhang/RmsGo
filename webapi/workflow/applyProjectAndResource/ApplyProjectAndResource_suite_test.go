package ApplyProjectAndResource_test

import (
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/workflow"
	"github.com/dkzhang/RmsGo/webapi/workflow/applyProjectAndResource"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var adm applicationDM.ApplicationDM
var pdm projectDM.ProjectDM

var adb applicationDB.ApplicationDB
var pdb projectDB.ProjectDB

var gwf workflow.GeneralWorkflow

func TestApplyProjectAndResource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "applyProjectAndResource Suite")
}

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	theLogMap := logMap.NewLogMap(`C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Parameter\logmap.yaml`)

	adb := applicationDB.NewApplicationPg(db, application.TableApp, application.TableAppOps)
	adm, _ = applicationDM.NewMemoryMap(adb, theLogMap)

	pdb := projectDB.NewProjectPg(db, project.TableName)
	pdm, _ = projectDM.NewMemoryMap(pdb, theLogMap)

	gwf = ApplyProjectAndResource.NewWorkflow(adm, pdm)

})

var _ = AfterSuite(func() {
	//adb.Close()
	//pdb.Close()
})
