package resGNodeTree

import "github.com/dkzhang/RmsGo/ResourceSM/model/resGNode"

func Count(t *Tree) {
	t.NodesNum = resGNode.CountRO(&t.Root)
}

func CountRO(t *Tree) (nodesNum int) {
	return resGNode.CountRO(&t.Root)
}
