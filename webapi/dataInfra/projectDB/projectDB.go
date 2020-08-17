package projectDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/jmoiron/sqlx"
)

type DBInfo struct {
	TheDB            *sqlx.DB
	StaticTableName  string
	DynamicTableName string
}

type ProjectDB interface {
	ProjectHistoryDB
	///////////////////////////////////////////////////////////////////////////////
	InsertAllInfo(project.StaticInfo, project.DynamicInfo) (projectID int, err error)
	UpdateStaticInfo(projectInfo project.StaticInfo) (err error)
	UpdateDynamicInfo(projectInfo project.DynamicInfo) (err error)

	// 同一数据库内归档
	InnerArchiveProject(stnHistory string, dtnHistory string, projectID int) (err error)
}

type ProjectHistoryDB interface {
	QueryStaticInfoByID(projectID int) (project.StaticInfo, error)
	QueryDynamicInfoByID(projectID int) (project.DynamicInfo, error)

	QueryInfoByOwner(userID int) ([]project.StaticInfo, []project.DynamicInfo, error)
	QueryInfoByDepartmentCode(dc string) ([]project.StaticInfo, []project.DynamicInfo, error)
	QueryAllInfo() ([]project.StaticInfo, []project.DynamicInfo, error)

	Close()
}
