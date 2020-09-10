package projectResDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/projectResDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/sirupsen/logrus"
	"time"
)

type MemoryMap struct {
	infoMap         map[int]*projectRes.ResInfo
	theProjectResDB projectResDB.ProjectResDB
}

func NewMemoryMap(pdb projectResDB.ProjectResDB, theLogMap logMap.LogMap) (nmm MemoryMap, err error) {
	nmm.theProjectResDB = pdb
	nmm.infoMap = make(map[int]*projectRes.ResInfo)

	prs, err := nmm.theProjectResDB.QueryAll()
	if err != nil {
		return MemoryMap{},
			fmt.Errorf("generate new MemoryMap failed since ProjectResDB.QueryAll error: %v", err)
	}
	theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"len(Info Array)": len(prs),
	}).Info("NewMemoryMap ProjectResDB.QueryAll success.")

	for _, pi := range prs {
		nmm.infoMap[pi.ProjectID] = &pi
	}

	theLogMap.Log(logMap.NORMAL).Info("NewMemoryMap load data to map success.")

	return nmm, nil
}

func (pdm MemoryMap) QueryByID(projectID int) (projectRes.ResInfo, error) {
	if psi, ok := pdm.infoMap[projectID]; ok {
		return *psi, nil
	} else {
		return projectRes.ResInfo{},
			fmt.Errorf("the project res (id = %d) info does not exist", projectID)
	}
}

func (pdm MemoryMap) QueryLiteByID(projectID int) (projectRes.ResInfoLite, error) {
	if psi, ok := pdm.infoMap[projectID]; ok {
		return projectRes.ResInfoLite{
			ProjectID:           psi.ProjectID,
			CpuNodesAcquired:    psi.CpuNodesAcquired,
			GpuNodesAcquired:    psi.GpuNodesAcquired,
			StorageSizeAcquired: psi.StorageSizeAcquired,
		}, nil
	} else {
		return projectRes.ResInfoLite{},
			fmt.Errorf("the project res (id = %d) info does not exist", projectID)
	}
}

func (pdm MemoryMap) QueryAll() (pis []projectRes.ResInfo, err error) {
	for _, pi := range pdm.infoMap {
		pis = append(pis, *pi)
	}
	return pis, nil
}

func (pdm MemoryMap) Insert(pr projectRes.ResInfo) (err error) {
	pr.CreatedAt = time.Now()
	pr.UpdatedAt = time.Now()

	// avoid slice to be nil
	if pr.CpuNodesArray == nil {
		pr.CpuNodesArray = make([]int64, 0)
	}
	if pr.GpuNodesArray == nil {
		pr.GpuNodesArray = make([]int64, 0)
	}

	//insert in db
	err = pdm.theProjectResDB.Insert(pr)
	if err != nil {
		return fmt.Errorf("ProjectResDB.Insert error: %v", err)
	}

	//insert in memoryMap
	pdm.infoMap[pr.ProjectID] = &pr

	return nil
}

func (pdm MemoryMap) Update(pr projectRes.ResInfo) (err error) {
	if _, ok := pdm.infoMap[pr.ProjectID]; !ok {
		return fmt.Errorf("the project res (id = %d) info does not exist", pr.ProjectID)
	}

	pr.UpdatedAt = time.Now()
	// avoid slice to be nil
	if pr.CpuNodesArray == nil {
		pr.CpuNodesArray = make([]int64, 0)
	}
	if pr.GpuNodesArray == nil {
		pr.GpuNodesArray = make([]int64, 0)
	}

	// update in DB
	err = pdm.theProjectResDB.Update(pr)
	if err != nil {
		return fmt.Errorf(" ProjectResDB.Update error: %v", err)
	}

	// update in MemoryMap
	pdm.infoMap[pr.ProjectID] = &pr

	return nil
}

func (pdm MemoryMap) Delete(projectID int) (err error) {
	if _, ok := pdm.infoMap[projectID]; !ok {
		return fmt.Errorf("the project res (id = %d) info does not exist", projectID)
	}

	// delete in DB
	err = pdm.theProjectResDB.Delete(projectID)
	if err != nil {
		return fmt.Errorf(" ProjectResDB.Delete error: %v", err)
	}

	// update in MemoryMap
	delete(pdm.infoMap, projectID)

	return nil
}
