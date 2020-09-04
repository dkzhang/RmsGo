package resNodeGroup

import "github.com/dkzhang/RmsGo/ResourceSM/model/resNode"

type Group struct {
	ID             int             `json:"group_id"`
	Name           string          `json:"group_name"`
	Status         int             `json:"group_status"`
	Description    string          `json:"description"`
	SubGroups      []*Group        `json:"sub_groups"`
	Nodes          []*resNode.Node `json:"nodes"`
	NodesNum       int             `json:"nodes_num"`
	NodesStatusMap map[int]int     `json:"nodes_status_map"`
}

const (
	StatusPartiallySelected = 4
)
