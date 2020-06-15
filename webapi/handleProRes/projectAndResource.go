package handleProRes

import "time"

type Application struct {
	ProjectID     int
	ApplicationID int
	Action        int
	BasicInfo     string
	ExtraInfo     string
}

type ProjectInfo struct {
	ID   int
	Name string

	DepartmentCode   string
	Department       string
	ChiefID          int
	ChiefChineseName string

	CreateDateTime  time.Time
	ArchiveDateTime time.Time
}

type ResourceApplication struct {
	CpuNodes    int
	GpuNodes    int
	StorageSize int
	StartDate   time.Time
	EndDate     time.Time
}

type ResourceExpansionApplication struct {
	CpuNodes    int
	GpuNodes    int
	StorageSize int
	EndDate     time.Time
	Type        int
}

const (
	ResExTypeCPU = 1 << iota
	ResExTypeGPU
	ResExTypeStorage
	ResExTypeDate
)
