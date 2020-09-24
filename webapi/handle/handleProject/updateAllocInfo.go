package handleProject

import (
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h HandleProject) UpdateAllocInfo(c *gin.Context) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	pi, err := h.extractAccessedProject(c)
	if err != nil {
		return
	}

	// only Controller can update alloc info
	if userLoginInfo.Role != user.RoleController {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "无权访问",
		})
		return
	}

	var allocInfo project.AllocInfo
	err = c.BindJSON(&allocInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无法解析project.AllocInfo 结构",
		})
		return
	}
	allocInfo.ProjectID = pi.ProjectID
	err = h.theProjectDM.UpdateAllocInfo(allocInfo)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "更新分配信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "更新分配信息成功",
	})
}
