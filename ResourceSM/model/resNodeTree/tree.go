package resNodeTree

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeGroup"
)

type Tree struct {
	RootGroup resNodeGroup.Group
	NodesMap  map[int64]*resNode.Node

	NodesNum       int
	NodesStatusMap map[int]int
}
