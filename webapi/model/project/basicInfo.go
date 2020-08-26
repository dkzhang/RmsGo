package project

import "time"

type BasicInfo struct {
	ProjectID        int    `db:"project_id" json:"project_id"`
	ProjectName      string `db:"project_name" json:"project_name"`
	ProjectCode      string `db:"project_code" json:"project_code"`
	DepartmentCode   string `db:"department_code" json:"department_code"`
	Department       string `db:"department" json:"department"`
	ChiefID          int    `db:"chief_id" json:"chief_id"`
	ChiefChineseName string `db:"chief_cn_name" json:"chief_cn_name"`

	BasicStatus          int `db:"basic_status" json:"basic_status"`
	ComputingAllocStatus int `db:"computing_alloc_status" json:"computing_alloc_status"`
	StorageAllocStatus   int `db:"storage_alloc_status" json:"storage_alloc_status"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
