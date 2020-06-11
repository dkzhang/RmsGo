package middleware

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

/////////////////////////////
// Securing Authenticated Routes
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := webapi.TheInfras.TheUserTempDM.ValidateToken(c.Request)
		if err != nil {
			logMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"error": err,
			}).Error("get userID from gin.Context failed.")
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "token验证未通过",
			})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
