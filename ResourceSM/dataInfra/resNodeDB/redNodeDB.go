package resNodeDB

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/jmoiron/sqlx"
)

type DBInfo struct {
	TheDB     *sqlx.DB
	TableName string
}

type ResNodeDB interface {
	QueryByID(nodeID int) (resNode.Node, error)
	QueryAll() ([]resNode.Node, error)
	Update(node resNode.Node) (err error)
	Insert(node resNode.Node) (err error)
	Close()
}
