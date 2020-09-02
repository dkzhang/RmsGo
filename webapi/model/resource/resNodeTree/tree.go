package resNodeTree

import (
	"github.com/dkzhang/RmsGo/webapi/model/resource/resNode"
	"github.com/dkzhang/RmsGo/webapi/model/resource/resNodeGroup"
)

type Tree struct {
	RootGroup resNodeGroup.Group
	NodesMap  map[int]*resNode.Node
}
