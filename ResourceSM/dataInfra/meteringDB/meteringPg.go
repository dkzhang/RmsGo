package meteringDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"github.com/jmoiron/sqlx"
)

type MeteringPg struct {
	DBInfo
}

func NewMeteringPg(sqlxdb *sqlx.DB, tn string) MeteringPg {
	return MeteringPg{
		DBInfo: DBInfo{
			TheDB:     sqlxdb,
			TableName: tn,
		},
	}
}

func (mpg MeteringPg) Close() {
	mpg.TheDB.Close()
}

func (mpg MeteringPg) Query(projectID int, mType int, typeInfo string) (ms metering.Statement, err error) {
	queryOne := fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1 AND metering_type=$2 AND metering_type_info=$3`, mpg.TableName)
	err = mpg.TheDB.Get(&ms, queryOne, projectID, mType, typeInfo)
	if err != nil {
		return metering.Statement{},
			fmt.Errorf("query One metering.Statement in TheDB error: %v", err)
	}
	return ms, nil
}

func (mpg MeteringPg) QueryAll(projectID int, mType int) (mss []metering.Statement, err error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1 AND metering_type & $2 != 0`, mpg.TableName)
	err = mpg.TheDB.Select(&mss, query, projectID, mType)
	if err != nil {
		return nil, fmt.Errorf("query all metering statement(projectID=%d AND type=%d) from TheDB error: %v", projectID, mType, err)
	}
	return mss, nil
}

func (mpg MeteringPg) Insert(ms metering.Statement) (mid int, err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s 
			(project_id, metering_type, metering_type_info,
			cpu_amount_in_days, gpu_amount_in_days, storage_amount_in_days, cpu_amount_in_hours, gpu_amount_in_hours, storage_amount_in_hours,
			cpu_node_metering, gpu_node_metering, storage_node_metering,
			created_at) 
			VALUES ($1, $2, $3,
			$4, $5, $6, $7, $8, $9, 
			$10, $11, $12,
			$13) 
			RETURNING metering_id`, mpg.TableName)
	err = mpg.TheDB.Get(&mid, execInsert,
		ms.ProjectID, ms.MeteringType, ms.MeteringTypeInfo,
		ms.CpuAmountInDays, ms.GpuAmountInDays, ms.StorageAmountInDays, ms.CpuAmountInHours, ms.GpuAmountInHours, ms.StorageAmountInHours,
		ms.CpuNodeMeteringJson, ms.GpuNodeMeteringJson, ms.StorageMeteringJson,
		ms.CreatedAt)
	if err != nil {
		return -1, fmt.Errorf("TheDB.Get Insert metering.Statement in TheDB error: %v", err)
	}
	return mid, nil
}
