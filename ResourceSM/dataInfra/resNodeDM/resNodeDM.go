package resNodeDM

import "github.com/dkzhang/RmsGo/ResourceSM/model/resNode"

type ResNodeDM interface {
	QueryByID(nodeID int64) (resNode.Node, error)
	QueryAll() (nodes []resNode.Node, err error)
	QueryAllMap() (nodesMap map[int64]*resNode.Node)
	Update(node resNode.Node) (err error)
	UpdateNodes(nodes []resNode.Node) (err error)
	Insert(node resNode.Node) (err error)
}
