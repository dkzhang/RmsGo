package project

import (
	"time"
)

type StaticInfo struct {
	ProjectID   int
	ProjectName string
	ProjectCode string

	DepartmentCode   string
	Department       string
	ChiefID          int
	ChiefChineseName string
	ExtraInfo        string

	CreatedAt time.Time
}
