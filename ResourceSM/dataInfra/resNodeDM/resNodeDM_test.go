package resNodeDM_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("ResNodeDM", func() {
	BeforeEach(func() {

	})

	Describe("Insert ResNode", func() {
		It("Insert 256 ResNode", func() {
			num := 256
			resNodes := make([]resNode.Node, num)
			for i := 0; i < num; i++ {
				nodeID := int64(i + 1)
				resNodes[i] = resNode.Node{
					ID:            nodeID,
					Name:          fmt.Sprintf("Node%d", nodeID),
					Status:        rand.Intn(100),
					Description:   fmt.Sprintf("Description of Node%d", nodeID),
					ProjectID:     rand.Intn(100),
					AllocatedTime: time.Now().Add(time.Duration(rand.Intn(1000)) * time.Second),
				}
				err := cpuDM.Insert(resNodes[i])
				Expect(err).ShouldNot(HaveOccurred())
			}

			for i := 0; i < num; i++ {
				nodeID := int64(i + 1)
				node, err := cpuDM.QueryByID(nodeID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(node.ID).Should(Equal(resNodes[i].ID))
				Expect(node.Name).Should(Equal(resNodes[i].Name))
				Expect(node.Status).Should(Equal(resNodes[i].Status))
				Expect(node.Description).Should(Equal(resNodes[i].Description))
				Expect(node.ProjectID).Should(Equal(resNodes[i].ProjectID))
				Expect(node.AllocatedTime).Should(BeTemporally("~", resNodes[i].AllocatedTime, time.Second))
			}

			_, err := cpuDM.GetAllArray()
			Expect(err).ShouldNot(HaveOccurred())

			_, err = cpuDM.GetAllMap()
			Expect(err).ShouldNot(HaveOccurred())
		})

		//It("Update", func() {
		//	for j := 0; j < 100; j++ {
		//		nodeID := rand.Int63n(100) + 1 // [1, 100]
		//
		//		nodeUpdated := resNode.Node{
		//			ID:            nodeID,
		//			Name:          fmt.Sprintf("Node%d", nodeID),
		//			Status:        rand.Intn(100),
		//			Description:   fmt.Sprintf("Description of Node%d", nodeID),
		//			ProjectID:     rand.Intn(100),
		//			AllocatedTime: time.Now().Add(time.Duration(rand.Intn(1000)) * time.Second),
		//		}
		//
		//		// update
		//		err := cpuDM.Update(nodeUpdated)
		//		Expect(err).ShouldNot(HaveOccurred())
		//
		//		// check
		//		node, err := cpuDM.QueryByID(nodeID)
		//		Expect(err).ShouldNot(HaveOccurred())
		//		Expect(node.ID).Should(Equal(nodeUpdated.ID))
		//		Expect(node.Name).Should(Equal(nodeUpdated.Name))
		//		Expect(node.Status).Should(Equal(nodeUpdated.Status))
		//		Expect(node.Description).Should(Equal(nodeUpdated.Description))
		//		Expect(node.ProjectID).Should(Equal(nodeUpdated.ProjectID))
		//		Expect(node.AllocatedTime).Should(BeTemporally("~", nodeUpdated.AllocatedTime, time.Second))
		//	}
		//})

		It("Update nodes", func() {
			nodes := make([]resNode.Node, 100)
			for j := 0; j < 100; j++ {
				nodeID := j + 1 // [1, 100]

				nodes[j] = resNode.Node{
					ID:            int64(nodeID),
					Name:          fmt.Sprintf("Node%d", nodeID),
					Status:        rand.Intn(100),
					Description:   fmt.Sprintf("Description of Node%d", nodeID),
					ProjectID:     rand.Intn(100),
					AllocatedTime: time.Now().Add(time.Duration(rand.Intn(1000)) * time.Second),
				}

				err := cpuDM.UpdateNodes(nodes)
				Expect(err).ShouldNot(HaveOccurred())
			}

			_, err := cpuDM.GetAllMap()
			Expect(err).ShouldNot(HaveOccurred())
			//for _,pNode := range nodesMap{
			//	By(fmt.Sprintf("node=%v", *pNode))
			//}

			for j := 60; j < 100; j++ {
				nodeID := j + 1
				// check
				node, err := cpuDM.QueryByID(int64(nodeID))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(node.ID).Should(Equal(nodes[j].ID), fmt.Sprintf("nodeID=%d, node=%v", nodeID, node))
				Expect(node.Name).Should(Equal(nodes[j].Name))
				Expect(node.Status).Should(Equal(nodes[j].Status))
				Expect(node.Description).Should(Equal(nodes[j].Description))
				Expect(node.ProjectID).Should(Equal(nodes[j].ProjectID))
				Expect(node.AllocatedTime).Should(BeTemporally("~", nodes[j].AllocatedTime, time.Second))
			}
		})
	})
})
