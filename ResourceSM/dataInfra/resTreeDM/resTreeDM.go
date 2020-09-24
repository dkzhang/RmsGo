package resTreeDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeTree"
)

type ResTreeDM struct {
	TheTree resNodeTree.Tree
	nodeDM  resNodeDM.ResNodeDM
}

func NewResTreeDM(ndm resNodeDM.ResNodeDM, strJson string) (ResTreeDM, error) {
	rtdm := ResTreeDM{
		TheTree: resNodeTree.Tree{},
		nodeDM:  ndm,
	}

	tree, err := resNodeTree.LoadTreeFromJson(strJson)
	if err != nil {
		return ResTreeDM{}, fmt.Errorf("resNodeTree.LoadTreeFromJson error: %v", err)
	}

	nodes, err := rtdm.nodeDM.GetAllArray()
	if err != nil {
		return ResTreeDM{}, fmt.Errorf("nodeDM.GetAllArray error: %v", err)
	}

	err = resNodeTree.SynchronizeNodesInfo(&tree, nodes)
	if err != nil {
		return ResTreeDM{}, fmt.Errorf("SynchronizeNodesInfo error: %v", err)
	}

	resNodeTree.CountTree(&tree)

	rtdm.TheTree = tree
	return rtdm, nil
}

func (rtdm ResTreeDM) QueryTree(nodeFilter func(node resNode.Node) bool) (jsonTree string, err error) {
	nodes, err := rtdm.nodeDM.GetAllArray()
	if err != nil {
		return "", fmt.Errorf("nodeDM.GetAllArray error: %v", err)
	}

	err = resNodeTree.SynchronizeNodesInfo(&rtdm.TheTree, nodes)
	if err != nil {
		return "", fmt.Errorf("SynchronizeNodesInfo error: %v", err)
	}

	treeF := resNodeTree.FiltrateTree(&rtdm.TheTree, nodeFilter)

	resNodeTree.CountTree(treeF)

	jsonTree, err = resNodeTree.TreeToJson(*treeF)
	if err != nil {
		return "", fmt.Errorf("resNodeTree.TreeToJson error: %v", err)
	}

	return jsonTree, nil
}

func (rtdm ResTreeDM) QueryTreeAllocated(projectID int) (jsonTree string, err error) {
	return rtdm.QueryTree(func(node resNode.Node) bool {
		return node.ProjectID == projectID
	})
}

func (rtdm ResTreeDM) QueryTreeIdleAndAllocated(projectID int) (jsonTree string, err error) {
	return rtdm.QueryTree(func(node resNode.Node) bool {
		return node.ProjectID == 0 || node.ProjectID == projectID
	})
}

func (rtdm ResTreeDM) QueryTreeAll() (jsonTree string, err error) {
	return rtdm.QueryTree(func(node resNode.Node) bool {
		return true
	})
}
