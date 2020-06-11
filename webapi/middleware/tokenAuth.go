package middleware

import (
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

/////////////////////////////
// Securing Authenticated Routes
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := webapi.TheInfras.TheUserTempDM.ValidateToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
