package projectDB

import "github.com/dkzhang/RmsGo/webapi/model/project"

type ProjectDM interface {
	ProjectHistoryDM
	///////////////////////////////////////////////////////////////////////////////
	InsertAllInfo(project.StaticInfo, project.DynamicInfo) (err error)
	UpdateStaticInfo(projectInfo project.StaticInfo) (err error)
	UpdateDynamicInfo(projectInfo project.DynamicInfo) (err error)

	ArchiveProject(projectID int) (err error)
}

type ProjectHistoryDM interface {
	QueryStaticInfoByID(projectID int) (project.StaticInfo, error)
	QueryDynamicInfoByID(projectID int) (project.DynamicInfo, error)

	QueryInfoByOwner(userID int) ([]project.StaticInfo, []project.DynamicInfo)
	QueryInfoByDepartmentCode(dc string) ([]project.StaticInfo, []project.DynamicInfo)
	QueryAllInfo() ([]project.StaticInfo, []project.DynamicInfo)
	QueryProjectByFilter(userFilter func(project.StaticInfo, project.DynamicInfo) bool) ([]project.StaticInfo, []project.DynamicInfo)
}
