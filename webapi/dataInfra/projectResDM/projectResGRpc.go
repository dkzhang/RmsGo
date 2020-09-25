package projectResDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/gRpcService/client"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/model/user"
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

func (prg ProjectResGRpc) QueryCpuTreeOccupied(projectID int, treeFormat int64) (
	jsonTree string, selected []int64, nodesNum int64, err error) {
	return prg.sClient.QueryCGpuTree(projectID, typeCPU, typeOccupied, treeFormat)
}

func (prg ProjectResGRpc) QueryCpuTreeAvailable(projectID int, treeFormat int64) (
	jsonTree string, selected []int64, nodesNum int64, err error) {
	return prg.sClient.QueryCGpuTree(projectID, typeCPU, typeAvailable, treeFormat)
}
func (prg ProjectResGRpc) QueryGpuTreeOccupied(projectID int, treeFormat int64) (
	jsonTree string, selected []int64, nodesNum int64, err error) {
	return prg.sClient.QueryCGpuTree(projectID, typeGPU, typeOccupied, treeFormat)
}
func (prg ProjectResGRpc) QueryGpuTreeAvailable(projectID int, treeFormat int64) (
	jsonTree string, selected []int64, nodesNum int64, err error) {
	return prg.sClient.QueryCGpuTree(projectID, typeGPU, typeAvailable, treeFormat)
}

func (prg ProjectResGRpc) QueryProjectRes(projectID int) (pr projectRes.ResInfo, err error) {
	return prg.sClient.QueryProjectRes(projectID)
}
func (prg ProjectResGRpc) QueryProjectResLite(projectID int) (prl projectRes.ResInfoLite, err error) {
	return prg.sClient.QueryProjectResLite(projectID)
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (prg ProjectResGRpc) SchedulingCpu(projectID int, nodesAfter []int64, ctrlUserInfo user.UserInfo) (err error) {
	allocInfo, err := prg.sClient.SchedulingCGpu(projectID, typeCPU, nodesAfter, ctrlUserInfo.UserID, ctrlUserInfo.ChineseName)
	if err != nil {
		return fmt.Errorf("grpc client SchedulingCGpu(projectID=%d, nodesAfter=%v) error: %v",
			projectID, nodesAfter, err)
	}

	err = prg.pdm.UpdateAllocNum(allocInfo)
	if err != nil {
		return fmt.Errorf("ProjectDM.UpdateAllocNum (%v) error: %v", allocInfo, err)
	}

	return nil
}

func (prg ProjectResGRpc) SchedulingGpu(projectID int, nodesAfter []int64, ctrlUserInfo user.UserInfo) (err error) {
	allocInfo, err := prg.sClient.SchedulingCGpu(projectID, typeGPU, nodesAfter, ctrlUserInfo.UserID, ctrlUserInfo.ChineseName)
	if err != nil {
		return fmt.Errorf("grpc client SchedulingCGpu(projectID=%d, nodesAfter=%v) error: %v",
			projectID, nodesAfter, err)
	}

	err = prg.pdm.UpdateAllocNum(allocInfo)
	if err != nil {
		return fmt.Errorf("ProjectDM.UpdateAllocNum (%v) error: %v", allocInfo, err)
	}

	return nil
}

func (prg ProjectResGRpc) SchedulingStorage(projectID int,
	storageSizeAfter int, storageAllocInfoAfter string, ctrlUserInfo user.UserInfo) (err error) {
	allocInfo, err := prg.sClient.SchedulingStorage(projectID, storageSizeAfter, storageAllocInfoAfter, ctrlUserInfo.UserID, ctrlUserInfo.ChineseName)
	if err != nil {
		return fmt.Errorf("grpc client SchedulingCGpu(projectID=%d, storageSizeAfter=%d) error: %v",
			projectID, storageSizeAfter, err)
	}

	err = prg.pdm.UpdateAllocNum(allocInfo)
	if err != nil {
		return fmt.Errorf("ProjectDM.UpdateAllocNum (%v) error: %v", allocInfo, err)
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////
const (
	typeCPU = 1
	typeGPU = 2
)

const (
	typeOccupied  = 1
	typeAvailable = 2
	typeAll       = 3
)
