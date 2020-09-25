package resScheduling_test

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/projectResDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/projectResDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resGTreeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/ResourceSM/resScheduling"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResScheduling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResScheduling Suite")
}

var prdm projectResDM.ProjectResDM
var cadm, gadm, sadm resAllocDM.ResAllocDM
var cndm, gndm resNodeDM.ResNodeDM
var ctdm, gtdm resGTreeDM.ResGTreeDM

var theResScheduling resScheduling.ResScheduling

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)
	pgOpsSqlx.SeedAllTable(db)

	var err error
	prdm, err = projectResDM.NewMemoryMap(projectResDB.NewProjectResPg(db, projectRes.TableName))
	if err != nil {
		panic(err)
	}

	cadm = resAllocDM.NewResAllocDirectDB(resAllocDB.NewResAllocPg(db, resAlloc.TableNameCPU))
	gadm = resAllocDM.NewResAllocDirectDB(resAllocDB.NewResAllocPg(db, resAlloc.TableNameGPU))
	sadm = resAllocDM.NewResAllocDirectDB(resAllocDB.NewResAllocPg(db, resAlloc.TableNameStorage))

	cndm, err = resNodeDM.NewMemoryMap(resNodeDB.NewResNodePg(db, resNode.TableNameCPU))
	if err != nil {
		panic(err)
	}
	gndm, err = resNodeDM.NewMemoryMap(resNodeDB.NewResNodePg(db, resNode.TableNameGPU))
	if err != nil {
		panic(err)
	}

	os.Setenv("TreeJsonCPU", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Parameter\tree_cpu256_8_32.json`)
	os.Setenv("TreeJsonGPU", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Parameter\tree_gpu66_2_33.json`)
	ctdm, err = resGTreeDM.NewResGTreeDM(cndm, os.Getenv("TreeJsonCPU"))
	if err != nil {
		panic(err)
	}

	gtdm, err = resGTreeDM.NewResGTreeDM(gndm, os.Getenv("TreeJsonGPU"))
	if err != nil {
		panic(err)
	}

	theResScheduling = resScheduling.NewResScheduling(prdm, cadm, gadm, sadm, cndm, gndm, ctdm, gtdm)
})

var _ = AfterSuite(func() {

})
