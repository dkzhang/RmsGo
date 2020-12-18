package resGNodeTree_test

import (
	"bufio"
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNodeTree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("CPU Tree", func() {
	var rootGNode resGNode.ResGNode
	BeforeEach(func() {
		rootGNode = resGNode.ResGNode{
			ID:       1 * resGNode.GroupBase,
			Label:    "偏移云",
			Children: nil,
		}

		var pGroup *resGNode.ResGNode
		///////////////////////////////////////////////////////////////////////
		lc1 := &resGNode.ResGNode{
			ID:       11 * resGNode.GroupBase,
			Label:    "浪潮云CPU 1组",
			Children: nil,
		}

		for ig := 287; ig <= 301; ig++ {
			pGroup = &resGNode.ResGNode{
				ID:       int64((1000 + ig) * resGNode.GroupBase),
				Label:    fmt.Sprintf("F%d", ig),
				Children: nil,
			}
			for in := 1; in <= 16; in++ {
				pNode := &resGNode.ResGNode{
					ID:       int64(ig*100 + in),
					Label:    fmt.Sprintf("C%d", in),
					Children: nil,
				}
				pGroup.Children = append(pGroup.Children, pNode)
			}
			lc1.Children = append(lc1.Children, pGroup)
		}

		{
			ig := 302
			pGroup = &resGNode.ResGNode{
				ID:       int64((1000 + ig) * resGNode.GroupBase),
				Label:    fmt.Sprintf("F%d", ig),
				Children: nil,
			}
			for in := 1; in <= 12; in++ {
				pNode := &resGNode.ResGNode{
					ID:       int64(ig*100 + in),
					Label:    fmt.Sprintf("C%d", in),
					Children: nil,
				}
				pGroup.Children = append(pGroup.Children, pNode)
			}
			lc1.Children = append(lc1.Children, pGroup)
		}

		rootGNode.Children = append(rootGNode.Children, lc1)

		///////////////////////////////////////////////////////////////////////
		lc2 := &resGNode.ResGNode{
			ID:       12 * resGNode.GroupBase,
			Label:    "浪潮云CPU 2组",
			Children: nil,
		}

		for ig := 303; ig <= 317; ig++ {
			pGroup = &resGNode.ResGNode{
				ID:       int64((1000 + ig) * resGNode.GroupBase),
				Label:    fmt.Sprintf("F%d", ig),
				Children: nil,
			}
			for in := 1; in <= 16; in++ {
				pNode := &resGNode.ResGNode{
					ID:       int64(ig*100 + in),
					Label:    fmt.Sprintf("C%d", in),
					Children: nil,
				}
				pGroup.Children = append(pGroup.Children, pNode)
			}
			lc2.Children = append(lc2.Children, pGroup)
		}

		{
			ig := 318
			pGroup = &resGNode.ResGNode{
				ID:       int64((1000 + ig) * resGNode.GroupBase),
				Label:    fmt.Sprintf("F%d", ig),
				Children: nil,
			}
			for in := 1; in <= 12; in++ {
				pNode := &resGNode.ResGNode{
					ID:       int64(ig*100 + in),
					Label:    fmt.Sprintf("C%d", in),
					Children: nil,
				}
				pGroup.Children = append(pGroup.Children, pNode)
			}
			lc2.Children = append(lc2.Children, pGroup)
		}

		rootGNode.Children = append(rootGNode.Children, lc2)

	})
	Context("Tree to Json", func() {
		It("Tree to Json", func() {
			t := resGNodeTree.Tree{
				Root:     rootGNode,
				NodesNum: 0,
			}

			str, err := resGNodeTree.ToJsonIndent(t)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("Tree to Json = %s", str))

			// write json text file
			filePath := "tree_cpu.json"
			file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
			Expect(err).ShouldNot(HaveOccurred())

			//及时关闭
			defer file.Close()
			//写入内容

			//写入时，使用带缓存的 *Writer
			writer := bufio.NewWriter(file)
			writer.WriteString(str)

			//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
			//所以要调用 flush方法，将缓存的数据真正写入到文件中。
			writer.Flush()
		})
	})

	//Context("Load Tree from Json", func() {
	//	It("Load Tree from Json", func() {
	//		t, err := resGNodeTree.LoadFromJsonFile("./tree256.json")
	//		Expect(err).ShouldNot(HaveOccurred())
	//		nodesNum := resGNodeTree.CountRO(&t)
	//		By(fmt.Sprintf("Count Tree Load from file = %d", nodesNum))
	//		resGNodeTree.Count(&t)
	//		Expect(t.NodesNum).Should(Equal(nodesNum))
	//	})
	//})
	//
	//Context("Filtrate Tree", func() {
	//	It("Filtrate Tree", func() {
	//		t := resGNodeTree.Tree{
	//			Root:     rootGNode,
	//			NodesNum: 0,
	//		}
	//
	//		nodesMap := make(map[int64]resNode.Node, 256)
	//		for i := int64(0); i < 256; i++ {
	//			nodesMap[i] = resNode.Node{
	//				ID:            i,
	//				ProjectID:     0,
	//				AllocatedTime: time.Now(),
	//			}
	//		}
	//		for j := int64(50); j < 100; j++ {
	//			nodesMap[j] = resNode.Node{
	//				ID:            j,
	//				ProjectID:     1,
	//				AllocatedTime: time.Now(),
	//			}
	//		}
	//
	//		for j := int64(100); j < 200; j++ {
	//			nodesMap[j] = resNode.Node{
	//				ID:            j,
	//				ProjectID:     2,
	//				AllocatedTime: time.Now(),
	//			}
	//		}
	//
	//		nt1, err := resGNodeTree.Filtrate(&t, nodesMap, func(node resNode.Node) bool {
	//			return node.ProjectID == 0 || node.ProjectID == 1
	//		})
	//		Expect(err).ShouldNot(HaveOccurred())
	//		By(fmt.Sprintf("t nodesNum = %d", resGNodeTree.CountRO(&t)))
	//		By(fmt.Sprintf("nt1 nodesNum = %d", resGNodeTree.CountRO(nt1)))
	//
	//		nt2, err := resGNodeTree.FiltrateMark(&t, nodesMap, func(node resNode.Node) bool {
	//			return node.ProjectID == 0 || node.ProjectID == 1
	//		})
	//		Expect(err).ShouldNot(HaveOccurred())
	//		By(fmt.Sprintf("nt2 nodesNum = %d", resGNodeTree.CountRO(nt2)))
	//
	//		str, err := resGNodeTree.ToJsonIndent(*nt2)
	//		Expect(err).ShouldNot(HaveOccurred())
	//		By(fmt.Sprintf("Tree to Json = %s", str))
	//
	//	})
	//})

	//Context("Syn Tree", func() {
	//	It("Syn Tree", func() {
	//		t := resGNodeTree.Tree{
	//			Root:     rootGNode,
	//			NodesNum: 0,
	//		}
	//
	//		nodesMap := make(map[int64]resNode.Node, 256)
	//		for i := int64(0); i < 256; i++ {
	//			nodesMap[i] = resNode.Node{
	//				ID:            i,
	//				Name:          fmt.Sprintf("Node%dNew", i),
	//				ProjectID:     0,
	//				AllocatedTime: time.Now(),
	//			}
	//		}
	//
	//		err := resGNodeTree.SynchronizeNodesInfo(&t, nodesMap)
	//		Expect(err).ShouldNot(HaveOccurred())
	//
	//		str, err := resGNodeTree.ToJsonIndent(t)
	//		Expect(err).ShouldNot(HaveOccurred())
	//		By(fmt.Sprintf("Tree to Json = %s", str))
	//
	//	})
	//})
})
