package metering

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MeteringList struct {
	gorm.Model

	CpuAmountInDays     int
	GpuAmountInDays     int
	StorageAmountInDays int

	CpuAmountInHours     int
	GpuAmountInHours     int
	StorageAmountInHours int

	CpuNodeMeteringList []CpuNodeMeteringItem
	GpuNodeMeteringList []GpuNodeMeteringItem
	StorageMeteringList []StorageMeteringItem
}

type CpuNodeMeteringItem struct {
	Number        int
	NodeMap       string
	StartDatetime time.Time
	EndDatetime   time.Time
	TheHours      int
	AmountInHours int
}

type GpuNodeMeteringItem struct {
	Number        int
	NodeMap       string
	StartDatetime time.Time
	EndDatetime   time.Time
	TheHours      int
	AmountInHours int
}

type StorageMeteringItem struct {
	Number        int
	StartDatetime time.Time
	EndDatetime   time.Time
	TheHours      int
	AmountInHours int
}
