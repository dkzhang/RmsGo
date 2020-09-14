package resGTreeDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNodeTree"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
)

type ResGTreeDM struct {
	TheTree resGNodeTree.Tree
	nodeDM  resNodeDM.ResNodeDM
}

func NewResGTreeDM(ndm resNodeDM.ResNodeDM, jsonFilename string) (ResGTreeDM, error) {
	rtdm := ResGTreeDM{
		TheTree: resGNodeTree.Tree{},
		nodeDM:  ndm,
	}

	tree, err := resGNodeTree.LoadFromJsonFile(jsonFilename)
	if err != nil {
		return ResGTreeDM{}, fmt.Errorf("resGNodeTree.LoadFromJsonFile error: %v", err)
	}

	nodesMap, err := rtdm.nodeDM.QueryAllMap()
	if err != nil {
		return ResGTreeDM{}, fmt.Errorf("nodeDM.QueryAll error: %v", err)
	}

	err = resGNodeTree.SynchronizeNodesInfo(&tree, nodesMap)
	if err != nil {
		return ResGTreeDM{}, fmt.Errorf("SynchronizeNodesInfo error: %v", err)
	}

	resGNodeTree.Count(&tree)

	rtdm.TheTree = tree
	return rtdm, nil
}

func (rtdm ResGTreeDM) QueryTree(nodeFilter func(node resNode.Node) bool) (jsonTree string, err error) {
	nodesMap, err := rtdm.nodeDM.QueryAllMap()
	if err != nil {
		return "", fmt.Errorf("nodeDM.QueryAll error: %v", err)
	}

	err = resGNodeTree.SynchronizeNodesInfo(&rtdm.TheTree, nodesMap)
	if err != nil {
		return "", fmt.Errorf("SynchronizeNodesInfo error: %v", err)
	}

	treeF, err := resGNodeTree.Filtrate(&rtdm.TheTree, nodesMap, nodeFilter)

	resGNodeTree.Count(treeF)

	jsonTree, err = resGNodeTree.ToJson(*treeF)
	if err != nil {
		return "", fmt.Errorf("resGNodeTree.ToJson error: %v", err)
	}

	return jsonTree, nil
}

func (rtdm ResGTreeDM) QueryTreeAllocated(projectID int) (jsonTree string, err error) {
	return rtdm.QueryTree(func(node resNode.Node) bool {
		return node.ProjectID == projectID
	})
}

func (rtdm ResGTreeDM) QueryTreeIdleAndAllocated(projectID int) (jsonTree string, err error) {
	return rtdm.QueryTree(func(node resNode.Node) bool {
		return node.ProjectID == 0 || node.ProjectID == projectID
	})
}

func (rtdm ResGTreeDM) QueryTreeAll() (jsonTree string, err error) {
	return rtdm.QueryTree(func(node resNode.Node) bool {
		return true
	})
}
