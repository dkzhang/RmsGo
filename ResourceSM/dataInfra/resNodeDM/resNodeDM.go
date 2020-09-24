package resNodeDM

import "github.com/dkzhang/RmsGo/ResourceSM/model/resNode"

type ResNodeDM interface {
	QueryByID(nodeID int64) (resNode.Node, error)
	GetAllArray() (nodes []resNode.Node, err error)
	GetAllMap() (nodesMap map[int64]*resNode.Node, err error)
	Update(node resNode.Node) (err error)
	UpdateNodes(nodes []resNode.Node) (err error)
	Insert(node resNode.Node) (err error)
}
