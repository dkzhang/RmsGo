package resNodeTree

import (
	"fmt"
)

type Tree struct {
	RootGroup Group
	NodesMap  map[int]*Node
}

func LoadTreeFromJson(str string) (t Tree, err error) {
	t.RootGroup, err = LoadGroupFromJson(str)
	if err != nil {
		return Tree{}, fmt.Errorf("LoadFromJson error since LoadGroupFromJson error: %v", err)
	}

	//GenNodesMap
	t.NodesMap = make(map[int]*Node)
	t.NodesMap, err = genNodesMap(&(t.RootGroup), t.NodesMap)
	if err != nil {
		return Tree{}, fmt.Errorf("tree genNodesMap error: %v", err)
	}
	return t, nil
}

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

func TreeToJson(t Tree) (string, error) {
	return GroupToJson(t.RootGroup)
}

func TreeToJsonIndent(t Tree) (string, error) {
	return GroupToJsonIndent(t.RootGroup)
}
