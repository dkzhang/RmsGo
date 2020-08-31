package resNodeTree

import "github.com/dkzhang/RmsGo/webapi/model/resource/resNodeTree/groupNode"

func CopyTree(t *Tree) (nt *Tree) {
	nt = &Tree{
		RootGroup: *(groupNode.CopyGroup(&(t.RootGroup))),
		NodesMap:  make(map[int]*groupNode.Node),
	}

	for k, v := range t.NodesMap {
		nt.NodesMap[k] = v
	}

	return nt
}
