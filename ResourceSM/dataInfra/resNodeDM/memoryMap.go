package resNodeDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/myUtils/deepCopy"
)

type MemoryMap struct {
	infoMap      map[int64]*resNode.Node
	theResNodeDB resNodeDB.ResNodeDB
}

func NewMemoryMap(rndb resNodeDB.ResNodeDB) (nmm MemoryMap, err error) {
	nmm.theResNodeDB = rndb
	nmm.infoMap = make(map[int64]*resNode.Node)

	nodes, err := nmm.theResNodeDB.QueryAll()
	if err != nil {
		return MemoryMap{},
			fmt.Errorf("generate new MemoryMap failed since ResNodeDB.QueryAll error: %v", err)
	}

	for i := range nodes {
		nmm.infoMap[nodes[i].ID] = &nodes[i]
	}

	return nmm, nil
}

func (rnm MemoryMap) QueryByID(nodeID int64) (resNode.Node, error) {
	if node, ok := rnm.infoMap[nodeID]; ok {
		return *node, nil
	} else {
		return resNode.Node{},
			fmt.Errorf("resNodeDM QueryByID error, resNode (id = %d) info does not exist", nodeID)
	}
}

func (rnm MemoryMap) GetAllArray() (nodes []resNode.Node, err error) {
	for _, pn := range rnm.infoMap {
		nodes = append(nodes, *pn)
	}
	return nodes, nil
}

func (rnm MemoryMap) GetAllMap() (nodesMap map[int64]*resNode.Node, err error) {
	err = deepCopy.DeepCopy(&nodesMap, &rnm.infoMap)
	if err != nil {
		return nil, fmt.Errorf("DeepCopy error=%v", err)
	}
	return nodesMap, nil
}

func (rnm MemoryMap) Update(node resNode.Node) (err error) {
	//update in db
	err = rnm.theResNodeDB.Update(node)
	if err != nil {
		return fmt.Errorf("ProjectResDB.Update error: %v", err)
	}

	//update in memoryMap
	rnm.infoMap[node.ID] = &node

	return nil
}

func (rnm MemoryMap) UpdateNodes(nodes []resNode.Node) (err error) {
	//update in db
	err = rnm.theResNodeDB.UpdateNodes(nodes)
	if err != nil {
		return fmt.Errorf("ProjectResDB.UpdateNodes error: %v", err)
	}

	//update in memoryMap
	for i := range nodes {
		rnm.infoMap[nodes[i].ID] = &nodes[i]
	}
	return nil
}

func (rnm MemoryMap) Insert(node resNode.Node) (err error) {
	//insert in db
	err = rnm.theResNodeDB.Insert(node)
	if err != nil {
		return fmt.Errorf("ProjectResDB.Insert error: %v", err)
	}

	//insert in memoryMap
	rnm.infoMap[node.ID] = &node

	return nil
}
