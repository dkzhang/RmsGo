package resourceManagement

import "time"

type AllocateRelease struct {
	OpsID             int
	ProjectID         int
	OpsUserID         int
	OpsDatetime       time.Time
	AllocateOrRelease int
	NodeBase64Str     string
	NodeCount         int
}
