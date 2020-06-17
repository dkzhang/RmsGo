package projectDM

import "github.com/dkzhang/RmsGo/webapi/model/project"

type projectDM interface {
	// What operations can the user do related to the project?

	// 新建项目，提交&审批新项目
	// 删除项目（如果该项目没有被分配过资源）
	// 新建、提交&审批变更申请

	InsertProRes(projectInfo project.ProjectInfo) (err error)
	UpdateProRes(projectInfo project.ProjectInfo) (err error)
	// InsertApplication 新增一个新申请单
	// UpdateApplication 更新一个新申请单
	// InsertApplicationOperation 对申请单的操作：提交

	/*
		项目长保存不提交项目申请单（项目号-1，申请单号-1）时：新增一个项目，新增一个申请单，新增一个申请单动作
		项目长提交项目申请单（项目号m，申请单号n）时：更新项目，更新申请单，新增一个申请单动作
	*/

	// 分配资源

	// 新建、提交&审批计算&存储资源归还申请

	// （系统新建）、审批计量清单
}
