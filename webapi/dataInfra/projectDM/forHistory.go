package projectDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDB"
	"github.com/dkzhang/RmsGo/webapi/model/project"
)

type ForHistory struct {
	theProjectDB projectDB.ProjectDB
}

func NewForHistory(pdb projectDB.ProjectDB) (nFHis ForHistory, err error) {
	nFHis.theProjectDB = pdb

	return nFHis, nil
}

func (fhis ForHistory) QueryStaticInfoByID(projectID int) (project.StaticInfo, error) {
	psi, err := fhis.theProjectDB.QueryStaticInfoByID(projectID)
	if err != nil {
		return project.StaticInfo{}, fmt.Errorf("ProjectDB QueryStaticInfoByID (id = %d) error", projectID)
	}
	return psi, nil
}

func (fhis ForHistory) QueryDynamicInfoByID(projectID int) (project.DynamicInfo, error) {
	pdi, err := fhis.theProjectDB.QueryDynamicInfoByID(projectID)
	if err != nil {
		return project.DynamicInfo{}, fmt.Errorf("ProjectDB QueryDynamicInfoByID (id = %d) error", projectID)
	}
	return pdi, nil
}

func (fhis ForHistory) QueryInfoByOwner(userID int) (pis []project.ProjectInfo, err error) {
	psis, pdis, err := fhis.theProjectDB.QueryInfoByOwner(userID)
	if err != nil {
		return nil, fmt.Errorf("ProjectDB.QueryInfoByOwner (userID = %d) error", userID)
	}

	pis, err = MergeStaticDynamic(psis, pdis)
	return pis, nil
}

func (fhis ForHistory) QueryInfoByDepartmentCode(dc string) (pis []project.ProjectInfo, err error) {
	psis, pdis, err := fhis.theProjectDB.QueryInfoByDepartmentCode(dc)
	if err != nil {
		return nil, fmt.Errorf("ProjectDB.QueryInfoByDepartmentCode (DepartmentCode = %s) error", dc)
	}

	pis, err = MergeStaticDynamic(psis, pdis)
	return pis, nil
}

func (fhis ForHistory) QueryAllInfo() (pis []project.ProjectInfo, err error) {
	psis, pdis, err := fhis.theProjectDB.QueryAllInfo()
	if err != nil {
		return nil, fmt.Errorf("ProjectDB.QueryAllInfo error")
	}

	pis, err = MergeStaticDynamic(psis, pdis)
	return pis, nil
}

func (fhis ForHistory) QueryProjectByFilter(userFilter func(project.StaticInfo, project.DynamicInfo) bool) (pis []project.ProjectInfo, err error) {
	psis, pdis, err := fhis.theProjectDB.QueryAllInfo()
	if err != nil {
		return nil, fmt.Errorf("ProjectDB.QueryAllInfo error")
	}

	pisAll, err := MergeStaticDynamic(psis, pdis)

	for _, pi := range pisAll {
		if userFilter(pi.TheStaticInfo, pi.TheDynamicInfo) == true {
			pis = append(pis, pi)
		}
	}
	return pis, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////

func MergeStaticDynamic(psis []project.StaticInfo, pdis []project.DynamicInfo) (pis []project.ProjectInfo, err error) {
	var pInfoMap map[int]project.ProjectInfo
	for _, psi := range psis {
		pInfoMap[psi.ProjectID] = project.ProjectInfo{
			ProjectID:      psi.ProjectID,
			TheStaticInfo:  psi,
			TheDynamicInfo: project.DynamicInfo{},
		}
	}

	for _, pdi := range pdis {
		if pi, ok := pInfoMap[pdi.ProjectID]; ok {
			pInfoMap[pdi.ProjectID] = project.ProjectInfo{
				ProjectID:      pdi.ProjectID,
				TheStaticInfo:  pi.TheStaticInfo,
				TheDynamicInfo: pdi,
			}
		} else {
			return nil, fmt.Errorf("project (id = %d) has no dynamic info", pdi.ProjectID)
		}
	}

	pis = make([]project.ProjectInfo, 0, len(pInfoMap))
	for _, v := range pInfoMap {
		pis = append(pis, v)
	}
	return pis, nil
}
