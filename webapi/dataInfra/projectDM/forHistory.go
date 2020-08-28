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

func (fhis ForHistory) QueryByID(projectID int) (project.Info, error) {
	return fhis.theProjectDB.QueryByID(projectID)
}

func (fhis ForHistory) QueryByOwner(userID int) (pis []project.Info, err error) {
	return fhis.theProjectDB.QueryByOwner(userID)
}

func (fhis ForHistory) QueryByDepartmentCode(dc string) (pis []project.Info, err error) {
	return fhis.theProjectDB.QueryByDepartmentCode(dc)
}

func (fhis ForHistory) QueryAllInfo() (pis []project.Info, err error) {
	return fhis.theProjectDB.QueryAllInfo()
}

func (fhis ForHistory) QueryProjectByFilter(userFilter func(project.Info) bool) (pis []project.Info, err error) {
	pisDb, err := fhis.theProjectDB.QueryAllInfo()
	if err != nil {
		return nil, fmt.Errorf("ProjectDB.QueryAllInfo error: %v", err)
	}
	for _, pi := range pisDb {
		if userFilter(pi) == true {
			pis = append(pis, pi)
		}
	}
	return pis, nil
}
