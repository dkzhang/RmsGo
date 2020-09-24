package project

import (
	"fmt"
	"time"
)

type Info struct {
	// Static 8
	ProjectID        int    `db:"project_id" json:"project_id"`
	ProjectName      string `db:"project_name" json:"project_name"`
	ProjectCode      string `db:"project_code" json:"project_code"`
	DepartmentCode   string `db:"department_code" json:"department_code"`
	Department       string `db:"department" json:"department"`
	ChiefID          int    `db:"chief_id" json:"chief_id"`
	ChiefChineseName string `db:"chief_cn_name" json:"chief_cn_name"`
	ExtraInfo        string `db:"extra_info" json:"extra_info"`

	// Status 4
	BasicStatus int `db:"basic_status" json:"basic_status"`

	// Apply Info 6
	StartDate           time.Time `db:"start_date" json:"start_date"`
	TotalDaysApply      int       `db:"total_days_apply" json:"total_days_apply"`
	EndReminderAt       time.Time `db:"end_reminder_at" json:"end_reminder_at"`
	CpuNodesExpected    int       `db:"cpu_nodes_expected" json:"cpu_nodes_expected"`
	GpuNodesExpected    int       `db:"gpu_nodes_expected" json:"gpu_nodes_expected"`
	StorageSizeExpected int       `db:"storage_size_expected" json:"storage_size_expected"`

	// Alloc 3+2
	CpuNodesAcquired    int    `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int    `db:"gpu_nodes_acquired" json:"gpu_nodes_acquired"`
	StorageSizeAcquired int    `db:"storage_size_acquired" json:"storage_size_acquired"`
	AccountAllocInfo    string `db:"account_alloc_info" json:"account_alloc_info"`
	StorageAllocInfo    string `db:"storage_alloc_info" json:"storage_alloc_info"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

var SchemaInfo = `
		CREATE TABLE %s (
    		project_id SERIAL PRIMARY KEY,
			project_name varchar(256) ,
			project_code varchar(32),
			department_code varchar(32),
			department varchar(256),
			chief_id int,
			chief_cn_name varchar(32), 
			extra_info varchar(16384),			
			
			basic_status int,			

			start_date TIMESTAMP WITH TIME ZONE,			
			total_days_apply int,
			end_reminder_at TIMESTAMP WITH TIME ZONE,			
			cpu_nodes_expected int,
			gpu_nodes_expected int,
			storage_size_expected int,

			cpu_nodes_acquired int,
			gpu_nodes_acquired int,
			storage_size_acquired int,		
			account_alloc_info varchar(4096), 
			storage_alloc_info varchar(4096), 

			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE
		);
		`

// 16K = 1024 * 16 = 16384

var TableName = "project_info"
var TableHistoryName = "history_project_info"

func GetSchema() string {
	return fmt.Sprintf(SchemaInfo, TableName)
}

func GetSchemaHistory() string {
	return fmt.Sprintf(SchemaInfo, TableHistoryName)
}

///////////////////////////////////////////////////////////

const (
	BasicStatusApplying   = 1   // 申请中，尚未走完初申流程
	BasicStatusWaiting    = 2   // 已建立，待分配
	BasicStatusRunning    = 4   // 已分配部分或全部资源
	BasicStatusSettlement = 8   // 已归还全部资源，项目结算单审阅中
	BasicStatusDeserted   = 64  // 已放弃的项目（只有从未分配过资源的项目可被放弃，即处于Applying和Waiting状态）
	BasicStatusArchived   = 128 // 已全部结束归档的项目，即将转入历史数据库
)
