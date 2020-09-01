package resNodeDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/resource/resNodeTree/groupNode"
	"github.com/jmoiron/sqlx"
)

type DBInfo struct {
	TheDB     *sqlx.DB
	TableName string
}

type ResNodeDB interface {
	QueryByID(nodeID int) (groupNode.Node, error)
	QueryAll() ([]groupNode.Node, error)
	Update(node groupNode.Node) (err error)
	Insert(node groupNode.Node) (err error)
}
