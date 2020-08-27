package resNodeTree

type Group struct {
	ID          int     `json:"group_id"`
	Name        string  `json:"group_name"`
	Status      int     `json:"group_status"`
	Description string  `json:"description"`
	SubGroups   []Group `json:"sub_groups"`
	Nodes       []Node  `json:"nodes"`
}

type Node struct {
	ID          int    `json:"node_id"`
	Name        string `json:"node_name"`
	Status      int    `json:"node_status"`
	Description string `json:"description"`
}

func Count(g Group) (nodesNum int, nodesStatusMap map[int]int) {
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
			num, sMap := Count(sg)
			nodesNum += num
			for k, v := range sMap {
				nodesStatusMap[k] += v
			}
		}
	}
	return nodesNum, nodesStatusMap
}

func OrganizeGroupStatus(g Group) (status int) {
	num, sMap := Count(g)
	if num != 0 {
		if sMap[StatusUnselected] == 0 {
			status = StatusSelected
		} else if sMap[StatusSelected] == 0 {
			status = StatusUnselected
		} else {
			status = StatusPartiallySelected
		}
	}
	return status
}

const (
	StatusUnselected        = 1
	StatusPartiallySelected = 2
	StatusSelected          = 4
)
