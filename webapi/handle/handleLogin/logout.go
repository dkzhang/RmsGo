package handleLogin

import (
	"github.com/dkzhang/RmsGo/webapi/infrastructure"
	"github.com/gin-gonic/gin"
)

func Logout(infra *infrastructure.Infrastructure, c *gin.Context) {
	userLoginInfo, err := infra.TheExtractor.Extract(c)
	if err != nil {
		return
	}

	infra.TheUserTempDM.DeleteToken(userLoginInfo.UserID)
}
