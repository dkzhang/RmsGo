package resNodeDB_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("ResNodeDB", func() {
	BeforeEach(func() {

	})

	Describe("Insert ResNode", func() {
		It("Insert 256 ResNode", func() {
			num := 256
			resNodes := make([]resNode.Node, num)
			for i := 0; i < num; i++ {
				nodeID := i + 1
				resNodes[i] = resNode.Node{
					ID:            nodeID,
					Name:          fmt.Sprintf("Node%d", nodeID),
					Status:        rand.Intn(100),
					Description:   fmt.Sprintf("Description of Node%d", nodeID),
					ProjectID:     rand.Intn(100),
					AllocatedTime: time.Now().Add(time.Duration(rand.Intn(1000)) * time.Second),
				}
				err := rndb.Insert(resNodes[i])
				Expect(err).ShouldNot(HaveOccurred())
			}

			for i := 0; i < num; i++ {
				nodeID := i + 1
				node, err := rndb.QueryByID(nodeID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(node.ID).Should(Equal(resNodes[i].ID))
				Expect(node.Name).Should(Equal(resNodes[i].Name))
				Expect(node.Status).Should(Equal(resNodes[i].Status))
				Expect(node.Description).Should(Equal(resNodes[i].Description))
				Expect(node.ProjectID).Should(Equal(resNodes[i].ProjectID))
				Expect(node.AllocatedTime).Should(BeTemporally("~", resNodes[i].AllocatedTime, time.Second))
			}

			_, err := rndb.QueryAll()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Update", func() {
			for j := 0; j < 100; j++ {
				nodeID := rand.Intn(100) + 1 // [1, 100]

				nodeUpdated := resNode.Node{
					ID:            nodeID,
					Name:          fmt.Sprintf("Node%d", nodeID),
					Status:        rand.Intn(100),
					Description:   fmt.Sprintf("Description of Node%d", nodeID),
					ProjectID:     rand.Intn(100),
					AllocatedTime: time.Now().Add(time.Duration(rand.Intn(1000)) * time.Second),
				}

				// update
				err := rndb.Update(nodeUpdated)
				Expect(err).ShouldNot(HaveOccurred())

				// check
				node, err := rndb.QueryByID(nodeID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(node.ID).Should(Equal(nodeUpdated.ID))
				Expect(node.Name).Should(Equal(nodeUpdated.Name))
				Expect(node.Status).Should(Equal(nodeUpdated.Status))
				Expect(node.Description).Should(Equal(nodeUpdated.Description))
				Expect(node.ProjectID).Should(Equal(nodeUpdated.ProjectID))
				Expect(node.AllocatedTime).Should(BeTemporally("~", nodeUpdated.AllocatedTime, time.Second))
			}
		})
	})
})
