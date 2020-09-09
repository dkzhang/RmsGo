package resTreeDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeTree"
)

type ResTreeDM struct {
	TheTree resNodeTree.Tree
}

func NewResTreeDM() ResTreeDM {
	return ResTreeDM{}
}

func (rtdm ResTreeDM) QueryTreeAllocated(projectID int) (jsonTree string, err error) {
	treeF := resNodeTree.FiltrateTree(&rtdm.TheTree, func(node resNode.Node) bool {
		if node.ProjectID == projectID {
			return true
		} else {
			return false
		}
	})

	resNodeTree.CountTree(treeF)

	jsonTree, err = resNodeTree.TreeToJson(*treeF)
	if err != nil {
		return "", fmt.Errorf("resNodeTree.TreeToJson error: %v", err)
	}

	return jsonTree, nil
}

func (rtdm ResTreeDM) QueryTreeIdleAndAllocated(projectID int) (jsonTree string, err error) {
	treeF := resNodeTree.FiltrateTree(&rtdm.TheTree, func(node resNode.Node) bool {
		if node.ProjectID == 0 || node.ProjectID == projectID {
			return true
		} else {
			return false
		}
	})

	resNodeTree.CountTree(treeF)

	jsonTree, err = resNodeTree.TreeToJson(*treeF)
	if err != nil {
		return "", fmt.Errorf("resNodeTree.TreeToJson error: %v", err)
	}

	return jsonTree, nil
}

func (rtdm ResTreeDM) QueryTreeAll() (jsonTree string, err error) {
	treeF := resNodeTree.FiltrateTree(&rtdm.TheTree, func(node resNode.Node) bool { return true })

	resNodeTree.CountTree(treeF)

	jsonTree, err = resNodeTree.TreeToJson(*treeF)
	if err != nil {
		return "", fmt.Errorf("resNodeTree.TreeToJson error: %v", err)
	}

	return jsonTree, nil
}
