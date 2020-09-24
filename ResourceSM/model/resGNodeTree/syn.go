package resGNodeTree

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/sirupsen/logrus"
)

func synNodesInfo(node *resGNode.ResGNode, nodesMap map[int64]resNode.Node) (err error) {
	if node.ID < resGNode.GroupBase {
		_, ok := nodesMap[node.ID]
		if !ok {
			return fmt.Errorf("GNode info (id=%d) does not exist in the NodesMap", node.ID)
		} else {
			node.Label = nodesMap[node.ID].Name
			logrus.Infof("node(id=%d) Label=%s", node.ID, node.Label)
		}
	}

	if node.Children != nil {
		for i, _ := range node.Children {
			err = synNodesInfo(node.Children[i], nodesMap)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SynchronizeNodesInfo(t *Tree, nodesMap map[int64]resNode.Node) (err error) {
	return synNodesInfo(&t.Root, nodesMap)
}

// syn the nodes info from DB to the Tree struct
//func SynchronizeNodesInfo(t *Tree, nodesMap map[int64]resNode.Node) (err error) {
//
//	giStack := stack.New()
//
//	giStack.Push(gNodeAndIndex{
//		gNode: &t.Root,
//		index: 0,
//	})
//
//	for {
//		if giStack.Len() == 0 {
//			break
//		}
//
//		gi, _ := giStack.Pop().(gNodeAndIndex)
//
//		if gi.gNode.Children != nil {
//			if gi.index < len(gi.gNode.Children) {
//				// push current group
//				giStack.Push(gNodeAndIndex{
//					gNode: gi.gNode,
//					index: gi.index + 1,
//				})
//
//				// push a sub group
//				giStack.Push(gNodeAndIndex{
//					gNode: gi.gNode.Children[gi.index],
//					index: 0,
//				})
//			} else {
//				//At the end of Children
//
//				// scan all current gNode children
//				// set all node children label = name from db
//
//				for _, child := range gi.gNode.Children {
//					if child.ID < resGNode.GroupBase {
//						_, ok := nodesMap[child.ID]
//						if !ok {
//							return fmt.Errorf("GNode info (id=%d) does not exist in the NodesMap", child.ID)
//						} else {
//							child.Label = nodesMap[child.ID].Name
//						}
//					}
//				}
//			}
//		}
//	}
//	return nil
//}
