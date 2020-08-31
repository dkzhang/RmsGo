package resNodeTree

import "github.com/dkzhang/RmsGo/webapi/model/resource/resNodeTree/groupNode"

type Tree struct {
	RootGroup groupNode.Group
	NodesMap  map[int]*groupNode.Node
}
