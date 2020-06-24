package project

import (
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

	CreatedAt time.Time
}
