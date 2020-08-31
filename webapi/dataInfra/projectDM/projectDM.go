package projectDM

import "github.com/dkzhang/RmsGo/webapi/model/project"

type ProjectDM interface {
	ProjectHistoryDM
	///////////////////////////////////////////////////////////////////////////////
	Insert(project.Info) (projectID int, err error)

	UpdateBasicInfo(bi project.BasicInfo) (err error)
	UpdateCodeInfo(pc project.CodeInfo) (err error)
	UpdateStatusInfo(si project.StatusInfo) (err error)
	UpdateApplyInfo(ai project.ApplyInfo) (err error)
	UpdateAllocInfo(ali project.AllocInfo) (err error)

	ArchiveProject(projectID int) (err error)
}

type ProjectHistoryDM interface {
	QueryByID(projectID int) (project.Info, error)
	QueryByOwner(userID int) ([]project.Info, error)
	QueryByDepartmentCode(dc string) ([]project.Info, error)
	QueryAllInfo() ([]project.Info, error)
	QueryProjectByFilter(userFilter func(project.Info) bool) ([]project.Info, error)
}
