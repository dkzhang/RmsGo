package projectRes

import (
	"fmt"
	"time"
)

type ResInfo struct {
	// Static 1
	ProjectID int `db:"project_id" json:"project_id"`

	// Alloc 8
	CpuNodesAcquired    int     `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int     `db:"gpu_nodes_acquired" json:"gpu_nodes_acquired"`
	StorageSizeAcquired int     `db:"storage_size_acquired" json:"storage_size_acquired"`
	CpuNodesArray       []int64 `db:"cpu_nodes_array" json:"cpu_nodes_array"`
	CpuNodesStr         string  `db:"cpu_nodes_str" json:"cpu_nodes_str"`
	GpuNodesArray       []int64 `db:"gpu_nodes_array" json:"gpu_nodes_array"`
	GpuNodesStr         string  `db:"gpu_nodes_str" json:"gpu_nodes_str"`
	StorageAllocInfo    string  `db:"storage_alloc_info" json:"storage_alloc_info"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type ResInfoLite struct {
	// Static 1
	ProjectID int `db:"project_id" json:"project_id"`

	// Alloc 3
	CpuNodesAcquired    int `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int `db:"gpu_nodes_acquired" json:"gpu_nodes_acquired"`
	StorageSizeAcquired int `db:"storage_size_acquired" json:"storage_size_acquired"`
}

var SchemaInfo = `
		CREATE TABLE %s (
    		project_id int PRIMARY KEY,			
			
			cpu_nodes_acquired int,
			gpu_nodes_acquired int,
			storage_size_acquired int,	
			cpu_nodes_array integer ARRAY,
			cpu_nodes_str varchar(1024),
			gpu_nodes_array integer ARRAY,
			gpu_nodes_str varchar(1024),
			storage_alloc_info varchar(1024),

			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE
		);
`

// no history table, delete res info after project archived.
var TableName = "project_res_info"

func GetSchema() string {
	return fmt.Sprintf(SchemaInfo, TableName)
}
