package resGNode

const GroupBase = 1e5

func CountRO(gn *ResGNode) (nodesNum int) {
	nodesNum = 0

	if gn == nil {
		return 0
	}

	if gn.ID < GroupBase {
		return 1
	}

	if gn.Children != nil {
		for _, child := range gn.Children {
			nodesNum += CountRO(child)
		}
	}
	return nodesNum
}
