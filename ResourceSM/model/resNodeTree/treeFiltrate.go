package resNodeTree

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeGroup"
	"github.com/golang-collections/collections/stack"
)

func FiltrateTree(t *Tree, filter func(node resNode.Node) bool) (nt *Tree) {
	nt = CopyTree(t)

	giStack := stack.New()

	giStack.Push(groupAndIndex{
		group: &nt.RootGroup,
		index: 0,
	})

	for {
		if giStack.Len() == 0 {
			break
		}

		gi, _ := giStack.Pop().(groupAndIndex)

		if gi.index == 0 {
			if gi.group.Nodes != nil {
				nodes := make([]*resNode.Node, 0, len(gi.group.Nodes))
				for _, node := range gi.group.Nodes {
					if filter(*node) == true {
						// reserve this node
						nodes = append(nodes, node)
					}
				}
				gi.group.Nodes = nodes
			}
		}

		if gi.group.SubGroups != nil {
			if gi.index < len(gi.group.SubGroups) {
				// push current group
				giStack.Push(groupAndIndex{
					group: gi.group,
					index: gi.index + 1,
				})

				// push a sub group
				giStack.Push(groupAndIndex{
					group: gi.group.SubGroups[gi.index],
					index: 0,
				})
			} else {
				//At the end of SubGroups
				groups := make([]*resNodeGroup.Group, 0, len(gi.group.SubGroups))
				for _, sg := range gi.group.SubGroups {
					if len(sg.Nodes) != 0 || len(sg.SubGroups) != 0 {
						groups = append(groups, sg)
					}
				}
				gi.group.SubGroups = groups
			}
		}
	}
	return nt
}
