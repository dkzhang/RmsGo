package resNodeTree

import "github.com/dkzhang/RmsGo/webapi/model/resource/resNodeTree/groupNode"

func TreeToJson(t Tree) (string, error) {
	return groupNode.GroupToJson(t.RootGroup)
}

func TreeToJsonIndent(t Tree) (string, error) {
	return groupNode.GroupToJsonIndent(t.RootGroup)
}
