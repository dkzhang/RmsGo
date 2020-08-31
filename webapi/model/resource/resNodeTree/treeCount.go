package resNodeTree

import "github.com/dkzhang/RmsGo/webapi/model/resource/resNodeTree/groupNode"

func CountTree(t *Tree) {
	groupNode.CountGroup(&t.RootGroup)
}

func CountTreeRO(t *Tree) (nodesNum int, nodesStatusMap map[int]int) {
	return groupNode.CountGroupRO(t.RootGroup)
}
