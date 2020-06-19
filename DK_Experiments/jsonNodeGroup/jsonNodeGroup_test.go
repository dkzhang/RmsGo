package jsonNodeGroup_test

import (
	"fmt"
	"github.com/dkzhang/RmsGo/DK_Experiments/jsonNodeGroup"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsonNodeGroup", func() {
	var (
		ng1 jsonNodeGroup.NodeGroup
	)

	BeforeEach(func() {
		ng1 = jsonNodeGroup.NodeGroup{
			Name:       "a1",
			IsSelected: false,
			SubGroup: []jsonNodeGroup.NodeGroup{{
				Name:       "b1",
				IsSelected: false,
				SubGroup: []jsonNodeGroup.NodeGroup{
					{
						Name:       "c1",
						IsSelected: false,
						SubGroup:   nil,
					},
					{
						Name:       "c2",
						IsSelected: false,
						SubGroup:   nil,
					},
				},
			},
				{
					Name:       "b2",
					IsSelected: false,
					SubGroup: []jsonNodeGroup.NodeGroup{
						{
							Name:       "c3",
							IsSelected: true,
							SubGroup:   nil,
						},
					},
				},
			},
		}
	})

	Describe("test NodeGroup marshal", func() {
		It("should no error", func() {
			str, err := ng1.ToJson()
			Expect(err).ShouldNot(HaveOccurred())
			By(fmt.Sprintf("json NodeGroup = %s", str))
		})
	})
})
