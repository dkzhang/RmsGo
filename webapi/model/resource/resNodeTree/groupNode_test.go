package resNodeTree_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/resource/resNodeTree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupNode", func() {
	var rootGroup resNodeTree.Group

	BeforeEach(func() {
		node1 := resNodeTree.Node{ID: 1, Name: "Node1", Status: 4, ProjectID: 2}
		node2 := resNodeTree.Node{ID: 2, Name: "Node2", Status: 1, ProjectID: 2}
		node3 := resNodeTree.Node{ID: 3, Name: "Node3", Status: 1, ProjectID: 3}
		node4 := resNodeTree.Node{ID: 4, Name: "Node4", Status: 4, ProjectID: 1}
		node5 := resNodeTree.Node{ID: 5, Name: "Node5", Status: 1, ProjectID: 3}
		node6 := resNodeTree.Node{ID: 6, Name: "Node6", Status: 4, ProjectID: 1}
		node7 := resNodeTree.Node{ID: 7, Name: "Node7", Status: 4, ProjectID: 2}
		node8 := resNodeTree.Node{ID: 8, Name: "Node8", Status: 1, ProjectID: 0}
		node9 := resNodeTree.Node{ID: 9, Name: "Node9", Status: 4, ProjectID: 0}
		node10 := resNodeTree.Node{ID: 10, Name: "Node10", Status: 4, ProjectID: 1}

		Group101 := resNodeTree.Group{
			ID:    101,
			Name:  "Group101",
			Nodes: []*resNodeTree.Node{&node1, &node2},
		}

		Group1 := resNodeTree.Group{
			ID:        1,
			Name:      "Group1",
			SubGroups: []*resNodeTree.Group{&Group101},
			Nodes:     []*resNodeTree.Node{&node3, &node4, &node5},
		}

		Group2 := resNodeTree.Group{
			ID:    2,
			Name:  "Group2",
			Nodes: []*resNodeTree.Node{&node6, &node7},
		}

		rootGroup = resNodeTree.Group{
			ID:        0,
			Name:      "rootGroup",
			SubGroups: []*resNodeTree.Group{&Group1, &Group2},
			Nodes:     []*resNodeTree.Node{&node8, &node9, &node10},
		}
	})

	Context("CountGroup", func() {
		It("GroupNode", func() {
			resNodeTree.CountGroup(&rootGroup)
			Expect(rootGroup.NodesNum).Should(Equal(10))
			Expect(rootGroup.NodesStatusMap[1]).Should(Equal(4))
			Expect(rootGroup.NodesStatusMap[4]).Should(Equal(6))
			By(fmt.Sprintf("num = %d", rootGroup.NodesNum))
			By(fmt.Sprintf("StatusMap = %v", rootGroup.NodesStatusMap))
			By(fmt.Sprintf("Status = %d", rootGroup.Status))

			strJson, err := resNodeTree.GroupToJsonIndent(rootGroup)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("strJson = %s", strJson))

		})
	})

	Context("CopyTree", func() {
		It("CopyTree", func() {
			strJson, err := resNodeTree.GroupToJsonIndent(rootGroup)
			Expect(err).ShouldNot(HaveOccurred())

			t1, err := resNodeTree.LoadTreeFromJson(strJson)
			Expect(err).ShouldNot(HaveOccurred())
			resNodeTree.CountGroup(&t1.RootGroup)

			tCopy := resNodeTree.CopyTree(&t1)
			resNodeTree.CountGroup(&tCopy.RootGroup)

			Expect(t1.RootGroup.NodesNum).Should(Equal(tCopy.RootGroup.NodesNum))
			Expect(t1.RootGroup.NodesStatusMap[resNodeTree.StatusSelected]).
				Should(Equal(tCopy.RootGroup.NodesStatusMap[resNodeTree.StatusSelected]))

			strJson2, err := resNodeTree.TreeToJsonIndent(*tCopy)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("strJson2 = %s", strJson2))

		})
	})

	Context("FiltrateTree", func() {
		It("FiltrateTree", func() {
			strJson, err := resNodeTree.GroupToJsonIndent(rootGroup)
			Expect(err).ShouldNot(HaveOccurred())

			t1, err := resNodeTree.LoadTreeFromJson(strJson)
			Expect(err).ShouldNot(HaveOccurred())
			resNodeTree.CountGroup(&t1.RootGroup)

			t2 := resNodeTree.FiltrateTree(&t1, 3)
			resNodeTree.CountGroup(&t2.RootGroup)
			strJson2, err := resNodeTree.TreeToJsonIndent(*t2)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("strJson2 = %s", strJson2))
		})
	})
})
