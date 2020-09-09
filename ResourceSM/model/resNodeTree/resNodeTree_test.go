package resNodeTree_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeGroup"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNodeTree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupNode", func() {
	var rootGroup resNodeGroup.Group

	BeforeEach(func() {
		node1 := resNode.Node{ID: 1, Name: "Node1", Status: 4, ProjectID: 2}
		node2 := resNode.Node{ID: 2, Name: "Node2", Status: 1, ProjectID: 2}
		node3 := resNode.Node{ID: 3, Name: "Node3", Status: 1, ProjectID: 3}
		node4 := resNode.Node{ID: 4, Name: "Node4", Status: 4, ProjectID: 1}
		node5 := resNode.Node{ID: 5, Name: "Node5", Status: 1, ProjectID: 3}
		node6 := resNode.Node{ID: 6, Name: "Node6", Status: 4, ProjectID: 1}
		node7 := resNode.Node{ID: 7, Name: "Node7", Status: 4, ProjectID: 2}
		node8 := resNode.Node{ID: 8, Name: "Node8", Status: 1, ProjectID: 0}
		node9 := resNode.Node{ID: 9, Name: "Node9", Status: 4, ProjectID: 0}
		node10 := resNode.Node{ID: 10, Name: "Node10", Status: 4, ProjectID: 1}

		Group101 := resNodeGroup.Group{
			ID:    101,
			Name:  "Group101",
			Nodes: []*resNode.Node{&node1, &node2},
		}

		Group1 := resNodeGroup.Group{
			ID:        1,
			Name:      "Group1",
			SubGroups: []*resNodeGroup.Group{&Group101},
			Nodes:     []*resNode.Node{&node3, &node4, &node5},
		}

		Group2 := resNodeGroup.Group{
			ID:    2,
			Name:  "Group2",
			Nodes: []*resNode.Node{&node6, &node7},
		}

		rootGroup = resNodeGroup.Group{
			ID:        0,
			Name:      "rootGroup",
			SubGroups: []*resNodeGroup.Group{&Group1, &Group2},
			Nodes:     []*resNode.Node{&node8, &node9, &node10},
		}
	})

	Context("CountGroup", func() {
		It("GroupNode", func() {
			resNodeGroup.CountGroup(&rootGroup)
			Expect(rootGroup.NodesNum).Should(Equal(10))
			Expect(rootGroup.NodesStatusMap[1]).Should(Equal(4))
			Expect(rootGroup.NodesStatusMap[4]).Should(Equal(6))
			By(fmt.Sprintf("num = %d", rootGroup.NodesNum))
			By(fmt.Sprintf("StatusMap = %v", rootGroup.NodesStatusMap))
			By(fmt.Sprintf("Status = %d", rootGroup.Status))

			strJson, err := resNodeGroup.GroupToJsonIndent(rootGroup)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("strJson = %s", strJson))

		})
	})

	Context("CopyTree", func() {
		It("CopyTree", func() {
			strJson, err := resNodeGroup.GroupToJsonIndent(rootGroup)
			Expect(err).ShouldNot(HaveOccurred())

			t1, err := resNodeTree.LoadTreeFromJson(strJson)
			Expect(err).ShouldNot(HaveOccurred())
			resNodeGroup.CountGroup(&t1.RootGroup)

			tCopy := resNodeTree.CopyTree(&t1)
			resNodeGroup.CountGroup(&tCopy.RootGroup)

			Expect(t1.RootGroup.NodesNum).Should(Equal(tCopy.RootGroup.NodesNum))
			Expect(t1.RootGroup.NodesStatusMap[resNode.StatusSelected]).
				Should(Equal(tCopy.RootGroup.NodesStatusMap[resNode.StatusSelected]))

			strJson2, err := resNodeTree.TreeToJsonIndent(*tCopy)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("strJson2 = %s", strJson2))

		})
	})

	Context("FiltrateTree", func() {
		It("FiltrateTree", func() {
			strJson, err := resNodeGroup.GroupToJsonIndent(rootGroup)
			Expect(err).ShouldNot(HaveOccurred())

			t1, err := resNodeTree.LoadTreeFromJson(strJson)
			Expect(err).ShouldNot(HaveOccurred())
			resNodeGroup.CountGroup(&t1.RootGroup)
			resNodeTree.CountTree(&t1)

			t2 := resNodeTree.FiltrateTree(&t1, func(node resNode.Node) bool {
				if node.ProjectID == 0 || node.ProjectID == 3 {
					return true
				} else {
					return false
				}
			})
			resNodeTree.CountTree(t2)
			strJson2, err := resNodeTree.TreeToJsonIndent(*t2)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("strJson2 = %s", strJson2))
		})
	})
})
