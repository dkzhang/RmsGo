package project

import "time"

type AllocInfoEx struct {
	ProjectID int `db:"project_id" json:"project_id"`
	AllocInfo
}

type AllocInfo struct {
	CpuNodesAcquired    int `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int `db:"gpu_nodes_acquired" json:"gpu_nodes_acquired"`
	StorageSizeAcquired int `db:"storage_size_acquired" json:"storage_size_acquired"`

	CpuNodesMap string `db:"cpu_nodes_map" json:"cpu_nodes_map"`
	GpuNodesMap string `db:"gpu_nodes_map" json:"gpu_nodes_map"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
