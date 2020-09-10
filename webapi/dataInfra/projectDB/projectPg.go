package projectDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/jmoiron/sqlx"
)

type ProjectPg struct {
	DBInfo
}

func NewProjectPg(sqlxdb *sqlx.DB, tn string) ProjectPg {
	return ProjectPg{
		DBInfo: DBInfo{
			TheDB:     sqlxdb,
			TableName: tn,
		},
	}
}

func (ppg ProjectPg) Close() {
	ppg.TheDB.Close()
}

func (ppg ProjectPg) QueryByID(projectID int) (pi project.Info, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1`, ppg.TableName)
	err = ppg.TheDB.Get(&pi, queryByID, projectID)
	if err != nil {
		return project.Info{},
			fmt.Errorf("QueryByID project in TheDB error: %v", err)
	}
	return pi, nil
}
func (ppg ProjectPg) QueryByOwner(userID int) (pis []project.Info, err error) {
	queryByOwner := fmt.Sprintf(`SELECT * FROM %s WHERE chief_id=$1`, ppg.TableName)
	err = ppg.TheDB.Select(&pis, queryByOwner, userID)
	if err != nil {
		return nil,
			fmt.Errorf("QueryByOwner project (userID = %d) in TheDB error: %v", userID, err)
	}
	return pis, nil
}
func (ppg ProjectPg) QueryByDepartmentCode(dc string) (pis []project.Info, err error) {
	queryByDC := fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1`, ppg.TableName)
	err = ppg.TheDB.Select(&pis, queryByDC, dc)
	if err != nil {
		return nil,
			fmt.Errorf("QueryByDepartmentCode project (DepartmentCode=%s) from TheDB error: %v", dc, err)
	}
	return pis, nil
}
func (ppg ProjectPg) QueryAllInfo() (pis []project.Info, err error) {
	queryAll := fmt.Sprintf(`SELECT * FROM %s `, ppg.TableName)
	err = ppg.TheDB.Select(&pis, queryAll)
	if err != nil {
		return nil,
			fmt.Errorf("QueryAllInfo project from TheDB error: %v", err)
	}
	return pis, nil
}

/////////////////////////////////////////////////

func (ppg ProjectPg) Insert(pi project.Info) (projectID int, err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s (
			project_name, project_code, department_code, department, chief_id, chief_cn_name, extra_info, 
			basic_status, 
			start_date, total_days_apply, end_reminder_at, cpu_nodes_expected, gpu_nodes_expected, storage_size_expected, 
			cpu_nodes_acquired, gpu_nodes_acquired, storage_size_acquired, 
			created_at, updated_at) 
			VALUES  ($1, $2, $3, $4, $5, $6, $7, 
					$8, 
					$9, $10, $11, $12, $13, $14, 
					$15, $16, $17, 
					$18, $19) RETURNING project_id`, ppg.TableName)
	err = ppg.TheDB.Get(&projectID, execInsert,
		pi.ProjectName, pi.ProjectCode, pi.DepartmentCode, pi.Department, pi.ChiefID, pi.ChiefChineseName, pi.ExtraInfo,
		pi.BasicStatus,
		pi.StartDate, pi.TotalDaysApply, pi.EndReminderAt, pi.CpuNodesExpected, pi.GpuNodesExpected, pi.StorageSizeExpected,
		pi.CpuNodesAcquired, pi.GpuNodesAcquired, pi.StorageSizeAcquired,
		pi.CreatedAt, pi.UpdatedAt)
	if err != nil {
		return -1,
			fmt.Errorf("TheDB.Get Insert Project in TheDB error: %v", err)
	}
	return projectID, nil
}

func (ppg ProjectPg) UpdateBasicInfo(bi project.BasicInfo) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET project_name=:project_name, extra_info=:extra_info, updated_at=:updated_at WHERE project_id=:project_id`, ppg.TableName)

	_, err = ppg.TheDB.NamedExec(execUpdate, bi)
	if err != nil {
		return fmt.Errorf("UpdateBasicInfo error: %v", err)
	}
	return nil
}

func (ppg ProjectPg) UpdateCodeInfo(pc project.CodeInfo) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET project_code=:project_code, updated_at=:updated_at WHERE project_id=:project_id`, ppg.TableName)

	_, err = ppg.TheDB.NamedExec(execUpdate, pc)
	if err != nil {
		return fmt.Errorf("UpdateCodeInfo error: %v", err)
	}
	return nil
}

func (ppg ProjectPg) UpdateStatusInfo(si project.StatusInfo) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET basic_status=:basic_status, updated_at=:updated_at WHERE project_id=:project_id`, ppg.TableName)

	_, err = ppg.TheDB.NamedExec(execUpdate, si)
	if err != nil {
		return fmt.Errorf("UpdateCodeInfo error: %v", err)
	}
	return nil
}

func (ppg ProjectPg) UpdateApplyInfo(ai project.ApplyInfo) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET start_date=:start_date, total_days_apply=:total_days_apply, end_reminder_at=:end_reminder_at, 
							cpu_nodes_expected=:cpu_nodes_expected, gpu_nodes_expected=:gpu_nodes_expected, storage_size_expected=:storage_size_expected,
							updated_at=:updated_at WHERE project_id=:project_id`, ppg.TableName)

	_, err = ppg.TheDB.NamedExec(execUpdate, ai)
	if err != nil {
		return fmt.Errorf("UpdateApplyInfo error: %v", err)
	}
	return nil
}

func (ppg ProjectPg) UpdateAllocInfo(ali project.AllocInfo) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET cpu_nodes_acquired=:cpu_nodes_acquired, gpu_nodes_acquired=:gpu_nodes_acquired, storage_size_acquired=:storage_size_acquired, 
							updated_at=:updated_at WHERE project_id=:project_id`, ppg.TableName)

	_, err = ppg.TheDB.NamedExec(execUpdate, ali)
	if err != nil {
		return fmt.Errorf("UpdateAllocInfo error: %v", err)
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////

func (ppg ProjectPg) InnerArchiveProject(historyTableName string, projectID int) (err error) {
	execCopy := fmt.Sprintf(`INSERT INTO %s SELECT * FROM %s WHERE project_id=$1`, historyTableName, ppg.TableName)
	execDel := fmt.Sprintf(`DELETE FROM %s WHERE project_id=$1`, ppg.TableName)

	tx, err := ppg.TheDB.Begin()
	_, err = ppg.TheDB.Exec(execCopy, projectID)
	_, err = ppg.TheDB.Exec(execDel, projectID)
	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("InnerArchiveProject Commit error: %v", err)
	}
	return nil
}
