package metering

import (
	"time"
)

type MeteringList struct {
	MeteringID          int
	CpuAmountInDays     int
	GpuAmountInDays     int
	StorageAmountInDays int

	CpuAmountInHours     int
	GpuAmountInHours     int
	StorageAmountInHours int

	CpuNodeMeteringList []CpuNodeMeteringItem
	GpuNodeMeteringList []GpuNodeMeteringItem
	StorageMeteringList []StorageMeteringItem

	CreatedAt time.Time
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
