package projectResDM_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/projectResDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/projectResDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProjectResDM(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProjectResDM Suite")
}

var prdb1 projectResDB.ProjectResDB
var prdm projectResDM.ProjectResDM

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	theLogMap := logMap.NewLogMap(`C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Parameter\logmap.yaml`)

	prdb1 = projectResDB.NewProjectResPg(db, projectRes.TableName)

	var err error
	prdm, err = projectResDM.NewMemoryMap(prdb1, theLogMap)
	if err != nil {
		panic(err)
	}
})

var _ = AfterSuite(func() {
	prdb1.Close()
})
