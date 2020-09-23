package handleProjectRes

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/authority/authApplication"
	"github.com/dkzhang/RmsGo/webapi/authority/authProject"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectResDM"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type HandleProjectRes struct {
	theProjectDM    projectDM.ProjectDM
	theProjectResDM projectResDM.ProjectResDM
	theExtractor    extractLoginUserInfo.Extractor
	theLogMap       logMap.LogMap
}

func NewHandleProjectRes(prdm projectResDM.ProjectResDM, pdm projectDM.ProjectDM,
	ext extractLoginUserInfo.Extractor, lm logMap.LogMap) HandleProjectRes {
	return HandleProjectRes{
		theProjectDM:    pdm,
		theProjectResDM: prdm,
		theExtractor:    ext,
		theLogMap:       lm,
	}
}

// 获取树
func (h HandleProjectRes) QueryCpuTreeOccupied(c *gin.Context) {
	h.queryTree(c, h.theProjectResDM.QueryCpuTreeOccupied)
}

func (h HandleProjectRes) QueryCpuTreeAvailable(c *gin.Context) {
	h.queryTree(c, h.theProjectResDM.QueryCpuTreeAvailable)
}
func (h HandleProjectRes) QueryGpuTreeOccupied(c *gin.Context) {
	h.queryTree(c, h.theProjectResDM.QueryGpuTreeOccupied)
}
func (h HandleProjectRes) QueryGpuTreeAvailable(c *gin.Context) {
	h.queryTree(c, h.theProjectResDM.QueryGpuTreeAvailable)
}

///////////////////////////////////////////////////////////
// 分配资源
func (h HandleProjectRes) SchedulingCpu(c *gin.Context) {
	h.schedulingCGpu(c, h.theProjectResDM.SchedulingCpu)
}

func (h HandleProjectRes) SchedulingGpu(c *gin.Context) {
	h.schedulingCGpu(c, h.theProjectResDM.SchedulingGpu)
}

func (h HandleProjectRes) schedulingCGpu(c *gin.Context,
	schFunc func(projectID int, nodesAfter []int64, ctrlUserInfo user.UserInfo) (err error)) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	pi, err := h.extractAccessedProject(c)
	if err != nil {
		return
	}

	// Only RoleController can actively schedule resources
	if userLoginInfo.Role != user.RoleController {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "无权访问",
		})
		return
	}

	var schCGpuNodes SchCGpuNodes
	err = c.BindJSON(&schCGpuNodes)
	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&SchCGpuNodes) error.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无权解析CGpuNodes结构",
		})
		return
	}

	err = schFunc(pi.ProjectID, schCGpuNodes.NodesAfter, userLoginInfo)
	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"projectInfo": pi,
			"error":       err,
		}).Error("ProjectResDM.SchedulingCGpu error.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "CPU/GPU资源分配操作失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "CPU/GPU资源分配成功",
	})
	return
}

func (h HandleProjectRes) SchedulingStorage(c *gin.Context) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	pi, err := h.extractAccessedProject(c)
	if err != nil {
		return
	}

	// Only RoleController can actively schedule resources
	if userLoginInfo.Role != user.RoleController {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "无权访问",
		})
		return
	}

	var schStorage SchStorage
	err = c.BindJSON(&schStorage)
	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"error":       err,
		}).Error("c.BindJSON(&SchStorage) error.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无权解析SchStorage结构",
		})
		return
	}

	err = h.theProjectResDM.SchedulingStorage(pi.ProjectID,
		schStorage.StorageSizeAfter, schStorage.StorageAllocInfoAfter, userLoginInfo)
	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"userLoginID": userLoginInfo.UserID,
			"projectInfo": pi,
			"error":       err,
		}).Error("ProjectResDM.SchedulingStorage error.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "存储资源分配操作失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "存储资源分配成功",
	})
	return
}

type SchCGpuNodes struct {
	NodesAfter []int64 `json:"nodes_after"`
}

type SchStorage struct {
	StorageSizeAfter      int    `json:"nodes_after"`
	StorageAllocInfoAfter string `json:"storage_alloc_info_after"`
}

///////////////////////////////////////////////////////////
func (h HandleProjectRes) queryTree(c *gin.Context, funcQuery func(projectID int) (jsonTree string, err error)) {
	userLoginInfo, err := h.theExtractor.Extract(c)
	if err != nil {
		return
	}

	pi, err := h.extractAccessedProject(c)
	if err != nil {
		return
	}

	permission := authProject.AuthorityCheck(h.theLogMap, userLoginInfo, pi, authApplication.OPS_RETRIEVE)
	if permission == true {
		treeJson, err := funcQuery(pi.ProjectID)
		if err != nil {
			h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"err":         err,
				"projectInfo": pi,
			}).Error("ProjectResDM Query Tree(using pid from gin.Context) failed.")

			c.JSON(http.StatusNotFound, gin.H{
				"msg": "查询失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tree": treeJson,
			"msg":  "查询成功",
		})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "无权访问",
		})
		return
	}
}

func (h HandleProjectRes) extractAccessedProject(c *gin.Context) (pi project.Info, err error) {
	idStr := c.Param("id")
	pid, err := strconv.Atoi(idStr)

	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"idStr": idStr,
			"error": err,
		}).Error("get Project ID from gin.Context failed.")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的Request：拟操作的Project ID无效",
		})
		return project.Info{},
			fmt.Errorf("get Project ID from gin.Context failed: %v", err)
	}

	pi, err = h.theProjectDM.QueryByID(pid)
	if err != nil {
		h.theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"pid": pid,
		}).Error("ProjectDM.QueryByID (using pid from gin.Context) failed.")

		c.JSON(http.StatusNotFound, gin.H{
			"msg": "无法找到该项目",
		})
		return project.Info{},
			fmt.Errorf("ProjectDM.QueryByID (using pid from gin.Context) error: %v", err)
	}
	return pi, nil
}
