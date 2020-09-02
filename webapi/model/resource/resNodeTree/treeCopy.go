package resNodeTree

import (
	"github.com/dkzhang/RmsGo/webapi/model/resource/resNode"
	"github.com/dkzhang/RmsGo/webapi/model/resource/resNodeGroup"
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
