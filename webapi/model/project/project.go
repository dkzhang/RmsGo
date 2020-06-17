package project

import (
	"github.com/dkzhang/RmsGo/webapi/model/metering"
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"github.com/dkzhang/RmsGo/webapi/model/resourceApplication"
	"time"
)

type ProjectInfo struct {
	ProjectID   int
	ProjectName string

	DepartmentCode   string
	Department       string
	ChiefID          int
	ChiefChineseName string
	ExtraInfo        string

	Status int

	ResourceExpected resource.Resource
	ResourceAcquired resource.Resource

	TheApplications    []resourceApplication.Application
	TheResAllocRecords []resource.ResAllocRecord
	TheMeteringList    []metering.MeteringList

	ArchivedAt time.Time
	CreatedAt  time.Time
}
