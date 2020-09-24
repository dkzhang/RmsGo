package resGTreeDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNodeTree"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
)

type ResGTreeDM struct {
	theTree resGNodeTree.Tree
	nodeDM  resNodeDM.ResNodeDM
}

func NewResGTreeDM(ndm resNodeDM.ResNodeDM, jsonFilename string) (ResGTreeDM, error) {
	rtdm := ResGTreeDM{
		theTree: resGNodeTree.Tree{},
		nodeDM:  ndm,
	}

	tree, err := resGNodeTree.LoadFromJsonFile(jsonFilename)
	if err != nil {
		return ResGTreeDM{}, fmt.Errorf("resGNodeTree.LoadFromJsonFile error: %v", err)
	}

	nodesMap, err := rtdm.nodeDM.GetAllMap()
	if err != nil {
		return ResGTreeDM{}, fmt.Errorf("nodeDM.GetAllMap error: %v", err)
	}

	err = resGNodeTree.SynchronizeNodesInfo(&tree, nodesMap)
	if err != nil {
		return ResGTreeDM{}, fmt.Errorf("SynchronizeNodesInfo error: %v", err)
	}

	resGNodeTree.Count(&tree)

	rtdm.theTree = tree
	return rtdm, nil
}

func (rtdm ResGTreeDM) QueryTree(treeFormat int, nodeFilter func(node resNode.Node) bool) (t *resGNodeTree.Tree, err error) {
	nodesMap, err := rtdm.nodeDM.GetAllMap()
	if err != nil {
		return nil, fmt.Errorf("nodeDM.GetAllMap error: %v", err)
	}

	switch treeFormat {
	case typeCutOut:
		t, err = resGNodeTree.Filtrate(&rtdm.theTree, nodesMap, nodeFilter)
		if err != nil {
			return nil, fmt.Errorf("filtrate tree typeCutOut error: %v", err)
		}
	case typeDisable:
		t, err = resGNodeTree.FiltrateMark(&rtdm.theTree, nodesMap, nodeFilter)
		if err != nil {
			return nil, fmt.Errorf("filtrate tree typeDisable error: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported tree format: %d", treeFormat)
	}
	resGNodeTree.Count(t)

	return t, nil
}

func (rtdm ResGTreeDM) QueryTreeAll() (t *resGNodeTree.Tree) {
	return &rtdm.theTree
}

const (
	typeCutOut  = 1
	typeDisable = 2
)
