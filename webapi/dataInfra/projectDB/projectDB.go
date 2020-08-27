package projectDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/jmoiron/sqlx"
)

type DBInfo struct {
	TheDB     *sqlx.DB
	TableName string
}

type ProjectDB interface {
	ProjectHistoryDB
	///////////////////////////////////////////////////////////////////////////////
	Insert(project.Info) (projectID int, err error)

	UpdateProjectCode(projectID int, projectCode string) (err error)
	UpdateBasicInfo()
	UpdateApplyInfo()
	UpdateAllocInfo()

	// 同一数据库内归档
	InnerArchiveProject(stnHistory string, dtnHistory string, projectID int) (err error)
}

type ProjectHistoryDB interface {
	QueryByID(projectID int) (project.Info, error)

	QueryByOwner(userID int) ([]project.Info, error)
	QueryByDepartmentCode(dc string) ([]project.Info, error)
	QueryAllInfo() ([]project.Info, error)

	Close()
}
