package projectResDM

import (
	"github.com/dkzhang/RmsGo/ResourceSM/gRpcService/client"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
)

type ProjectResGRpc struct {
	pdm     projectDM.ProjectDM
	sClient client.SchedulingClient
}

func NewProjectResGRpc(pdm projectDM.ProjectDM, sc client.SchedulingClient) ProjectResGRpc {
	return ProjectResGRpc{
		pdm:     pdm,
		sClient: sc,
	}
}

func (prg ProjectResGRpc) QueryCpuTreeOccupied(projectID int) (jsonTree string, err error) {
	return prg.sClient.QueryCGpuTree(projectID, typeCPU, typeOccupied)
}

func (prg ProjectResGRpc) QueryCpuTreeAvailable(projectID int) (jsonTree string, err error) {
	return prg.sClient.QueryCGpuTree(projectID, typeCPU, typeAvailable)
}
func (prg ProjectResGRpc) QueryGpuTreeOccupied(projectID int) (jsonTree string, err error) {
	return prg.sClient.QueryCGpuTree(projectID, typeGPU, typeOccupied)
}
func (prg ProjectResGRpc) QueryGpuTreeAvailable(projectID int) (jsonTree string, err error) {
	return prg.sClient.QueryCGpuTree(projectID, typeGPU, typeAvailable)
}

func (prg ProjectResGRpc) QueryProjectRes(projectID int) (pr projectRes.ResInfo, err error) {

}
func (prg ProjectResGRpc) QueryProjectResLite(projectID int) (prl projectRes.ResInfoLite, err error) {

}

const (
	typeCPU = 1
	typeGPU = 2
)

const (
	typeOccupied  = 1
	typeAvailable = 2
	typeAll       = 3
)
