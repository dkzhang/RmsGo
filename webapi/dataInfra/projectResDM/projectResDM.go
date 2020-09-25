package projectResDM

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/webapi/model/user"
)

type ProjectResDM interface {
	ProjectResReadOnlyDM
	///////////////////////////////////////////////////////////////////////////////
	SchedulingCpu(projectID int, nodesAfter []int64, ctrlUserInfo user.UserInfo) (err error)
	SchedulingGpu(projectID int, nodesAfter []int64, ctrlUserInfo user.UserInfo) (err error)
	SchedulingStorage(projectID int, storageSizeAfter int, storageAllocInfoAfter string,
		ctrlUserInfo user.UserInfo) (err error)
}

type ProjectResReadOnlyDM interface {
	QueryCpuTreeOccupied(projectID int, treeFormat int64) (jsonTree string, selected []int64, nodesNum int64, err error)
	QueryCpuTreeAvailable(projectID int, treeFormat int64) (jsonTree string, selected []int64, nodesNum int64, err error)
	QueryGpuTreeOccupied(projectID int, treeFormat int64) (jsonTree string, selected []int64, nodesNum int64, err error)
	QueryGpuTreeAvailable(projectID int, treeFormat int64) (jsonTree string, selected []int64, nodesNum int64, err error)

	QueryProjectRes(projectID int) (pr projectRes.ResInfo, err error)
	QueryProjectResLite(projectID int) (prl projectRes.ResInfoLite, err error)
}
