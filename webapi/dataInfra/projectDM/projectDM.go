package projectDM

import (
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"time"
)

type ProjectDM interface {
	ProjectReadOnlyDM
	///////////////////////////////////////////////////////////////////////////////
	Insert(project.Info) (projectID int, err error)

	UpdateBasicInfo(bi project.BasicInfo) (err error)
	UpdateCodeInfo(pc project.CodeInfo) (err error)
	UpdateStatusInfo(si project.StatusInfo) (err error)
	UpdateApplyInfo(ai project.ApplyInfo) (err error)
	UpdateAllocNum(ali project.AllocNum) (err error)

	//TODO
	//UpdateAllocInfo(ali project.AllocInfo) (err error)

	ArchiveProject(projectID int) (err error)
}

type ProjectReadOnlyDM interface {
	QueryByID(projectID int) (project.Info, error)
	QueryByOwner(userID int) ([]project.Info, error)
	QueryByDepartmentCode(dc string) ([]project.Info, error)
	QueryAllInfo() ([]project.Info, error)
	QueryProjectByFilter(userFilter func(project.Info) bool) ([]project.Info, error)

	QueryIDsByTimeRange(from, to time.Time) ([]int, error)
}
