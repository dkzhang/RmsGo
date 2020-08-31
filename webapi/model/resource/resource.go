package resource

type Resource struct {
	CpuNodes    int `json:"cpu_nodes"`
	GpuNodes    int `json:"gpu_nodes"`
	StorageSize int `json:"storage_size"`
}
