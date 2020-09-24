package resGNodeTree

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/golang-collections/collections/stack"
)

// generate a NodesMap from the Tree By iteration
type gNodeAndIndex struct {
	gNode *resGNode.ResGNode
	index int
}

func Filtrate(t *Tree, nodesMap map[int64]*resNode.Node, filter func(node resNode.Node) bool) (nt *Tree, err error) {
	nt = Copy(t)

	giStack := stack.New()

	giStack.Push(gNodeAndIndex{
		gNode: &nt.Root,
		index: 0,
	})

	for {
		if giStack.Len() == 0 {
			break
		}

		gi, _ := giStack.Pop().(gNodeAndIndex)

		if gi.gNode.Children != nil {
			if gi.index < len(gi.gNode.Children) {
				// push current group
				giStack.Push(gNodeAndIndex{
					gNode: gi.gNode,
					index: gi.index + 1,
				})

				// push a sub group
				giStack.Push(gNodeAndIndex{
					gNode: gi.gNode.Children[gi.index],
					index: 0,
				})
			} else {
				//At the end of Children
				newChildren := make([]*resGNode.ResGNode, 0, len(gi.gNode.Children))

				// scan all current gNode children
				// keep all node children which eligible giving filter func
				// keep all group children which children is not empty
				for _, child := range gi.gNode.Children {
					if child.ID < resGNode.GroupBase {
						ni, ok := nodesMap[child.ID]
						if !ok {
							return nil,
								fmt.Errorf("GNode info (id=%d) does not exist in the NodesMap", child.ID)
						}
						if filter(*ni) == true {
							// reserve this node
							newChildren = append(newChildren, child)
						}
					} else {
						if len(child.Children) != 0 {
							// reserve this type=group node
							newChildren = append(newChildren, child)
						}
					}
				}
				gi.gNode.Children = newChildren
			}
		}
	}
	return nt, nil
}

// Filtrate the tree and mark all unusable node (Disabled=true)
func FiltrateMark(t *Tree, nodesMap map[int64]resNode.Node, filter func(node resNode.Node) bool) (nt *Tree, err error) {
	nt = Copy(t)

	giStack := stack.New()

	giStack.Push(gNodeAndIndex{
		gNode: &nt.Root,
		index: 0,
	})

	for {
		if giStack.Len() == 0 {
			break
		}

		gi, _ := giStack.Pop().(gNodeAndIndex)

		if gi.gNode.Children != nil {
			if gi.index < len(gi.gNode.Children) {
				// push current group
				giStack.Push(gNodeAndIndex{
					gNode: gi.gNode,
					index: gi.index + 1,
				})

				// push a sub group
				giStack.Push(gNodeAndIndex{
					gNode: gi.gNode.Children[gi.index],
					index: 0,
				})
			} else {
				//At the end of Children

				mark := true
				// scan all current gNode children
				for index, child := range gi.gNode.Children {
					if child.ID < resGNode.GroupBase {
						ni, ok := nodesMap[child.ID]
						if !ok {
							return nil,
								fmt.Errorf("GNode info (id=%d) does not exist in the NodesMap", child.ID)
						}
						if filter(ni) == true {
							gi.gNode.Children[index].Disabled = false
						} else {
							// mark this node disabled
							gi.gNode.Children[index].Disabled = true
						}
					}
					mark = mark && gi.gNode.Children[index].Disabled
				}
				gi.gNode.Disabled = mark
			}
		}
	}
	return nt, nil
}
