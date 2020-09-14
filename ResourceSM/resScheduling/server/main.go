package main

import (
	"fmt"
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
	pb "github.com/dkzhang/RmsGo/ResourceSM/resScheduling/grpc"
	"github.com/dkzhang/RmsGo/ResourceSM/resScheduling/resourceScheduling"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

const (
	port = ":50071"
)

func main() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	theLogMap := logMap.NewLogMap(`C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Parameter\logmap.yaml`)

	var prdm projectResDM.ProjectResDM
	var cadm, gadm, sadm resAllocDM.ResAllocDM
	var cndm, gndm resNodeDM.ResNodeDM
	var ctdm, gtdm resGTreeDM.ResGTreeDM
	var err error
	prdm, err = projectResDM.NewMemoryMap(projectResDB.NewProjectResPg(db, projectRes.TableName), theLogMap)
	if err != nil {
		panic(err)
	}

	cadm = resAllocDM.NewResAllocDirectDB(resAllocDB.NewResAllocPg(db, resAlloc.TableNameCPU))
	gadm = resAllocDM.NewResAllocDirectDB(resAllocDB.NewResAllocPg(db, resAlloc.TableNameGPU))
	sadm = resAllocDM.NewResAllocDirectDB(resAllocDB.NewResAllocPg(db, resAlloc.TableNameStorage))

	cndm, err = resNodeDM.NewMemoryMap(resNodeDB.NewResNodePg(db, resNode.TableNameCPU), theLogMap)
	if err != nil {
		panic(err)
	}
	gndm, err = resNodeDM.NewMemoryMap(resNodeDB.NewResNodePg(db, resNode.TableNameGPU), theLogMap)
	if err != nil {
		panic(err)
	}

	// TODO load jsonFileName
	jsonFilename := ""
	ctdm, err = resGTreeDM.NewResGTreeDM(cndm, jsonFilename)
	if err != nil {
		panic(err)
	}
	gtdm, err = resGTreeDM.NewResGTreeDM(gndm, jsonFilename)
	if err != nil {
		panic(err)
	}

	theServer := server{
		TheResScheduling: resourceScheduling.NewResScheduling(prdm, cadm, gadm, sadm, cndm, gndm, ctdm, gtdm),
	}

	//////////////////////////////////////////////////////////////////////////////
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf(" fatal error! failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchedulingServer(s, &theServer)
	fmt.Printf("Begin to serve %v \n", time.Now())
	if err := s.Serve(lis); err != nil {
		log.Printf(" fatal error! failed to serve: %v", err)
	}
}
