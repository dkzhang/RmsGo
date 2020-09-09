package resNodeDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/jmoiron/sqlx"
)

type ResNodePg struct {
	DBInfo
}

func NewResNodePg(sqlxdb *sqlx.DB, tn string) ResNodePg {
	return ResNodePg{
		DBInfo: DBInfo{
			TheDB:     sqlxdb,
			TableName: tn,
		},
	}
}

func (rnpg ResNodePg) Close() {
	rnpg.TheDB.Close()
}

func (rnpg ResNodePg) QueryByID(nodeID int64) (ni resNode.Node, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE node_id=$1`, rnpg.TableName)
	err = rnpg.TheDB.Get(&ni, queryByID, nodeID)
	if err != nil {
		return resNode.Node{},
			fmt.Errorf("QueryByID ResNode in TheDB error: %v", err)
	}
	return ni, nil
}

func (rnpg ResNodePg) QueryAll() (nis []resNode.Node, err error) {
	queryAll := fmt.Sprintf(`SELECT * FROM %s `, rnpg.TableName)
	err = rnpg.TheDB.Select(&nis, queryAll)
	if err != nil {
		return nil,
			fmt.Errorf("QueryAllInfo ResNode from TheDB error: %v", err)
	}
	return nis, nil
}

func (rnpg ResNodePg) Update(ni resNode.Node) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET node_name=:node_name, node_status=:node_status, description=:description, project_id=:project_id, allocated_time=:allocated_time WHERE node_id=:node_id`, rnpg.TableName)

	_, err = rnpg.TheDB.NamedExec(execUpdate, ni)
	if err != nil {
		return fmt.Errorf(" Update ResNode info in DB error: %v", err)
	}
	return nil
}

func (rnpg ResNodePg) UpdateNodes(nodes []resNode.Node) (err error) {
	tx, err := rnpg.TheDB.Begin()
	if err != nil {
		return fmt.Errorf("db transactions begin error: %v", err)
	}

	txStmt, err := tx.Prepare(fmt.Sprintf(
		`UPDATE %s SET node_name=$2, node_status=$3, description=$4, 
			project_id=$5, allocated_time=$6 WHERE node_id=$1`, rnpg.TableName))

	for _, node := range nodes {
		_, err = txStmt.Exec(node.ID,
			node.Name, node.Status, node.Description,
			node.ProjectID, node.AllocatedTime)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf(" Update node (nodeID=%d) in DB error: %v", node.ID, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(" Update nodes in DB error: %v", err)
	}
	return nil
}

func (rnpg ResNodePg) Insert(ni resNode.Node) (err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s (
			node_id, node_name, node_status, description, project_id, allocated_time) 
			VALUES  ($1, $2, $3, $4, $5, $6)`, rnpg.TableName)
	_, err = rnpg.TheDB.Exec(execInsert,
		ni.ID, ni.Name, ni.Status, ni.Description, ni.ProjectID, ni.AllocatedTime)
	if err != nil {
		return fmt.Errorf("TheDB.Exec Insert ResNode info in TheDB error: %v", err)
	}
	return nil
}
