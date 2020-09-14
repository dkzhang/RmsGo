package resGNodeTree_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNodeTree"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Tree", func() {
	var rootGNode resGNode.ResGNode
	BeforeEach(func() {
		rootGNode = resGNode.ResGNode{
			ID:       1e4,
			Label:    "偏移云",
			Children: nil,
		}

		lc := &resGNode.ResGNode{
			ID:       11e4,
			Label:    "浪潮云",
			Children: nil,
		}
		rootGNode.Children = append(rootGNode.Children, lc)

		tempGroup := &resGNode.ResGNode{
			ID:       111e4,
			Label:    "Group1-1",
			Children: nil,
		}
		for i := int64(0); i < 256; i++ {
			p := &resGNode.ResGNode{
				ID:       i,
				Label:    fmt.Sprintf("Node%d", i),
				Children: nil,
			}
			tempGroup.Children = append(tempGroup.Children, p)

			if i%32 == 31 {

				lc.Children = append(lc.Children, tempGroup)

				groupID := (i+1)/32 + 1
				tempGroup = &resGNode.ResGNode{
					ID:       110e4 + groupID*1e4,
					Label:    fmt.Sprintf("Group1-%d", groupID),
					Children: nil,
				}
			}
		}

	})
	Context("Tree to Json", func() {
		It("Tree to Json", func() {
			t := resGNodeTree.Tree{
				Root:     rootGNode,
				NodesNum: 0,
			}

			_, err := resGNodeTree.ToJsonIndent(t)
			Expect(err).ShouldNot(HaveOccurred())
			//By(fmt.Sprintf("Tree to Json = %s", str))
		})
	})

	Context("Load Tree from Json", func() {
		It("Load Tree from Json", func() {
			t, err := resGNodeTree.LoadFromJsonFile("./tree256.json")
			Expect(err).ShouldNot(HaveOccurred())
			nodesNum := resGNodeTree.CountRO(&t)
			By(fmt.Sprintf("Count Tree Load from file = %d", nodesNum))
			resGNodeTree.Count(&t)
			Expect(t.NodesNum).Should(Equal(nodesNum))
		})
	})

	Context("Filtrate Tree", func() {
		It("Filtrate Tree", func() {
			t := resGNodeTree.Tree{
				Root:     rootGNode,
				NodesNum: 0,
			}

			nodesMap := make(map[int64]resNode.Node, 256)
			for i := int64(0); i < 256; i++ {
				nodesMap[i] = resNode.Node{
					ID:            i,
					ProjectID:     0,
					AllocatedTime: time.Now(),
				}
			}
			for j := int64(50); j < 100; j++ {
				nodesMap[j] = resNode.Node{
					ID:            j,
					ProjectID:     1,
					AllocatedTime: time.Now(),
				}
			}

			for j := int64(100); j < 200; j++ {
				nodesMap[j] = resNode.Node{
					ID:            j,
					ProjectID:     2,
					AllocatedTime: time.Now(),
				}
			}

			unArray, err := resGNodeTree.FiltrateMark(&t, nodesMap, func(node resNode.Node) bool {
				return node.ProjectID == 0 || node.ProjectID == 1
			})
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("unusable array = %v", unArray))

			nt, err := resGNodeTree.Filtrate(&t, nodesMap, func(node resNode.Node) bool {
				return node.ProjectID == 0 || node.ProjectID == 1
			})
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("t nodesNum = %d", resGNodeTree.CountRO(&t)))
			By(fmt.Sprintf("nt nodesNum = %d", resGNodeTree.CountRO(nt)))

			//str, err := resGNodeTree.ToJsonIndent(*nt)
			//Expect(err).ShouldNot(HaveOccurred())
			//By(fmt.Sprintf("Tree to Json = %s", str))

		})
	})

	Context("Syn Tree", func() {
		It("Syn Tree", func() {
			t := resGNodeTree.Tree{
				Root:     rootGNode,
				NodesNum: 0,
			}

			nodesMap := make(map[int64]resNode.Node, 256)
			for i := int64(0); i < 256; i++ {
				nodesMap[i] = resNode.Node{
					ID:            i,
					Name:          fmt.Sprintf("Node%dNew", i),
					ProjectID:     0,
					AllocatedTime: time.Now(),
				}
			}

			err := resGNodeTree.SynchronizeNodesInfo(&t, nodesMap)
			Expect(err).ShouldNot(HaveOccurred())

			str, err := resGNodeTree.ToJsonIndent(t)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("Tree to Json = %s", str))

		})
	})
})
