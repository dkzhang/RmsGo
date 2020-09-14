package resGNode

func Copy(gn *ResGNode) (ngn *ResGNode) {
	ngn = &ResGNode{
		ID:       gn.ID,
		Label:    gn.Label,
		Children: make([]*ResGNode, 0),
	}

	if gn.Children != nil {
		for _, child := range gn.Children {
			ngn.Children = append(ngn.Children, Copy(child))
		}
	}
	return ngn
}
