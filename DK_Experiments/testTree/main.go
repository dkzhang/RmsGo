package main

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNode"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNodeTree"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var t resGNodeTree.Tree
var nodesMap map[int64]resNode.Node

func init() {
	rootGNode := resGNode.ResGNode{
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

	t = resGNodeTree.Tree{
		Root:     rootGNode,
		NodesNum: 0,
	}

	resGNodeTree.Count(&t)

	///////////////////////////////////////////////////////

	nodesMap = make(map[int64]resNode.Node, 256)
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
}

func main() {
	r := gin.Default()

	webAPIv1 := r.Group("/webapi")
	{
		hTree := webAPIv1.Group("/Tree")
		{
			hTree.GET("/All", RetrieveAll)
			hTree.GET("/ProRes/:id", RetrieveProRes)
			hTree.GET("/ProResD/:id", RetrieveProResD)
			hTree.POST("/ProRes/:id", SubmitProRes)
		}
	}

	r.Run(":8084")
}

func RetrieveAll(c *gin.Context) {
	jsonTree, err := resGNodeTree.ToJsonForVue(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "tree转json失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tree":     jsonTree,
		"nodesNum": resGNodeTree.CountRO(&t),
		"unfolded": []int64{1e4, 11e4, 115e4},
		"selected": []int64{10, 11, 12, 14, 15},
	})
}

func RetrieveProRes(c *gin.Context) {
	idStr := c.Param("id")
	projectID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("projectID格式错误%s", idStr),
		})
		return
	}
	nt, err := resGNodeTree.Filtrate(&t, nodesMap, func(node resNode.Node) bool {
		return node.ProjectID == 0 || node.ProjectID == projectID
	})

	jsonTree, err := resGNodeTree.ToJsonForVue(*nt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "tree转json失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tree":     jsonTree,
		"nodesNum": resGNodeTree.CountRO(nt),
		"unfolded": []int64{1e4, 11e4},
		"selected": []int64{},
	})
	return
}

func RetrieveProResD(c *gin.Context) {
	idStr := c.Param("id")
	projectID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("projectID格式错误%s", idStr),
		})
		return
	}
	nt, err := resGNodeTree.FiltrateMark(&t, nodesMap, func(node resNode.Node) bool {
		return node.ProjectID == 0 || node.ProjectID == projectID
	})

	jsonTree, err := resGNodeTree.ToJsonForVue(*nt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "tree转json失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tree":     jsonTree,
		"nodesNum": resGNodeTree.CountRO(nt),
		"unfolded": []int64{1e4, 11e4},
		"selected": []int64{},
	})
	return
}

type SubmitSelect struct {
	Selected []int64 `json:"selected"`
}

func SubmitProRes(c *gin.Context) {
	idStr := c.Param("id")
	projectID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("projectID格式错误%s", idStr),
		})
		return
	}

	ss := SubmitSelect{}
	err = c.BindJSON(&ss)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("BindJSON error: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("receive (projectID=%d) selected array: %v", projectID, ss.Selected),
	})
	return
}
