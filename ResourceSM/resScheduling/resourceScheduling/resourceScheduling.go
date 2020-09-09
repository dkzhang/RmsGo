package resourceScheduling

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/projectResDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resTreeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/myUtils/arrayMerge"
	"github.com/dkzhang/RmsGo/myUtils/nodeEncode"
	"time"
)

type ResScheduling struct {
	prdm      projectResDM.ProjectResDM
	allocDM   resAllocDM.ResAllocDM
	cpuNodeDM resNodeDM.ResNodeDM
	gpuNodeDM resNodeDM.ResNodeDM
	cpuTreeDM resTreeDM.ResTreeDM
	gpuTreeDM resTreeDM.ResTreeDM
}

func NewResScheduling(prdm projectResDM.ProjectResDM,
	radm resAllocDM.ResAllocDM,
	cndm resNodeDM.ResNodeDM, gndm resNodeDM.ResNodeDM,
	ctdm resTreeDM.ResTreeDM, gtdm resTreeDM.ResTreeDM) ResScheduling {
	return ResScheduling{
		prdm:      prdm,
		allocDM:   radm,
		cpuNodeDM: cndm,
		gpuNodeDM: gndm,
		cpuTreeDM: ctdm,
		gpuTreeDM: gtdm,
	}
}

func (rs ResScheduling) SchedulingCPU(projectID int, nodesAfter []int64, ctrlID int, ctrlCN string) (err error) {

	pr, err := rs.prdm.QueryByID(projectID)
	if err != nil {
		return fmt.Errorf("cannot find ProjectResource info with projectID = %d", projectID)
	}

	// (1) create the Resource Allocate Record
	nodesChange, increased, reduced, err := arrayMerge.ComputeChange(pr.CpuNodesArray, nodesAfter)
	if err != nil {
		return fmt.Errorf("arrayMerge.ComputeChange error: %v", err)
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
				return fmt.Errorf("cpuNodeDM.QueryByID (nodeID = %d) error: %v", ni, err)
			}
			node.ProjectID = projectID
			node.AllocatedTime = time.Now()
		} else {
			node, err := rs.cpuNodeDM.QueryByID(-ni)
			if err != nil {
				return fmt.Errorf("cpuNodeDM.QueryByID (nodeID = %d) error: %v", -ni, err)
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
	err = rs.allocDM.Insert(rar)
	if err != nil {
		return fmt.Errorf("allocDM.Insert error: %v", err)
	}

	// (2)
	err = rs.cpuNodeDM.UpdateNodes(nodes)
	if err != nil {
		return fmt.Errorf("cpuNodeDM.UpdateNodes error: %v", err)
	}

	// (3)
	err = rs.prdm.Update(pr)
	if err != nil {
		return fmt.Errorf("prdm.Update error: %v", err)
	}

	return nil
}

func (rs ResScheduling) SchedulingGPU(projectID int, nodesAfter []int64, ctrlID int, ctrlCN string) (err error) {

	pr, err := rs.prdm.QueryByID(projectID)
	if err != nil {
		return fmt.Errorf("cannot find ProjectResource info with projectID = %d", projectID)
	}

	// (1) create the Resource Allocate Record
	nodesChange, increased, reduced, err := arrayMerge.ComputeChange(pr.CpuNodesArray, nodesAfter)
	if err != nil {
		return fmt.Errorf("arrayMerge.ComputeChange error: %v", err)
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
			node, err := rs.gpuNodeDM.QueryByID(ni)
			if err != nil {
				return fmt.Errorf("gpuNodeDM.QueryByID (nodeID = %d) error: %v", ni, err)
			}
			node.ProjectID = projectID
			node.AllocatedTime = time.Now()
		} else {
			node, err := rs.gpuNodeDM.QueryByID(-ni)
			if err != nil {
				return fmt.Errorf("gpuNodeDM.QueryByID (nodeID = %d) error: %v", -ni, err)
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
	err = rs.allocDM.Insert(rar)
	if err != nil {
		return fmt.Errorf("allocDM.Insert error: %v", err)
	}

	// (2)
	err = rs.gpuNodeDM.UpdateNodes(nodes)
	if err != nil {
		return fmt.Errorf("gpuNodeDM.UpdateNodes error: %v", err)
	}

	// (3)
	err = rs.prdm.Update(pr)
	if err != nil {
		return fmt.Errorf("prdm.Update error: %v", err)
	}

	return nil
}

func (rs ResScheduling) SchedulingStorage(projectID int,
	storageSizeAfter int, storageAllocInfoAfter string, ctrlID int, ctrlCN string) (err error) {

	// TODO
	return fmt.Errorf("not accomplished")
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (rs ResScheduling) QueryCpuTreeAllocated(projectID int) (jsonTree string, err error) {
	jsonTree, err = rs.cpuTreeDM.QueryTreeAllocated(projectID)
	if err != nil {
		return "", fmt.Errorf("rs.cpuTreeDM.QueryTreeAllocated error: %v", err)
	}

	return jsonTree, nil
}

func (rs ResScheduling) QueryCpuTreeIdleAndAllocated(projectID int) (jsonTree string, err error) {
	jsonTree, err = rs.cpuTreeDM.QueryTreeIdleAndAllocated(projectID)
	if err != nil {
		return "", fmt.Errorf("rs.cpuTreeDM.QueryTreeIdleAndAllocated error: %v", err)
	}

	return jsonTree, nil
}

func (rs ResScheduling) QueryCpuTreeAll() (jsonTree string, err error) {
	jsonTree, err = rs.cpuTreeDM.QueryTreeAll()
	if err != nil {
		return "", fmt.Errorf("rs.cpuTreeDM.QueryTreeAll error: %v", err)
	}

	return jsonTree, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (rs ResScheduling) QueryGpuTreeAllocated(projectID int) (jsonTree string, err error) {

	// TODO
	return "", fmt.Errorf("not accomplished")
}

func (rs ResScheduling) QueryGpuTreeIdleAndAllocated(projectID int) (jsonTree string, err error) {

	// TODO
	return "", fmt.Errorf("not accomplished")
}

func (rs ResScheduling) QueryGpuTreeAll() (jsonTree string, err error) {

	// TODO
	return "", fmt.Errorf("not accomplished")
}
