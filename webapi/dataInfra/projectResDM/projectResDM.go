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
	QueryCpuTreeOccupied(projectID int) (jsonTree string, err error)
	QueryCpuTreeAvailable(projectID int) (jsonTree string, err error)
	QueryGpuTreeOccupied(projectID int) (jsonTree string, err error)
	QueryGpuTreeAvailable(projectID int) (jsonTree string, err error)

	QueryProjectRes(projectID int) (pr projectRes.ResInfo, err error)
	QueryProjectResLite(projectID int) (prl projectRes.ResInfoLite, err error)
}
