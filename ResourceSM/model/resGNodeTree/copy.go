package resGNodeTree

import "github.com/dkzhang/RmsGo/ResourceSM/model/resGNode"

func Copy(t *Tree) (nt *Tree) {
	nt = &Tree{
		Root:     *(resGNode.Copy(&t.Root)),
		NodesNum: t.NodesNum,
	}
	return nt
}
