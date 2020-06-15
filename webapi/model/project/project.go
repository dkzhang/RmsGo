package project

import (
	"github.com/dkzhang/RmsGo/webapi/model/metering"
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"github.com/dkzhang/RmsGo/webapi/model/resourceApplication"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type ProjectInfo struct {
	gorm.Model

	ProjectName string

	DepartmentCode   string
	Department       string
	ChiefID          int
	ChiefChineseName string

	Status int

	ResourceExpected resource.Resource
	ResourceAcquired resource.Resource

	TheApplications    []resourceApplication.Application
	TheResAllocRecords []resource.ResAllocRecord
	TheMeteringList    []metering.MeteringList

	ArchiveDateTime time.Time
}
