package resNodeTree

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeGroup"
)

func CopyTree(t *Tree) (nt *Tree) {
	nt = &Tree{
		RootGroup: *(resNodeGroup.CopyGroup(&(t.RootGroup))),
		NodesMap:  make(map[int]*resNode.Node),
	}

	for k, v := range t.NodesMap {
		nt.NodesMap[k] = v
	}

	return nt
}
