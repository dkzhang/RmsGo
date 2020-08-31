package resNodeTree

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func LoadTreeFromJson(str string) (t Tree, err error) {
	t.RootGroup, err = LoadGroupFromJson(str)
	if err != nil {
		return Tree{}, fmt.Errorf("LoadFromJson error since LoadGroupFromJson error: %v", err)
	}

	//GenNodesMap
	t.NodesMap = make(map[int]*Node)
	//t.NodesMap, err = genNodesMap(&(t.RootGroup), t.NodesMap)
	t.NodesMap, err = genNodesMapIteration(&(t.RootGroup))
	if err != nil {
		return Tree{}, fmt.Errorf("tree genNodesMap error: %v", err)
	}
	return t, nil
}

// syn the nodes info from DB to the Tree struct
func SynchronizeNodesInfo(t *Tree, nodes []Node) (err error) {
	for _, n := range nodes {
		if _, ok := t.NodesMap[n.ID]; ok {
			t.NodesMap[n.ID].Name = n.Name
			t.NodesMap[n.ID].Description = n.Description
			t.NodesMap[n.ID].ProjectID = n.ProjectID
			t.NodesMap[n.ID].AllocatedTime = n.AllocatedTime
		}
	}
	return nil
}

// generate a NodesMap from the Tree By iteration
type groupAndIndex struct {
	group *Group
	index int
}

func genNodesMapIteration(g *Group) (map[int]*Node, error) {
	nodesMap := make(map[int]*Node)
	giStack := stack.New()

	giStack.Push(groupAndIndex{
		group: g,
		index: 0,
	})

	for {
		if giStack.Len() == 0 {
			break
		}

		gi, _ := giStack.Pop().(groupAndIndex)

		if gi.index == 0 {
			if gi.group.Nodes != nil {
				for _, node := range gi.group.Nodes {
					if _, ok := nodesMap[node.ID]; ok {
						return nil,
							fmt.Errorf(" Group<%d> error: duplicated ID <%d> detected", gi.group.ID, node.ID)
					}
					nodesMap[node.ID] = node
				}
			}
		}

		if gi.group.SubGroups != nil && gi.index < len(gi.group.SubGroups) {
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
		}
	}
	return nodesMap, nil
}

// generate a NodesMap from the Tree
func genNodesMap(g *Group, nm map[int]*Node) (map[int]*Node, error) {
	var err error
	if g.Nodes != nil {
		for _, node := range g.Nodes {
			if _, ok := nm[node.ID]; ok {
				return nil, fmt.Errorf(" Group<%d> error: duplicated ID <%d> detected", g.ID, node.ID)
			}
			nm[node.ID] = node
		}
	}

	if g.SubGroups != nil {
		for _, sg := range g.SubGroups {
			nm, err = genNodesMap(sg, nm)
			if err != nil {
				return nil, fmt.Errorf(" Group<%d>-subGroup<<%d>> error: %v", g.ID, sg.ID, err)
			}
		}
	}
	return nm, nil
}
