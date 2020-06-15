package projectDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/project"
)

type ProjectDB interface {
	QueryProjectByID(projectID int) (project.ProjectInfo, error)
	QueryProjectByOwner(userID int) []project.ProjectInfo
	QueryProjectByDepartmentCode(dc string) []project.ProjectInfo
	QueryProjectByFilter(userFilter func(info project.ProjectInfo) bool) []project.ProjectInfo

	InsertProject(projectInfo project.ProjectInfo) (err error)
}
