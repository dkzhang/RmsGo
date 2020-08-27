package resNodeTree

type Group struct {
	ID             int         `json:"group_id"`
	Name           string      `json:"group_name"`
	Status         int         `json:"group_status"`
	Description    string      `json:"description"`
	SubGroups      []Group     `json:"sub_groups"`
	Nodes          []Node      `json:"nodes"`
	NodesNum       int         `json:"nodes_num"`
	NodesStatusMap map[int]int `json:"nodes_status_map"`
}

type Node struct {
	ID          int    `json:"node_id"`
	Name        string `json:"node_name"`
	Status      int    `json:"node_status"`
	Description string `json:"description"`
}

func (g Group) Count() {
	g.NodesStatusMap = make(map[int]int)

	for _, node := range g.Nodes {
		g.NodesNum++
		g.NodesStatusMap[node.Status]++
	}
}
