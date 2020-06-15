package resource

import "github.com/jinzhu/gorm"

type ResAllocRecord struct {
	gorm.Model

	AllocResDetails ResourceDetails
	UserID          int
	UserChineseName string
}

type Resource struct {
	CpuNodes    int `json:"cpu_nodes"`
	GpuNodes    int `json:"gpu_nodes"`
	StorageSize int `json:"storage_size"`
}

type ResourceDetails struct {
	TheCpuResDetails     CpuResDetails
	TheGpuResDetails     GpuResDetails
	TheStorageResDetails StorageResDetails
}

type CpuResDetails struct {
	Number     int
	NodeMap    string
	NodeBefore string
	NodeAfter  string
}

type GpuResDetails struct {
	Number     int
	NodeMap    string
	NodeBefore string
	NodeAfter  string
}

type StorageResDetails struct {
	Size       int
	SizeBefore int
	SizeAfter  int
}
