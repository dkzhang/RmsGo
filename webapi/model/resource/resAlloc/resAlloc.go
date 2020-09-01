package resAlloc

import "time"

type Record struct {
	RecordID  int `db:"record_id" json:"record_id"`
	ProjectID int `db:"project_id" json:"project_id"`

	NumBefore          int    `db:"num_before" json:"num_before"`
	AllocInfoBefore    []int  `db:"alloc_info_before" json:"alloc_info_before"`
	AllocInfoBeforeStr string `db:"alloc_info_before_str" json:"alloc_info_before_str"`
	NumAfter           int    `db:"num_after" json:"num_after"`
	AllocInfoAfter     []int  `db:"alloc_info_after" json:"alloc_info_after"`
	AllocInfoAfterStr  string `db:"alloc_info_after_str" json:"alloc_info_after_str"`
	NumChange          int    `db:"num_change" json:"num_change"`
	AllocInfoChange    []int  `db:"alloc_info_change" json:"alloc_info_change"`
	AllocInfoChangeStr string `db:"alloc_info_change_str" json:"alloc_info_change_str"`

	CtrlID          int    `db:"ctrl_id" json:"ctrl_id"`
	CtrlChineseName string `db:"ctrl_cn_name" json:"ctrl_cn_name"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

var SchemaRecord = `
		CREATE TABLE %s (
			record_id SERIAL PRIMARY KEY,
    		project_id int,			
			num_before int,
			alloc_info_before integer ARRAY,
			alloc_info_before_str varchar(1024),
			num_after int,
			alloc_info_after integer ARRAY,
			alloc_info_after_str varchar(1024),			
			num_change int,
			alloc_info_change integer ARRAY,
			alloc_info_change_str varchar(1024),			
			ctrl_id	 int,
			ctrl_cn_name varchar(32),
			created_at TIMESTAMP WITH TIME ZONE
		);
		`

var TableNameCPU = "res_alloc_cpu"
var TableNameGPU = "res_alloc_gpu"
var TableNameStorage = "res_alloc_storage"
var TableHistoryNameCPU = "history_res_alloc_cpu"
var TableHistoryNameGPU = "history_res_alloc_gpu"
var TableHistoryNameStorage = "history_res_alloc_storage"
