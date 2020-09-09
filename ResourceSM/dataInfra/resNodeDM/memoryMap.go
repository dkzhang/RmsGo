package resNodeDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/sirupsen/logrus"
)

type MemoryMap struct {
	infoMap      map[int64]*resNode.Node
	theResNodeDB resNodeDB.ResNodeDB
}

func NewMemoryMap(rndb resNodeDB.ResNodeDB, theLogMap logMap.LogMap) (nmm MemoryMap, err error) {
	nmm.theResNodeDB = rndb
	nmm.infoMap = make(map[int64]*resNode.Node)

	nodes, err := nmm.theResNodeDB.QueryAll()
	if err != nil {
		return MemoryMap{},
			fmt.Errorf("generate new MemoryMap failed since ResNodeDB.QueryAll error: %v", err)
	}
	theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"len(Info Array)": len(nodes),
	}).Info("NewMemoryMap ResNodeDB.QueryAll success.")

	for _, node := range nodes {
		nmm.infoMap[node.ID] = &node
	}

	theLogMap.Log(logMap.NORMAL).Info("NewMemoryMap load data to map success.")

	return nmm, nil
}

func (rnm MemoryMap) QueryByID(nodeID int64) (resNode.Node, error) {
	if node, ok := rnm.infoMap[nodeID]; ok {
		return *node, nil
	} else {
		return resNode.Node{},
			fmt.Errorf("the project res (id = %d) info does not exist", nodeID)
	}
}

func (rnm MemoryMap) QueryAll() (nodes []resNode.Node, err error) {
	for _, pn := range rnm.infoMap {
		nodes = append(nodes, *pn)
	}
	return nodes, nil
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
	for _, node := range nodes {
		rnm.infoMap[node.ID] = &node
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
