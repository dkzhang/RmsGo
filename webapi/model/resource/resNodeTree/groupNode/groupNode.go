package groupNode

import "time"

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
	ID            int       `db:"node_id" json:"node_id"`
	Name          string    `db:"node_name" json:"node_name"`
	Status        int       `db:"node_status" json:"node_status"`
	Description   string    `db:"description" json:"description"`
	ProjectID     int       `db:"project_id" json:"project_id"`
	AllocatedTime time.Time `db:"allocated_time" json:"allocated_time"`
}

///////////////////////////////////////////////////////////////////////////////

const (
	StatusUnselected        = 1
	StatusPartiallySelected = 2
	StatusSelected          = 4
)
