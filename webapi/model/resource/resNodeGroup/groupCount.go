package resNodeGroup

import "github.com/dkzhang/RmsGo/webapi/model/resource/resNode"

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

	if g.NodesStatusMap[resNode.StatusSelected]+g.NodesStatusMap[resNode.StatusUnselected] != 0 {
		if g.NodesStatusMap[resNode.StatusUnselected] == 0 {
			g.Status = resNode.StatusSelected
		} else if g.NodesStatusMap[resNode.StatusSelected] == 0 {
			g.Status = resNode.StatusUnselected
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
