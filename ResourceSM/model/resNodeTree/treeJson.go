package resNodeTree

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeGroup"
)

func TreeToJson(t Tree) (string, error) {
	return resNodeGroup.GroupToJson(t.RootGroup)
}

func TreeToJsonIndent(t Tree) (string, error) {
	return resNodeGroup.GroupToJsonIndent(t.RootGroup)
}
