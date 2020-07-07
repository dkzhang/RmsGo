package projectDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/project"
)

type ProjectDB interface {
	QueryStaticInfoByID(projectID int) (project.StaticInfo, error)
	QueryDynamicInfoByID(projectID int) (project.DynamicInfo, error)

	QueryAllInfoByOwner(userID int) ([]project.StaticInfo, []project.DynamicInfo)
	QueryAllInfoByDepartmentCode(dc string) ([]project.StaticInfo, []project.DynamicInfo)
	QueryProjectByFilter(userFilter func(project.StaticInfo, project.DynamicInfo) bool) ([]project.StaticInfo, []project.DynamicInfo)

	///////////////////////////////////////////////////////////////////////////////
	InsertAllInfo(project.StaticInfo, project.DynamicInfo) (err error)
	UpdateStaticInfo(projectInfo project.StaticInfo) (err error)
	UpdateDynamicInfo(projectInfo project.DynamicInfo) (err error)

	DeleteProject(projectID int) (err error)
}
