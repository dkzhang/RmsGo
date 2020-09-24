package resNodeDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
)

type DirectDB struct {
	nodeDB resNodeDB.ResNodeDB
}

func NewDirectDB(rndb resNodeDB.ResNodeDB) (nmm DirectDB, err error) {
	nmm.nodeDB = rndb
	return nmm, nil
}

func (rnm DirectDB) QueryByID(nodeID int64) (resNode.Node, error) {
	return rnm.nodeDB.QueryByID(nodeID)
}

func (rnm DirectDB) GetAllArray() (nodes []resNode.Node, err error) {
	return rnm.nodeDB.QueryAll()
}

func (rnm DirectDB) GetAllMap() (nodesMap map[int64]*resNode.Node, err error) {
	nodes, err := rnm.nodeDB.QueryAll()
	if err != nil {
		return nil, fmt.Errorf("nodeDB.QueryAll error: %v", err)
	}

	nodesMap = make(map[int64]*resNode.Node)
	for i := range nodes {
		nodesMap[nodes[i].ID] = &nodes[i]
	}
	return nodesMap, nil
}

func (rnm DirectDB) Update(node resNode.Node) (err error) {
	return rnm.nodeDB.Update(node)
}

func (rnm DirectDB) UpdateNodes(nodes []resNode.Node) (err error) {
	return rnm.nodeDB.UpdateNodes(nodes)
}

func (rnm DirectDB) Insert(node resNode.Node) (err error) {
	return rnm.nodeDB.Insert(node)
}
