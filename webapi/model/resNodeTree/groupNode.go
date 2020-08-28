package resNodeTree

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	ID             int         `json:"group_id"`
	Name           string      `json:"group_name"`
	Status         int         `json:"group_status"`
	Description    string      `json:"description"`
	SubGroups      []*Group    `json:"sub_groups"`
	Nodes          []*Node     `json:"nodes"`
	NodesNum       int         `json:"nodes_num"`
	NodesStatusMap map[int]int `json:"nodes_status_map"`
}

type Node struct {
	ID          int    `db:"node_id" json:"node_id"`
	Name        string `db:"node_name" json:"node_name"`
	Status      int    `db:"node_status" json:"node_status"`
	Description string `db:"description" json:"description"`
	ProjectID   int    `db:"project_id" json:"project_id"`
}

///////////////////////////////////////////////////////////////////////////////

const (
	StatusUnselected        = 1
	StatusPartiallySelected = 2
	StatusSelected          = 4
)

///////////////////////////////////////////////////////////////////////////////

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

func GroupToJson(g Group) (string, error) {
	bj, err := json.Marshal(g)
	if err != nil {
		return "", fmt.Errorf("json Marshal Group error: %v", err)
	}
	return string(bj), nil
}

func GroupToJsonIndent(g Group) (string, error) {
	bj, err := json.MarshalIndent(g, "", "    ")
	if err != nil {
		return "", fmt.Errorf("json Marshal Group error: %v", err)
	}
	return string(bj), nil
}

func LoadGroupFromJson(str string) (g Group, err error) {
	err = json.Unmarshal([]byte(str), &g)
	if err != nil {
		return Group{}, fmt.Errorf("json UnMarshal Group error: %v", err)
	}
	return g, nil
}
