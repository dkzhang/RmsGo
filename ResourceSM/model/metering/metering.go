package metering

import (
	"fmt"
	"time"
)

type Statement struct {
	MeteringID       int    `db:"metering_id" json:"metering_id"`
	ProjectID        int    `db:"project_id" json:"project_id"`
	MeteringType     int    `db:"metering_type" json:"metering_type"`
	MeteringTypeInfo string `db:"metering_type_info" json:"metering_type_info"`

	CpuAmountInDays      int `db:"cpu_amount_in_days" json:"cpu_amount_in_days"`
	GpuAmountInDays      int `db:"gpu_amount_in_days" json:"gpu_amount_in_days"`
	StorageAmountInDays  int `db:"storage_amount_in_days" json:"storage_amount_in_days"`
	CpuAmountInHours     int `db:"cpu_amount_in_hours" json:"cpu_amount_in_hours"`
	GpuAmountInHours     int `db:"gpu_amount_in_hours" json:"gpu_amount_in_hours"`
	StorageAmountInHours int `db:"storage_amount_in_hours" json:"storage_amount_in_hours"`

	CpuNodeMeteringJson string `db:"cpu_node_metering" json:"cpu_node_metering"`
	GpuNodeMeteringJson string `db:"gpu_node_metering" json:"gpu_node_metering"`
	StorageMeteringJson string `db:"storage_node_metering" json:"storage_metering"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type MeteringItem struct {
	Number        int       `json:"number"`
	AllocInfo     string    `json:"alloc_info"`
	StartDatetime time.Time `json:"start_datetime"`
	EndDatetime   time.Time `json:"end_datetime"`
	TheHours      int       `json:"the_hours"`
	AmountInHours int       `json:"amount_in_hours"`
}

var SchemaInfo = `
		CREATE TABLE %s (
    		metering_id SERIAL PRIMARY KEY,
			project_id int,
			metering_type int,
			metering_type_info varchar(256) ,
			cpu_amount_in_days int,
			gpu_amount_in_days int,
			storage_amount_in_days int,
			cpu_amount_in_hours int,
			gpu_amount_in_hours int,
			storage_amount_in_hours int,
			cpu_node_metering varchar(16384),
			gpu_node_metering varchar(16384),
			storage_node_metering varchar(16384),
			created_at TIMESTAMP WITH TIME ZONE			
		);
		`

// 16K = 1024 * 16 = 16384

var TableName = "metering_statement"
var TableHistoryName = "history_metering_statement"

func GetSchema() string {
	return fmt.Sprintf(SchemaInfo, TableName)
}

func GetSchemaHistory() string {
	return fmt.Sprintf(SchemaInfo, TableHistoryName)
}

const (
	TypeMonthly    = 1
	TypeQuarterly  = 2
	TypeAnnual     = 4
	TypeAnyPeriod  = 8
	TypeSettlement = 16
)
