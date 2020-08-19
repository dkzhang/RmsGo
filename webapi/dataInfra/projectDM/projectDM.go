package projectDM

import "github.com/dkzhang/RmsGo/webapi/model/project"

type ProjectDM interface {
	ProjectHistoryDM
	///////////////////////////////////////////////////////////////////////////////
	InsertAllInfo(project.ProjectInfo) (projectID int, err error)
	UpdateStaticInfo(psi project.StaticInfo) (err error)
	UpdateDynamicInfo(pdi project.DynamicInfo) (err error)

	ArchiveProject(projectID int) (err error)
}

type ProjectHistoryDM interface {
	QueryStaticInfoByID(projectID int) (project.StaticInfo, error)
	QueryDynamicInfoByID(projectID int) (project.DynamicInfo, error)

	QueryInfoByOwner(userID int) ([]project.ProjectInfo, error)
	QueryInfoByDepartmentCode(dc string) ([]project.ProjectInfo, error)
	QueryAllInfo() ([]project.ProjectInfo, error)
	QueryProjectByFilter(userFilter func(project.StaticInfo, project.DynamicInfo) bool) ([]project.ProjectInfo, error)
}
