package resNodeTree

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeGroup"
)

func CountTree(t *Tree) {
	resNodeGroup.CountGroup(&t.RootGroup)
	t.NodesNum = t.RootGroup.NodesNum
	t.NodesStatusMap = t.RootGroup.NodesStatusMap
}

func CountTreeRO(t *Tree) (nodesNum int, nodesStatusMap map[int]int) {
	return resNodeGroup.CountGroupRO(t.RootGroup)
}
