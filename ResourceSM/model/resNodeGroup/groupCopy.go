package resNodeGroup

import "github.com/dkzhang/RmsGo/ResourceSM/model/resNode"

// make new Group, but keep the *Node pointer
func CopyGroup(g *Group) (ng *Group) {
	ng = &Group{
		ID:             g.ID,
		Name:           g.Name,
		Status:         g.Status,
		Description:    g.Description,
		SubGroups:      make([]*Group, 0),
		Nodes:          make([]*resNode.Node, 0),
		NodesNum:       g.NodesNum,
		NodesStatusMap: make(map[int]int),
	}

	// copy Nodes Array
	if g.Nodes != nil {
		for _, node := range g.Nodes {
			ng.Nodes = append(ng.Nodes, node)
		}
	}

	// recursion copy SubGroups Array
	if g.SubGroups != nil {
		for _, sg := range g.SubGroups {
			ng.SubGroups = append(ng.SubGroups, CopyGroup(sg))
		}
	}

	// copy NodesStatusMap
	for k, v := range g.NodesStatusMap {
		ng.NodesStatusMap[k] = v
	}

	return ng
}
