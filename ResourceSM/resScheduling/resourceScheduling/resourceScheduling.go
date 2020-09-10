package resourceScheduling

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/projectResDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resTreeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/myUtils/arrayMerge"
	"github.com/dkzhang/RmsGo/myUtils/nodeEncode"
	"time"
)

type ResScheduling struct {
	prdm projectResDM.ProjectResDM

	cpuAllocDM     resAllocDM.ResAllocDM
	gpuAllocDM     resAllocDM.ResAllocDM
	storageAllocDM resAllocDM.ResAllocDM

	cpuNodeDM resNodeDM.ResNodeDM
	gpuNodeDM resNodeDM.ResNodeDM

	cpuTreeDM resTreeDM.ResTreeDM
	gpuTreeDM resTreeDM.ResTreeDM
}

func NewResScheduling(prdm projectResDM.ProjectResDM,
	cadm resAllocDM.ResAllocDM, gadm resAllocDM.ResAllocDM, sadm resAllocDM.ResAllocDM,
	cndm resNodeDM.ResNodeDM, gndm resNodeDM.ResNodeDM,
	ctdm resTreeDM.ResTreeDM, gtdm resTreeDM.ResTreeDM) ResScheduling {
	return ResScheduling{
		prdm:           prdm,
		cpuAllocDM:     cadm,
		gpuAllocDM:     gadm,
		storageAllocDM: sadm,
		cpuNodeDM:      cndm,
		gpuNodeDM:      gndm,
		cpuTreeDM:      ctdm,
		gpuTreeDM:      gtdm,
	}
}

func (rs ResScheduling) SchedulingCPU(projectID int, nodesAfter []int64, ctrlID int, ctrlCN string) (isFirstAlloc bool, err error) {
	var pr projectRes.ResInfo
	isFirstAlloc = rs.prdm.IsExist(projectID)
	if isFirstAlloc {
		pr, err = rs.prdm.QueryByID(projectID)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf(" QueryByID ProjectResourceInfo (projectID=%d) error: %v", projectID, err)
		}
	} else {
		pr = projectRes.ResInfo{ProjectID: projectID}
	}

	// (1) create the Resource Allocate Record
	nodesChange, increased, reduced, err := arrayMerge.ComputeChange(pr.CpuNodesArray, nodesAfter)
	if err != nil {
		return isFirstAlloc,
			fmt.Errorf("arrayMerge.ComputeChange error: %v", err)
	}

	rar := resAlloc.Record{
		ProjectID:          pr.ProjectID,
		NumBefore:          pr.CpuNodesAcquired,
		AllocInfoBefore:    pr.CpuNodesArray,
		AllocInfoBeforeStr: nodeEncode.IntArrayToBase64Str(pr.CpuNodesArray),
		NumAfter:           len(nodesAfter),
		AllocInfoAfter:     nodesAfter,
		AllocInfoAfterStr:  nodeEncode.IntArrayToBase64Str(nodesAfter),
		NumChange:          increased - reduced,
		AllocInfoChange:    nodesChange,
		AllocInfoChangeStr: nodeEncode.IntArrayToBase64Str(nodesChange),
		CtrlID:             ctrlID,
		CtrlChineseName:    ctrlCN,
	}

	// (2) modify Resource Node alloc info
	nodes := make([]resNode.Node, 0)
	for _, ni := range nodesChange {
		if ni > 0 {
			node, err := rs.cpuNodeDM.QueryByID(ni)
			if err != nil {
				return isFirstAlloc,
					fmt.Errorf("cpuNodeDM.QueryByID (nodeID = %d) error: %v", ni, err)
			}
			node.ProjectID = projectID
			node.AllocatedTime = time.Now()
		} else {
			node, err := rs.cpuNodeDM.QueryByID(-ni)
			if err != nil {
				return isFirstAlloc,
					fmt.Errorf("cpuNodeDM.QueryByID (nodeID = %d) error: %v", -ni, err)
			}
			node.ProjectID = 0
			node.AllocatedTime = time.Time{}
		}
	}

	// (3) modify Project Resource info
	pr.CpuNodesAcquired = rar.NumAfter
	pr.CpuNodesArray = rar.AllocInfoAfter
	pr.CpuNodesStr = rar.AllocInfoAfterStr

	// DM DB Ops
	// (1)
	err = rs.cpuAllocDM.Insert(rar)
	if err != nil {
		return isFirstAlloc,
			fmt.Errorf("allocDM.Insert error: %v", err)
	}

	// (2)
	err = rs.cpuNodeDM.UpdateNodes(nodes)
	if err != nil {
		return isFirstAlloc,
			fmt.Errorf("cpuNodeDM.UpdateNodes error: %v", err)
	}

	// (3)
	if isFirstAlloc {
		err = rs.prdm.Insert(pr)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf("prdm.Update error: %v", err)
		}
	} else {
		err = rs.prdm.Update(pr)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf("prdm.Update error: %v", err)
		}
	}

	return isFirstAlloc, nil
}

func (rs ResScheduling) SchedulingGPU(projectID int, nodesAfter []int64, ctrlID int, ctrlCN string) (isFirstAlloc bool, err error) {
	var pr projectRes.ResInfo
	isFirstAlloc = rs.prdm.IsExist(projectID)
	if isFirstAlloc {
		pr, err = rs.prdm.QueryByID(projectID)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf(" QueryByID ProjectResourceInfo (projectID=%d) error: %v", projectID, err)
		}
	} else {
		pr = projectRes.ResInfo{ProjectID: projectID}
	}

	// (1) create the Resource Allocate Record
	nodesChange, increased, reduced, err := arrayMerge.ComputeChange(pr.GpuNodesArray, nodesAfter)
	if err != nil {
		return isFirstAlloc,
			fmt.Errorf("arrayMerge.ComputeChange error: %v", err)
	}

	rar := resAlloc.Record{
		ProjectID:          pr.ProjectID,
		NumBefore:          pr.GpuNodesAcquired,
		AllocInfoBefore:    pr.GpuNodesArray,
		AllocInfoBeforeStr: nodeEncode.IntArrayToBase64Str(pr.GpuNodesArray),
		NumAfter:           len(nodesAfter),
		AllocInfoAfter:     nodesAfter,
		AllocInfoAfterStr:  nodeEncode.IntArrayToBase64Str(nodesAfter),
		NumChange:          increased - reduced,
		AllocInfoChange:    nodesChange,
		AllocInfoChangeStr: nodeEncode.IntArrayToBase64Str(nodesChange),
		CtrlID:             ctrlID,
		CtrlChineseName:    ctrlCN,
	}

	// (2) modify Resource Node alloc info
	nodes := make([]resNode.Node, 0)
	for _, ni := range nodesChange {
		if ni > 0 {
			node, err := rs.gpuNodeDM.QueryByID(ni)
			if err != nil {
				return isFirstAlloc,
					fmt.Errorf("gpuNodeDM.QueryByID (nodeID = %d) error: %v", ni, err)
			}
			node.ProjectID = projectID
			node.AllocatedTime = time.Now()
		} else {
			node, err := rs.gpuNodeDM.QueryByID(-ni)
			if err != nil {
				return isFirstAlloc,
					fmt.Errorf("gpuNodeDM.QueryByID (nodeID = %d) error: %v", -ni, err)
			}
			node.ProjectID = 0
			node.AllocatedTime = time.Time{}
		}
	}

	// (3) modify Project Resource info
	pr.GpuNodesAcquired = rar.NumAfter
	pr.GpuNodesArray = rar.AllocInfoAfter
	pr.GpuNodesStr = rar.AllocInfoAfterStr

	// DM DB Ops
	// (1)
	err = rs.gpuAllocDM.Insert(rar)
	if err != nil {
		return isFirstAlloc,
			fmt.Errorf("allocDM.Insert error: %v", err)
	}

	// (2)
	err = rs.gpuNodeDM.UpdateNodes(nodes)
	if err != nil {
		return isFirstAlloc,
			fmt.Errorf("gpuNodeDM.UpdateNodes error: %v", err)
	}

	// (3)
	if isFirstAlloc {
		err = rs.prdm.Insert(pr)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf("prdm.Update error: %v", err)
		}
	} else {
		err = rs.prdm.Update(pr)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf("prdm.Update error: %v", err)
		}
	}

	return isFirstAlloc, nil
}

func (rs ResScheduling) SchedulingStorage(projectID int,
	storageSizeAfter int, storageAllocInfoAfter string, ctrlID int, ctrlCN string) (isFirstAlloc bool, err error) {

	var pr projectRes.ResInfo
	isFirstAlloc = rs.prdm.IsExist(projectID)
	if isFirstAlloc {
		pr, err = rs.prdm.QueryByID(projectID)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf(" QueryByID ProjectResourceInfo (projectID=%d) error: %v", projectID, err)
		}
	} else {
		pr = projectRes.ResInfo{ProjectID: projectID}
	}

	// (1) create the Resource Allocate Record

	rar := resAlloc.Record{
		ProjectID:          pr.ProjectID,
		NumBefore:          pr.StorageSizeAcquired,
		AllocInfoBefore:    nil,
		AllocInfoBeforeStr: pr.StorageAllocInfo,
		NumAfter:           storageSizeAfter,
		AllocInfoAfter:     nil,
		AllocInfoAfterStr:  storageAllocInfoAfter,
		NumChange:          storageSizeAfter - pr.StorageSizeAcquired,
		AllocInfoChange:    nil,
		AllocInfoChangeStr: fmt.Sprintf("%s => => => %s", pr.StorageAllocInfo, storageAllocInfoAfter),
		CtrlID:             ctrlID,
		CtrlChineseName:    ctrlCN,
	}

	// (2) modify Resource Node alloc info

	// (3) modify Project Resource info
	pr.StorageSizeAcquired = storageSizeAfter
	pr.StorageAllocInfo = storageAllocInfoAfter

	// DM DB Ops
	// (1)
	err = rs.storageAllocDM.Insert(rar)
	if err != nil {
		return isFirstAlloc,
			fmt.Errorf("allocDM.Insert error: %v", err)
	}

	// (2)

	// (3)
	if isFirstAlloc {
		err = rs.prdm.Insert(pr)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf("prdm.Update error: %v", err)
		}
	} else {
		err = rs.prdm.Update(pr)
		if err != nil {
			return isFirstAlloc,
				fmt.Errorf("prdm.Update error: %v", err)
		}
	}

	return isFirstAlloc, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (rs ResScheduling) QueryCpuTreeAllocated(projectID int) (jsonTree string, err error) {
	return rs.cpuTreeDM.QueryTreeAllocated(projectID)
}

func (rs ResScheduling) QueryCpuTreeIdleAndAllocated(projectID int) (jsonTree string, err error) {
	return rs.cpuTreeDM.QueryTreeIdleAndAllocated(projectID)
}

func (rs ResScheduling) QueryCpuTreeAll() (jsonTree string, err error) {
	return rs.cpuTreeDM.QueryTreeAll()
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (rs ResScheduling) QueryGpuTreeAllocated(projectID int) (jsonTree string, err error) {
	return rs.gpuTreeDM.QueryTreeAllocated(projectID)
}

func (rs ResScheduling) QueryGpuTreeIdleAndAllocated(projectID int) (jsonTree string, err error) {
	return rs.gpuTreeDM.QueryTreeIdleAndAllocated(projectID)
}

func (rs ResScheduling) QueryGpuTreeAll() (jsonTree string, err error) {
	return rs.gpuTreeDM.QueryTreeAll()
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (rs ResScheduling) QueryProjectResByID(projectID int) (projectRes.ResInfo, error) {
	return rs.prdm.QueryByID(projectID)
}

func (rs ResScheduling) QueryProjectResLiteByID(projectID int) (projectRes.ResInfoLite, error) {
	return rs.prdm.QueryLiteByID(projectID)
}
