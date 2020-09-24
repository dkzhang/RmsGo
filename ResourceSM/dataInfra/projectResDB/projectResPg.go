package projectResDB

import (
	"database/sql"
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type ProjectResPg struct {
	DBInfo
}

func NewProjectResPg(sqlxdb *sqlx.DB, tn string) ProjectResPg {
	return ProjectResPg{
		DBInfo: DBInfo{
			TheDB:     sqlxdb,
			TableName: tn,
		},
	}
}

func (prpg ProjectResPg) Close() {
	prpg.TheDB.Close()
}

func (prpg ProjectResPg) QueryByID(projectID int) (pr projectRes.ResInfo, err error) {
	stmt, err := prpg.TheDB.Prepare(fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1`, prpg.TableName))
	if err != nil {
		logrus.Fatalf("ProjectResPg QueryByID TheDB.Prepare statement error: %v", err)
	}

	defer stmt.Close()

	pr, err = rowScan(stmt.QueryRow(projectID))
	if err != nil {
		return projectRes.ResInfo{},
			fmt.Errorf("query Project Resource Info (ID = %d) in TheDB error: %v", projectID, err)
	}
	return pr, nil
}

func (prpg ProjectResPg) QueryAll() (prs []projectRes.ResInfo, err error) {

	stmt, err := prpg.TheDB.Prepare(fmt.Sprintf(`SELECT * FROM %s`, prpg.TableName))
	if err != nil {
		logrus.Fatalf("ProjectResPg GetAllArray TheDB.Prepare statement error: %v", err)
	}

	defer stmt.Close()

	//////
	rows, err := stmt.Query()
	if err != nil {
		return nil,
			fmt.Errorf("QueryByProjectID stmt.Query(projectID) error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan()
		ra, err := rowsScan(rows)
		if err != nil {
			return nil,
				fmt.Errorf("query all Project Resource Info from TheDB error: %v", err)
		}
		prs = append(prs, ra)
	}
	return prs, nil
}

func (prpg ProjectResPg) Insert(pr projectRes.ResInfo) (err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s (
			project_id, 
			cpu_nodes_acquired, gpu_nodes_acquired, storage_size_acquired, 
			cpu_nodes_array, cpu_nodes_str, 
			gpu_nodes_array, gpu_nodes_str, 
			storage_alloc_info,
			created_at, updated_at) 
			VALUES  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, prpg.TableName)

	stmt, err := prpg.TheDB.Prepare(execInsert)
	if err != nil {
		logrus.Fatalf("ProjectResPg Insert TheDB.Prepare statement error: %v", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		pr.ProjectID,
		pr.CpuNodesAcquired, pr.GpuNodesAcquired, pr.StorageSizeAcquired,
		pq.Array(pr.CpuNodesArray), pr.CpuNodesStr,
		pq.Array(pr.GpuNodesArray), pr.GpuNodesStr,
		pr.StorageAllocInfo,
		pr.CreatedAt, pr.UpdatedAt)
	if err != nil {
		return fmt.Errorf("TheDB.Exec Insert Project Resource Info in TheDB error: %v", err)
	}
	return nil
}

func (prpg ProjectResPg) Update(pr projectRes.ResInfo) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET			
			cpu_nodes_acquired = $2, gpu_nodes_acquired = $3, storage_size_acquired = $4, 
			cpu_nodes_array = $5, cpu_nodes_str = $6, 
			gpu_nodes_array = $7, gpu_nodes_str = $8, 
			storage_alloc_info = $9,
			updated_at = $10
			WHERE project_id=$1`, prpg.TableName)

	stmt, err := prpg.TheDB.Prepare(execUpdate)
	if err != nil {
		logrus.Fatalf("ProjectResPg Update TheDB.Prepare statement error: %v", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		pr.ProjectID,
		pr.CpuNodesAcquired, pr.GpuNodesAcquired, pr.StorageSizeAcquired,
		pq.Array(pr.CpuNodesArray), pr.CpuNodesStr,
		pq.Array(pr.GpuNodesArray), pr.GpuNodesStr,
		pr.StorageAllocInfo,
		pr.UpdatedAt)
	if err != nil {
		return fmt.Errorf("TheDB.Exec Update Project Resource Info in TheDB error: %v", err)
	}
	return nil
}

func (prpg ProjectResPg) Delete(projectID int) (err error) {
	deleteProjectRes := fmt.Sprintf(`DELETE FROM %s WHERE project_id=$1`, prpg.TableName)

	result, err := prpg.TheDB.Exec(deleteProjectRes, projectID)
	if err != nil {
		return fmt.Errorf("db.Exec(deleteProjectRes, projectID), projectID = %d", projectID)
	}
	fmt.Printf("Delete ProjectRes success: %v \n", result)
	return nil
}

/////////////////////////////////////////////////////////////////////
func rowScan(r *sql.Row) (pr projectRes.ResInfo, err error) {
	err = r.Scan(&pr.ProjectID,
		&pr.CpuNodesAcquired, &pr.GpuNodesAcquired, &pr.StorageSizeAcquired,
		pq.Array(&(pr.CpuNodesArray)), &pr.CpuNodesStr,
		pq.Array(&(pr.GpuNodesArray)), &pr.GpuNodesStr,
		&pr.StorageAllocInfo,
		&pr.CreatedAt, &pr.UpdatedAt)
	if err != nil {
		return projectRes.ResInfo{},
			fmt.Errorf("rowScan projectRes.ResInfo error: %v", err)
	}
	return pr, nil
}
func rowsScan(rs *sql.Rows) (pr projectRes.ResInfo, err error) {
	err = rs.Scan(&pr.ProjectID,
		&pr.CpuNodesAcquired, &pr.GpuNodesAcquired, &pr.StorageSizeAcquired,
		pq.Array(&(pr.CpuNodesArray)), &pr.CpuNodesStr,
		pq.Array(&(pr.GpuNodesArray)), &pr.GpuNodesStr,
		&pr.StorageAllocInfo,
		&pr.CreatedAt, &pr.UpdatedAt)
	if err != nil {
		return projectRes.ResInfo{},
			fmt.Errorf("rowScan projectRes.ResInfo error: %v", err)
	}
	return pr, nil
}
