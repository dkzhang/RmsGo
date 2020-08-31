package resNodeTree

func CopyTree(t *Tree) (nt *Tree) {
	nt = &Tree{
		RootGroup: *(copyGroup(&(t.RootGroup))),
		NodesMap:  make(map[int]*Node),
	}

	for k, v := range t.NodesMap {
		nt.NodesMap[k] = v
	}

	return nt
}
