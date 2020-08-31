package groupNode

func CountGroup(g *Group) {
	if g == nil {
		return
	}

	g.NodesStatusMap = make(map[int]int)
	g.NodesNum = 0

	if g.Nodes != nil {
		for _, node := range g.Nodes {
			g.NodesNum++
			g.NodesStatusMap[node.Status]++
		}
	}

	if g.SubGroups != nil {
		for _, sg := range g.SubGroups {
			CountGroup(sg)
			g.NodesNum += sg.NodesNum
			for k, v := range sg.NodesStatusMap {
				g.NodesStatusMap[k] += v
			}
		}
	}

	if g.NodesStatusMap[StatusSelected]+g.NodesStatusMap[StatusUnselected] != 0 {
		if g.NodesStatusMap[StatusUnselected] == 0 {
			g.Status = StatusSelected
		} else if g.NodesStatusMap[StatusSelected] == 0 {
			g.Status = StatusUnselected
		} else {
			g.Status = StatusPartiallySelected
		}
	}
	return
}

func CountGroupRO(g Group) (nodesNum int, nodesStatusMap map[int]int) {
	nodesStatusMap = make(map[int]int)
	nodesNum = 0

	if g.Nodes != nil {
		for _, node := range g.Nodes {
			nodesNum++
			nodesStatusMap[node.Status]++
		}
	}

	if g.SubGroups != nil {
		for _, sg := range g.SubGroups {
			num, sMap := CountGroupRO(*sg)
			nodesNum += num
			for k, v := range sMap {
				nodesStatusMap[k] += v
			}
		}
	}
	return nodesNum, nodesStatusMap
}
