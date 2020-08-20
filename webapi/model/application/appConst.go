package application

import "github.com/dkzhang/RmsGo/webapi/model/user"

const (
	AppStatusProjectChief = user.RoleProjectChief
	AppStatusApprover     = user.RoleApprover
	AppStatusController   = user.RoleController
	AppStatusArchived     = 16

	AppStatusALL = -1
)

const (
	AppTypeNew = 1 << iota
	AppTypeChange
	AppTypeReturnCompute
	AppTypeReturnStorage

	AppTypeALL = -1
)

const (
	AppActionSubmit = 1
	AppActionPass   = 1
	AppActionReject = -1
)
