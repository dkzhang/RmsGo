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
		node1 := resNodeTree.Node{ID: 1, Status: 4}
		node2 := resNodeTree.Node{ID: 2, Status: 1}
		node3 := resNodeTree.Node{ID: 3, Status: 1}
		node4 := resNodeTree.Node{ID: 4, Status: 4}
		node5 := resNodeTree.Node{ID: 5, Status: 1}
		node6 := resNodeTree.Node{ID: 6, Status: 4}
		node7 := resNodeTree.Node{ID: 7, Status: 4}
		node8 := resNodeTree.Node{ID: 8, Status: 1}
		node9 := resNodeTree.Node{ID: 9, Status: 4}
		node10 := resNodeTree.Node{ID: 10, Status: 4}

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

		It("Load from json", func() {
			strJson := ` 
 {
    "group_id": 0,
    "group_name": "rootGroup",
    "group_status": 2,
    "description": "",
    "sub_groups": [
        {
            "group_id": 1,
            "group_name": "Group1",
            "group_status": 2,
            "description": "",
            "sub_groups": [
                {
                    "group_id": 101,
                    "group_name": "Group101",
                    "group_status": 2,
                    "description": "",
                    "sub_groups": null,
                    "nodes": [
                        {
                            "node_id": 1,
                            "node_name": "",
                            "node_status": 4,
                            "description": ""
                        },
                        {
                            "node_id": 2,
                            "node_name": "",
                            "node_status": 1,
                            "description": ""
                        }
                    ],
                    "nodes_num": 2,
                    "nodes_status_map": {
                        "1": 1,
                        "4": 1
                    }
                }
            ],
            "nodes": [
                {
                    "node_id": 3,
                    "node_name": "",
                    "node_status": 1,
                    "description": ""
                },
                {
                    "node_id": 4,
                    "node_name": "",
                    "node_status": 4,
                    "description": ""
                },
                {
                    "node_id": 5,
                    "node_name": "",
                    "node_status": 1,
                    "description": ""
                }
            ],
            "nodes_num": 5,
            "nodes_status_map": {
                "1": 3,
                "4": 2
            }
        },
        {
            "group_id": 2,
            "group_name": "Group2",
            "group_status": 4,
            "description": "",
            "sub_groups": null,
            "nodes": [
                {
                    "node_id": 6,
                    "node_name": "",
                    "node_status": 4,
                    "description": ""
                },
                {
                    "node_id": 7,
                    "node_name": "",
                    "node_status": 4,
                    "description": ""
                }
            ],
            "nodes_num": 2,
            "nodes_status_map": {
                "4": 2
            }
        }
    ],
    "nodes": [
        {
            "node_id": 8,
            "node_name": "",
            "node_status": 1,
            "description": ""
        },
        {
            "node_id": 9,
            "node_name": "",
            "node_status": 4,
            "description": ""
        },
        {
            "node_id": 10,
            "node_name": "",
            "node_status": 4,
            "description": ""
        }
    ],
    "nodes_num": 10,
    "nodes_status_map": {
        "1": 4,
        "4": 6
    }
}
`
			t, err := resNodeTree.LoadTreeFromJson(strJson)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("t = %v", t))
			By(fmt.Sprintf("t.NodesMap = %v", t.NodesMap))

			strJson2, err := resNodeTree.TreeToJsonIndent(t)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("strJson2 = %s", strJson2))

		})
	})
})
