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
	projectStaticInfo  map[int]*project.StaticInfo
	projectDynamicInfo map[int]*project.DynamicInfo
	theProjectDB       projectDB.ProjectDB
}

func NewMemoryMap(pdb projectDB.ProjectDB, theLogMap logMap.LogMap) (nmm MemoryMap, err error) {
	nmm.theProjectDB = pdb
	nmm.projectStaticInfo = make(map[int]*project.StaticInfo)
	nmm.projectDynamicInfo = make(map[int]*project.DynamicInfo)

	psis, pdis, err := nmm.theProjectDB.QueryAllInfo()
	if err != nil {
		return MemoryMap{},
			fmt.Errorf("generate new MemoryMap failed since ProjectDB.QueryAllInfo error: %v", err)
	}
	theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"len(StaticInfo)":  len(psis),
		"len(DynamicInfo)": len(pdis),
	}).Info("NewMemoryMap ProjectDB.QueryAllInfo success.")

	for _, psi := range psis {
		nmm.projectStaticInfo[psi.ProjectID] = &psi
	}

	for _, pdi := range pdis {
		nmm.projectDynamicInfo[pdi.ProjectID] = &pdi
	}

	theLogMap.Log(logMap.NORMAL).Info("NewMemoryMap load data to map success.")

	return nmm, nil
}

func (pdm MemoryMap) QueryByID(projectID int) (project.Info, error) {
	// TODO
	return project.Info{}, fmt.Errorf("need to be done")
}

func (pdm MemoryMap) QueryStaticInfoByID(projectID int) (project.StaticInfo, error) {
	if psi, ok := pdm.projectStaticInfo[projectID]; ok {
		return *psi, nil
	} else {
		return project.StaticInfo{}, fmt.Errorf("the project (id = %d) StaticInfo does not exist", projectID)
	}
}

func (pdm MemoryMap) QueryDynamicInfoByID(projectID int) (project.DynamicInfo, error) {
	if pdi, ok := pdm.projectDynamicInfo[projectID]; ok {
		return *pdi, nil
	} else {
		return project.DynamicInfo{}, fmt.Errorf("the project (id = %d) DynamicInfo does not exist", projectID)
	}
}

func (pdm MemoryMap) QueryInfoByOwner(userID int) (pis []project.ProjectInfo, err error) {
	for pid, psi := range pdm.projectStaticInfo {
		if psi.ChiefID == userID {
			if pdi, ok := pdm.projectDynamicInfo[pid]; ok {
				pis = append(pis, project.ProjectInfo{
					ProjectID:      pid,
					TheStaticInfo:  *psi,
					TheDynamicInfo: *pdi,
				})
			} else {
				return nil, fmt.Errorf("project DynamicInfo (id = %d) does not exist", pid)
			}
		}
	}
	return pis, nil
}

func (pdm MemoryMap) QueryInfoByDepartmentCode(dc string) (pis []project.ProjectInfo, err error) {
	for pid, psi := range pdm.projectStaticInfo {
		if psi.DepartmentCode == dc {
			if pdi, ok := pdm.projectDynamicInfo[pid]; ok {
				pis = append(pis, project.ProjectInfo{
					ProjectID:      pid,
					TheStaticInfo:  *psi,
					TheDynamicInfo: *pdi,
				})
			} else {
				return nil, fmt.Errorf("project DynamicInfo (id = %d) does not exist", pid)
			}
		}
	}
	return pis, nil
}

func (pdm MemoryMap) QueryAllInfo() (pis []project.ProjectInfo, err error) {
	for pid, psi := range pdm.projectStaticInfo {
		if pdi, ok := pdm.projectDynamicInfo[pid]; ok {
			pis = append(pis, project.ProjectInfo{
				ProjectID:      pid,
				TheStaticInfo:  *psi,
				TheDynamicInfo: *pdi,
			})
		} else {
			return nil, fmt.Errorf("project DynamicInfo (id = %d) does not exist", pid)
		}
	}
	return pis, nil
}

func (pdm MemoryMap) QueryProjectByFilter(userFilter func(project.StaticInfo, project.DynamicInfo) bool) (pis []project.ProjectInfo, err error) {
	for pid, psi := range pdm.projectStaticInfo {
		if pdi, ok := pdm.projectDynamicInfo[pid]; ok {
			if userFilter(*psi, *pdi) == true {
				pis = append(pis, project.ProjectInfo{
					ProjectID:      pid,
					TheStaticInfo:  *psi,
					TheDynamicInfo: *pdi,
				})
			}
		} else {
			return nil, fmt.Errorf("project DynamicInfo (id = %d) does not exist", pid)
		}
	}
	return pis, nil
}

func (pdm MemoryMap) InsertAllInfo(pInfo project.ProjectInfo) (projectID int, err error) {
	pInfo.TheStaticInfo.CreatedAt = time.Now()
	pInfo.TheStaticInfo.UpdatedAt = time.Now()

	pInfo.TheDynamicInfo.CreatedAt = time.Now()
	pInfo.TheDynamicInfo.UpdatedAt = time.Now()

	projectID, err = pdm.theProjectDB.InsertAllInfo(pInfo.TheStaticInfo, pInfo.TheDynamicInfo)
	if err != nil {
		return -1, fmt.Errorf("ProjectDB.InsertAllInfo error: %v", err)
	}

	pInfo.TheStaticInfo.ProjectID = projectID
	pInfo.TheDynamicInfo.ProjectID = projectID
	pdm.projectStaticInfo[projectID] = &(pInfo.TheStaticInfo)
	pdm.projectDynamicInfo[projectID] = &(pInfo.TheDynamicInfo)

	return projectID, nil
}
func (pdm MemoryMap) UpdateStaticInfo(psi project.StaticInfo) (err error) {
	psi.UpdatedAt = time.Now()
	err = pdm.theProjectDB.UpdateStaticInfo(psi)
	if err != nil {
		return fmt.Errorf("ProjectDB.UpdateStaticInfo error: %v", err)
	}

	pdm.projectStaticInfo[psi.ProjectID] = &(psi)

	return nil
}
func (pdm MemoryMap) UpdateDynamicInfo(pdi project.DynamicInfo) (err error) {
	pdi.UpdatedAt = time.Now()
	err = pdm.theProjectDB.UpdateDynamicInfo(pdi)
	if err != nil {
		return fmt.Errorf("ProjectDB.UpdateDynamicInfo error: %v", err)
	}

	pdm.projectDynamicInfo[pdi.ProjectID] = &(pdi)
	return nil
}

func (pdm MemoryMap) ArchiveProject(projectID int) (err error) {
	return fmt.Errorf("TODO")
}
