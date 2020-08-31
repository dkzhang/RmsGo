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

	UpdateBasicInfo(bi project.BasicInfo) (err error)
	UpdateCodeInfo(pc project.CodeInfo) (err error)
	UpdateStatusInfo(si project.StatusInfo) (err error)
	UpdateApplyInfo(ai project.ApplyInfo) (err error)
	UpdateAllocInfo(ali project.AllocInfo) (err error)

	// 同一数据库内归档
	InnerArchiveProject(historyTableName string, projectID int) (err error)
}

type ProjectHistoryDB interface {
	QueryByID(projectID int) (project.Info, error)
	QueryByOwner(userID int) ([]project.Info, error)
	QueryByDepartmentCode(dc string) ([]project.Info, error)
	QueryAllInfo() ([]project.Info, error)

	Close()
}
