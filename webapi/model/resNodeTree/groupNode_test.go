package resNodeTree_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/resNodeTree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupNode", func() {
	var rootGroup resNodeTree.Group

	BeforeEach(func() {
		node1 := resNodeTree.Node{Status: 4}
		node2 := resNodeTree.Node{Status: 1}
		node3 := resNodeTree.Node{Status: 1}
		node4 := resNodeTree.Node{Status: 4}
		node5 := resNodeTree.Node{Status: 1}
		node6 := resNodeTree.Node{Status: 4}
		node7 := resNodeTree.Node{Status: 1}
		node8 := resNodeTree.Node{Status: 1}
		node9 := resNodeTree.Node{Status: 4}
		node10 := resNodeTree.Node{Status: 4}

		Group101 := resNodeTree.Group{
			Name:  "Group101",
			Nodes: []resNodeTree.Node{node1, node2},
		}

		Group1 := resNodeTree.Group{
			Name:      "Group1",
			SubGroups: []resNodeTree.Group{Group101},
			Nodes:     []resNodeTree.Node{node3, node4, node5},
		}

		Group2 := resNodeTree.Group{
			Name:  "Group2",
			Nodes: []resNodeTree.Node{node6, node7},
		}

		rootGroup = resNodeTree.Group{
			Name:      "rootGroup",
			SubGroups: []resNodeTree.Group{Group1, Group2},
			Nodes:     []resNodeTree.Node{node8, node9, node10},
		}
	})

	Context("Count", func() {
		It("GroupNode", func() {
			nodesNum, nodesStatusMap := resNodeTree.Count(rootGroup)
			Expect(nodesNum).Should(Equal(10))
			Expect(nodesStatusMap[1]).Should(Equal(5))
			Expect(nodesStatusMap[4]).Should(Equal(5))
			By(fmt.Sprintf("num = %d", nodesNum))
			By(fmt.Sprintf("StatusMap = %v", nodesStatusMap))
		})
	})
})
