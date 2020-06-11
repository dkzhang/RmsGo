package middleware

import (
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

//engine := gin.Default() //在这行后新增
//engine.Use(middleware.LoggerToFile())

// 日志记录到文件
func LoggerGinToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logMap.Log(logMap.GIN).WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
