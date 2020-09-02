package resNodeTree

import (
	"github.com/dkzhang/RmsGo/webapi/model/resource/resNodeGroup"
)

func CountTree(t *Tree) {
	resNodeGroup.CountGroup(&t.RootGroup)
}

func CountTreeRO(t *Tree) (nodesNum int, nodesStatusMap map[int]int) {
	return resNodeGroup.CountGroupRO(t.RootGroup)
}
