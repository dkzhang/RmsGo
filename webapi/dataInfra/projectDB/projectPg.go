package projectDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/jmoiron/sqlx"
)

type ProjectPg struct {
	DBInfo
}

func NewProjectPg(sqlxdb *sqlx.DB, stn string, dtn string) ProjectPg {
	return ProjectPg{
		DBInfo: DBInfo{
			TheDB:            sqlxdb,
			StaticTableName:  stn,
			DynamicTableName: dtn,
		},
	}
}

func (ppg ProjectPg) Close() {
	ppg.TheDB.Close()
}

func (ppg ProjectPg) QueryStaticInfoByID(projectID int) (psi project.StaticInfo, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1`, ppg.StaticTableName)
	err = ppg.TheDB.Get(&psi, queryByID, projectID)
	if err != nil {
		return project.StaticInfo{},
			fmt.Errorf("QueryStaticInfoByID project.StaticInfo in TheDB error: %v", err)
	}
	return psi, nil
}
func (ppg ProjectPg) QueryDynamicInfoByID(projectID int) (pdi project.DynamicInfo, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1`, ppg.DynamicTableName)
	err = ppg.TheDB.Get(&pdi, queryByID, projectID)
	if err != nil {
		return project.DynamicInfo{},
			fmt.Errorf("QueryStaticInfoByID project.DynamicInfo in TheDB error: %v", err)
	}
	return pdi, nil
}

func (ppg ProjectPg) QueryInfoByOwner(userID int) (psis []project.StaticInfo, pdis []project.DynamicInfo, err error) {
	queryByOwner := fmt.Sprintf(`SELECT * FROM %s WHERE chief_id=$1`, ppg.StaticTableName)
	err = ppg.TheDB.Select(&psis, queryByOwner, userID)
	if err != nil {
		return nil, nil,
			fmt.Errorf("QueryInfoByOwner project.StaticInfo (userID=%d) from TheDB error: %v", err, userID)
	}

	queryByOwnerJoin := fmt.Sprintf(`SELECT %s.* FROM %s JOIN %s USING(project_id) WHERE %s.chief_id=$1`,
		ppg.DynamicTableName, ppg.StaticTableName, ppg.DynamicTableName, ppg.StaticTableName)
	err = ppg.TheDB.Select(&pdis, queryByOwnerJoin, userID)
	if err != nil {
		return psis, nil,
			fmt.Errorf("QueryInfoByOwner project.DynamicInfo (userID=%d) from TheDB error: %v", err, userID)
	}

	return psis, pdis, nil
}

func (ppg ProjectPg) QueryInfoByDepartmentCode(dc string) (psis []project.StaticInfo, pdis []project.DynamicInfo, err error) {
	queryByDC := fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1`, ppg.StaticTableName)
	err = ppg.TheDB.Select(&psis, queryByDC, dc)
	if err != nil {
		return nil, nil,
			fmt.Errorf("QueryInfoByDepartmentCode project.StaticInfo (DepartmentCode=%s) from TheDB error: %v", err, dc)
	}

	queryByDCJoin := fmt.Sprintf(`SELECT %s.* FROM %s JOIN %s USING(project_id) WHERE %s.department_code=$1`,
		ppg.DynamicTableName, ppg.StaticTableName, ppg.DynamicTableName, ppg.StaticTableName)
	err = ppg.TheDB.Select(&pdis, queryByDCJoin, dc)
	if err != nil {
		return psis, nil,
			fmt.Errorf("QueryInfoByDepartmentCode project.DynamicInfo (DepartmentCode=%s) from TheDB error: %v", err, dc)
	}

	return psis, pdis, nil
}
func (ppg ProjectPg) QueryAllInfo() (psis []project.StaticInfo, pdis []project.DynamicInfo, err error) {
	queryAll := fmt.Sprintf(`SELECT * FROM %s `, ppg.StaticTableName)
	err = ppg.TheDB.Select(&psis, queryAll)
	if err != nil {
		return nil, nil,
			fmt.Errorf("QueryInfoByOwner project.StaticInfo ALL from TheDB error: %v", err)
	}

	queryAllJoin := fmt.Sprintf(`SELECT %s.* FROM %s JOIN %s USING(project_id)`,
		ppg.DynamicTableName, ppg.StaticTableName, ppg.DynamicTableName)
	err = ppg.TheDB.Select(&pdis, queryAllJoin)
	if err != nil {
		return psis, nil,
			fmt.Errorf("QueryInfoByOwner project.DynamicInfo ALL from TheDB error: %v", err)
	}

	return psis, pdis, nil
}

///////////////////////////////////////////////////////////////////////////////
func (ppg ProjectPg) InsertAllInfo(psi project.StaticInfo, pdi project.DynamicInfo) (projectID int, err error) {
	execInsertStatic := fmt.Sprintf(`INSERT INTO %s (project_name, project_code, department_code, department, chief_id, chief_cn_name, extra_info, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING project_id`, ppg.StaticTableName)
	err = ppg.TheDB.Get(&projectID, execInsertStatic,
		psi.ProjectName, psi.ProjectCode,
		psi.DepartmentCode, psi.Department,
		psi.ChiefID, psi.ChiefChineseName, psi.ExtraInfo,
		psi.CreatedAt, psi.UpdatedAt)
	if err != nil {
		return -1, fmt.Errorf("TheDB.Get Insert Project StaticInfo in TheDB error: %v", err)
	}

	execInsertDynamic := fmt.Sprintf(`INSERT INTO %s (project_id, basic_status, computing_alloc_status, storage_alloc_status, start_date, total_days_apply, end_reminder_at, cpu_nodes_expected, gpu_nodes_expected, storage_size_expected, cpu_nodes_acquired, gpu_nodes_acquired, storage_size_acquired, created_at, updated_at)  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING project_id`, ppg.DynamicTableName)
	err = ppg.TheDB.Get(&projectID, execInsertDynamic,
		projectID, pdi.BasicStatus,
		pdi.ComputingAllocStatus, pdi.StorageAllocStatus,
		pdi.StartDate, pdi.TotalDaysApply, pdi.EndReminderAt,
		pdi.CpuNodesExpected, pdi.GpuNodesExpected, pdi.StorageSizeExpected,
		pdi.CpuNodesAcquired, pdi.GpuNodesAcquired, pdi.StorageSizeAcquired,
		pdi.CreatedAt, pdi.UpdatedAt)
	if err != nil {
		return -1, fmt.Errorf("TheDB.Get Insert Project DynamicInfo in TheDB error: %v", err)
	}
	return projectID, nil
}
func (ppg ProjectPg) UpdateStaticInfo(psi project.StaticInfo) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET project_name=:project_name, project_code=:project_code, extra_info=:extra_info, updated_at=:updated_at WHERE project_id=:project_id`, ppg.StaticTableName)

	_, err = ppg.TheDB.NamedExec(execUpdate, psi)
	if err != nil {
		return fmt.Errorf("TheDB.NamedExec UPDATE project.StaticInfo error: %v", err)
	}
	return nil
}
func (ppg ProjectPg) UpdateDynamicInfo(pdi project.DynamicInfo) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET basic_status=:basic_status, computing_alloc_status=:computing_alloc_status, storage_alloc_status=:storage_alloc_status, start_date=:start_date, total_days_apply=:total_days_apply, end_reminder_at=:end_reminder_at, cpu_nodes_expected=:cpu_nodes_expected, gpu_nodes_expected=:gpu_nodes_expected, storage_size_expected=:storage_size_expected, cpu_nodes_acquired=:cpu_nodes_acquired, gpu_nodes_acquired=:gpu_nodes_acquired, storage_size_acquired=:storage_size_acquired, updated_at=:updated_at WHERE project_id=:project_id`, ppg.DynamicTableName)

	_, err = ppg.TheDB.NamedExec(execUpdate, pdi)
	if err != nil {
		return fmt.Errorf("TheDB.NamedExec UPDATE project.DynamicInfo error: %v", err)
	}
	return nil
}

func (ppg ProjectPg) InnerArchiveProject(stnHistory string, dtnHistory string, projectID int) (err error) {
	execCopyS := fmt.Sprintf(`INSERT INTO %s SELECT * FROM %s WHERE project_id=$1`, stnHistory, ppg.StaticTableName)
	execCopyD := fmt.Sprintf(`INSERT INTO %s SELECT * FROM %s WHERE project_id=$1`, dtnHistory, ppg.DynamicTableName)
	execDelS := fmt.Sprintf(`DELETE FROM %s WHERE project_id=$1`, ppg.StaticTableName)
	execDelD := fmt.Sprintf(`DELETE FROM %s WHERE project_id=$1`, ppg.DynamicTableName)

	tx, err := ppg.TheDB.Begin()
	_, err = ppg.TheDB.Exec(execCopyS, projectID)
	_, err = ppg.TheDB.Exec(execCopyD, projectID)
	_, err = ppg.TheDB.Exec(execDelS, projectID)
	_, err = ppg.TheDB.Exec(execDelD, projectID)
	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("InnerArchiveProject Commit error: %v", err)
	}
	return nil
}
