package projectDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/resource"
)

type ProjectDB interface {
	QueryProjectByID(projectID int) (project.ProjectInfo, error)
	QueryProjectByOwner(userID int) []project.ProjectInfo
	QueryProjectByDepartmentCode(dc string) []project.ProjectInfo
	QueryProjectByFilter(userFilter func(info project.ProjectInfo) bool) []project.ProjectInfo

	///////////////////////////////////////////////////////////////////////////////
	InsertProject(projectInfo project.ProjectInfo) (err error)
	UpdateProject(projectInfo project.ProjectInfo) (err error)
	DeleteProject(projectID int) (err error)
	// InsertApplication 新增一个新申请单
	// UpdateApplication 更新一个新申请单
	// InsertApplicationOperation 对申请单的操作：提交

	/*
		项目长保存不提交项目申请单（项目号-1，申请单号-1）时：新增一个项目，新增一个申请单，新增一个申请单动作
		项目长提交项目申请单（项目号m，申请单号n）时：更新项目，更新申请单，新增一个申请单动作
	*/

	InsertResAllocRecord(record resource.ResAllocRecord) (err error)
}
