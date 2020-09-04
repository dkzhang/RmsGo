package resAllocDB

import (
	"database/sql"
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type ResAllocPg struct {
	DBInfo
}

func NewResAllocPg(sqlxdb *sqlx.DB, tn string) ResAllocPg {
	return ResAllocPg{
		DBInfo: DBInfo{
			TheDB:     sqlxdb,
			TableName: tn,
		},
	}
}

func (rnpg ResAllocPg) Close() {
	rnpg.TheDB.Close()
}

func (rnpg ResAllocPg) QueryByID(recordID int) (ra resAlloc.Record, err error) {
	stmt, err := rnpg.TheDB.Prepare(fmt.Sprintf(`SELECT * FROM %s WHERE record_id=$1`, rnpg.TableName))
	if err != nil {
		logrus.Fatalf("QueryByID TheDB.Prepare statement error: %v", err)
	}

	defer stmt.Close()

	ra, err = rowScan(stmt.QueryRow(recordID))
	if err != nil {
		return resAlloc.Record{},
			fmt.Errorf("query ResourceAllocateRecord (ID = %d) in TheDB error: %v", recordID, err)
	}
	return ra, nil
}

func (rnpg ResAllocPg) QueryByProjectID(projectID int) (rs []resAlloc.Record, err error) {
	stmt, err := rnpg.TheDB.Prepare(fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1`, rnpg.TableName))
	if err != nil {
		logrus.Fatalf("QueryByProjectID TheDB.Prepare statement error: %v", err)
	}

	defer stmt.Close()

	//////
	rows, err := stmt.Query(projectID)
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
				fmt.Errorf("query ResourceAllocateRecord (projectID = %d) in TheDB error: %v", projectID, err)
		}
		rs = append(rs, ra)
	}
	return rs, nil
}

func (rnpg ResAllocPg) QueryAll() (rs []resAlloc.Record, err error) {

	stmt, err := rnpg.TheDB.Prepare(fmt.Sprintf(`SELECT * FROM %s`, rnpg.TableName))
	if err != nil {
		logrus.Fatalf("QueryByProjectID TheDB.Prepare statement error: %v", err)
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
				fmt.Errorf("query all ResourceAllocateRecord from TheDB error: %v", err)
		}
		rs = append(rs, ra)
	}
	return rs, nil
}

func (rnpg ResAllocPg) Insert(rar resAlloc.Record) (err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s (
			project_id, 
			num_before, alloc_info_before, alloc_info_before_str, 
			num_after, alloc_info_after, alloc_info_after_str,
			num_change, alloc_info_change, alloc_info_change_str,
			ctrl_id, ctrl_cn_name, created_at) 
			VALUES  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
			$11, $12, $13)`, rnpg.TableName)
	_, err = rnpg.TheDB.Exec(execInsert,
		rar.ProjectID,
		rar.NumBefore, pq.Array(rar.AllocInfoBefore), rar.AllocInfoBeforeStr,
		rar.NumAfter, pq.Array(rar.AllocInfoAfter), rar.AllocInfoAfterStr,
		rar.NumChange, pq.Array(rar.AllocInfoChange), rar.AllocInfoChangeStr,
		rar.CtrlID, rar.CtrlChineseName, rar.CreatedAt)
	if err != nil {
		return fmt.Errorf("TheDB.Exec Insert ResourceAllocateRecord in TheDB error: %v", err)
	}
	return nil
}

/////////////////////////////////////////////////////////////////////
func rowScan(r *sql.Row) (ra resAlloc.Record, err error) {
	err = r.Scan(&ra.RecordID, &ra.ProjectID,
		&ra.NumBefore, pq.Array(&(ra.AllocInfoBefore)), &ra.AllocInfoBeforeStr,
		&ra.NumAfter, pq.Array(&(ra.AllocInfoAfter)), &ra.AllocInfoAfterStr,
		&ra.NumChange, pq.Array(&(ra.AllocInfoChange)), &ra.AllocInfoAfterStr,
		&ra.CtrlID, &ra.CtrlChineseName, &ra.CtrlChineseName)
	if err != nil {
		return resAlloc.Record{},
			fmt.Errorf("rowScan resAlloc.Record error: %v", err)
	}
	return ra, nil
}
func rowsScan(rs *sql.Rows) (ra resAlloc.Record, err error) {
	err = rs.Scan(&ra.RecordID, &ra.ProjectID,
		&ra.NumBefore, pq.Array(&(ra.AllocInfoBefore)), &ra.AllocInfoBeforeStr,
		&ra.NumAfter, pq.Array(&(ra.AllocInfoAfter)), &ra.AllocInfoAfterStr,
		&ra.NumChange, pq.Array(&(ra.AllocInfoChange)), &ra.AllocInfoAfterStr,
		&ra.CtrlID, &ra.CtrlChineseName, &ra.CtrlChineseName)
	if err != nil {
		return resAlloc.Record{},
			fmt.Errorf("rowsScan resAlloc.Record error: %v", err)
	}
	return ra, nil
}
