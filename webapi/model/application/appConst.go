package application

import "github.com/dkzhang/RmsGo/webapi/model/user"

const (
	AppStatusProjectChief = user.RoleProjectChief
	AppStatusApprover     = user.RoleApprover
	AppStatusApprover2    = user.RoleApprover2
	AppStatusController   = user.RoleController
	AppStatusArchived     = user.RoleSystemArchiver

	AppStatusALL = -1
)

const (
	AppTypeNew = 1 << iota
	AppTypeChange
	AppTypeReturnCompute
	AppTypeReturnStorage

	AppTypeBrowseMetering

	AppTypeALL = -1
)

const (
	AppActionSubmit = 1
	AppActionPass   = 1
	AppActionReject = -1
)
