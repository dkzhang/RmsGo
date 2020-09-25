package resScheduling_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResSch", func() {
	BeforeEach(func() {

	})

	Describe("Query Tree", func() {
		It("SchedulingCPU", func() {
			projectID := 1
			nodes := make([]int64, 100)
			for i := 0; i < 100; i++ {
				nodes[i] = int64(i + 1)
			}
			ctrlID := 1
			ctrlCN := "zhang"

			isFirstAlloc, err := theResScheduling.SchedulingCPU(projectID, nodes, ctrlID, ctrlCN)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(isFirstAlloc).Should(Equal(true))

			cpuAllocRecords, err := cadm.QueryAll()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(cpuAllocRecords)).Should(Equal(1))

			cpuNodesMap, err := cndm.GetAllMap()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(cpuNodesMap)).Should(Equal(256))
			for i := int64(0); i < 100; i++ {
				Expect(cpuNodesMap[i+1].ProjectID).Should(Equal(projectID))
			}

			cpuTree, selected, err := theResScheduling.QueryCpuTreeIdleAndAllocated(projectID+1, 1)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("cpuTree.NodesNum=%d", cpuTree.NodesNum))
			By(fmt.Sprintf("selected=%v", selected))

			cpuTree2, err := ctdm.QueryTree(1, func(node resNode.Node) bool {
				return node.ProjectID == projectID
			})
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("cpuTree.NodesNum=%d", cpuTree2.NodesNum))
		})

		It("SchedulingGPU", func() {
			projectID := 1
			nodes := make([]int64, 33)
			for i := 0; i < 33; i++ {
				nodes[i] = int64(i + 1)
			}
			ctrlID := 1
			ctrlCN := "zhang"

			_, err := theResScheduling.SchedulingGPU(projectID, nodes, ctrlID, ctrlCN)
			Expect(err).ShouldNot(HaveOccurred())

			gpuAllocRecords, err := gadm.QueryAll()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(gpuAllocRecords)).Should(Equal(1))

			gpuNodesMap, err := gndm.GetAllMap()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(gpuNodesMap)).Should(Equal(66))
			for i := int64(0); i < 33; i++ {
				Expect(gpuNodesMap[i+1].ProjectID).Should(Equal(projectID))
			}

			gpuTree, selected, err := theResScheduling.QueryGpuTreeIdleAndAllocated(projectID+1, 1)
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("gpuTree.NodesNum=%d", gpuTree.NodesNum))
			By(fmt.Sprintf("selected=%v", selected))

			gpuTree2, err := gtdm.QueryTree(1, func(node resNode.Node) bool {
				return node.ProjectID == projectID
			})
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("gpuTree.NodesNum=%d", gpuTree2.NodesNum))
		})
	})
})
