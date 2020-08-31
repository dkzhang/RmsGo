package projectDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDB"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/sirupsen/logrus"
	"time"
)

type MemoryMap struct {
	infoMap      map[int]*project.Info
	theProjectDB projectDB.ProjectDB
}

func NewMemoryMap(pdb projectDB.ProjectDB, theLogMap logMap.LogMap) (nmm MemoryMap, err error) {
	nmm.theProjectDB = pdb
	nmm.infoMap = make(map[int]*project.Info)

	pis, err := nmm.theProjectDB.QueryAllInfo()
	if err != nil {
		return MemoryMap{},
			fmt.Errorf("generate new MemoryMap failed since ProjectDB.QueryAllInfo error: %v", err)
	}
	theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"len(Info Array)": len(pis),
	}).Info("NewMemoryMap ProjectDB.QueryAllInfo success.")

	for _, pi := range pis {
		nmm.infoMap[pi.ProjectID] = &pi
	}

	theLogMap.Log(logMap.NORMAL).Info("NewMemoryMap load data to map success.")

	return nmm, nil
}

func (pdm MemoryMap) QueryByID(projectID int) (project.Info, error) {
	if psi, ok := pdm.infoMap[projectID]; ok {
		return *psi, nil
	} else {
		return project.Info{},
			fmt.Errorf("the project (id = %d) info does not exist", projectID)
	}
}

func (pdm MemoryMap) QueryByOwner(userID int) (pis []project.Info, err error) {
	for _, pi := range pdm.infoMap {
		if pi.ChiefID == userID {
			pis = append(pis, *pi)
		}
	}
	return pis, nil
}

func (pdm MemoryMap) QueryByDepartmentCode(dc string) (pis []project.Info, err error) {
	for _, pi := range pdm.infoMap {
		if pi.DepartmentCode == dc {
			pis = append(pis, *pi)
		}
	}
	return pis, nil
}

func (pdm MemoryMap) QueryAllInfo() (pis []project.Info, err error) {
	for _, pi := range pdm.infoMap {
		pis = append(pis, *pi)
	}
	return pis, nil
}

func (pdm MemoryMap) QueryProjectByFilter(userFilter func(project.Info) bool) (pis []project.Info, err error) {
	for _, pi := range pdm.infoMap {
		if userFilter(*pi) == true {
			pis = append(pis, *pi)
		}
	}
	return pis, nil
}

func (pdm MemoryMap) Insert(pInfo project.Info) (projectID int, err error) {
	pInfo.CreatedAt = time.Now()
	pInfo.UpdatedAt = time.Now()

	//insert in db
	projectID, err = pdm.theProjectDB.Insert(pInfo)
	if err != nil {
		return -1, fmt.Errorf("ProjectDM.Insert error: %v", err)
	}
	pInfo.ProjectID = projectID

	//insert in memoryMap
	pdm.infoMap[projectID] = &pInfo

	return projectID, nil
}

func (pdm MemoryMap) UpdateBasicInfo(bi project.BasicInfo) (err error) {
	if _, ok := pdm.infoMap[bi.ProjectID]; !ok {
		return fmt.Errorf("the project (id = %d) info does not exist", bi.ProjectID)
	}

	bi.UpdatedAt = time.Now()

	// update in DB
	err = pdm.theProjectDB.UpdateBasicInfo(bi)
	if err != nil {
		return fmt.Errorf(" ProjectDB.UpdateBasicInfo error: %v", err)
	}

	// update in MemoryMap
	pdm.infoMap[bi.ProjectID].ProjectName = bi.ProjectName
	pdm.infoMap[bi.ProjectID].ExtraInfo = bi.ExtraInfo
	pdm.infoMap[bi.ProjectID].UpdatedAt = bi.UpdatedAt

	return nil
}

func (pdm MemoryMap) UpdateCodeInfo(pc project.CodeInfo) (err error) {
	if _, ok := pdm.infoMap[pc.ProjectID]; !ok {
		return fmt.Errorf("the project (id = %d) info does not exist", pc.ProjectID)
	}

	pc.UpdatedAt = time.Now()

	// update in DB
	err = pdm.theProjectDB.UpdateCodeInfo(pc)
	if err != nil {
		return fmt.Errorf(" ProjectDB.UpdateCodeInfo error: %v", err)
	}

	// update in MemoryMap
	pdm.infoMap[pc.ProjectID].ProjectCode = pc.ProjectCode
	pdm.infoMap[pc.ProjectID].UpdatedAt = pc.UpdatedAt

	return nil
}

func (pdm MemoryMap) UpdateStatusInfo(si project.StatusInfo) (err error) {
	if _, ok := pdm.infoMap[si.ProjectID]; !ok {
		return fmt.Errorf("the project (id = %d) info does not exist", si.ProjectID)
	}

	si.UpdatedAt = time.Now()

	// update in DB
	err = pdm.theProjectDB.UpdateStatusInfo(si)
	if err != nil {
		return fmt.Errorf(" ProjectDB.UpdateStatusInfo error: %v", err)
	}

	// update in MemoryMap
	pdm.infoMap[si.ProjectID].BasicStatus = si.BasicStatus
	pdm.infoMap[si.ProjectID].ComputingAllocStatus = si.ComputingAllocStatus
	pdm.infoMap[si.ProjectID].StorageAllocStatus = si.StorageAllocStatus
	pdm.infoMap[si.ProjectID].UpdatedAt = si.UpdatedAt

	return nil
}

func (pdm MemoryMap) UpdateApplyInfo(ai project.ApplyInfo) (err error) {
	if _, ok := pdm.infoMap[ai.ProjectID]; !ok {
		return fmt.Errorf("the project (id = %d) info does not exist", ai.ProjectID)
	}

	ai.UpdatedAt = time.Now()

	// update in DB
	err = pdm.theProjectDB.UpdateApplyInfo(ai)
	if err != nil {
		return fmt.Errorf(" ProjectDB.UpdateApplyInfo error: %v", err)
	}

	// update in MemoryMap
	pdm.infoMap[ai.ProjectID].StartDate = ai.StartDate
	pdm.infoMap[ai.ProjectID].TotalDaysApply = ai.TotalDaysApply
	pdm.infoMap[ai.ProjectID].EndReminderAt = ai.EndReminderAt
	pdm.infoMap[ai.ProjectID].CpuNodesExpected = ai.CpuNodesExpected
	pdm.infoMap[ai.ProjectID].GpuNodesExpected = ai.GpuNodesExpected
	pdm.infoMap[ai.ProjectID].StorageSizeExpected = ai.StorageSizeExpected

	pdm.infoMap[ai.ProjectID].UpdatedAt = ai.UpdatedAt

	return nil
}

func (pdm MemoryMap) UpdateAllocInfo(ali project.AllocInfo) (err error) {
	if _, ok := pdm.infoMap[ali.ProjectID]; !ok {
		return fmt.Errorf("the project (id = %d) info does not exist", ali.ProjectID)
	}

	ali.UpdatedAt = time.Now()

	// update in DB
	err = pdm.theProjectDB.UpdateAllocInfo(ali)
	if err != nil {
		return fmt.Errorf(" ProjectDB.UpdateAllocInfo error: %v", err)
	}

	// update in MemoryMap
	pdm.infoMap[ali.ProjectID].CpuNodesAcquired = ali.CpuNodesAcquired
	pdm.infoMap[ali.ProjectID].GpuNodesAcquired = ali.GpuNodesAcquired
	pdm.infoMap[ali.ProjectID].StorageSizeAcquired = ali.StorageSizeAcquired
	pdm.infoMap[ali.ProjectID].CpuNodesMap = ali.CpuNodesMap
	pdm.infoMap[ali.ProjectID].GpuNodesMap = ali.GpuNodesMap
	pdm.infoMap[ali.ProjectID].StorageAllocInfo = ali.StorageAllocInfo

	pdm.infoMap[ali.ProjectID].UpdatedAt = ali.UpdatedAt

	return nil
}

func (pdm MemoryMap) ArchiveProject(projectID int) (err error) {
	err = pdm.theProjectDB.InnerArchiveProject(project.TableHistoryName, projectID)
	if err != nil {
		return fmt.Errorf("ProjectDB.InnerArchiveProject error: %v", err)
	}

	delete(pdm.infoMap, projectID)

	return nil
}
