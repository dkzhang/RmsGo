package resNodeDM

import "github.com/dkzhang/RmsGo/ResourceSM/model/resNode"

type ResNodeDM interface {
	QueryByID(nodeID int) (resNode.Node, error)
	QueryAll() ([]resNode.Node, error)
	Update(node resNode.Node) (err error)
	Insert(node resNode.Node) (err error)
}
