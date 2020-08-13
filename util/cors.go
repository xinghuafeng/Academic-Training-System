package util

import (
	"github.com/gin-gonic/gin"
)

const (
	AuthWebOrigin = "http://127.0.0.1:9000"
)

/*
	允许跨域请求
	支持cookie
*/
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin") // 获取 origin
		/*
			检查 origin 是否被允许跨域请求
			可通过数据库或缓存进行验证
		*/
		if AuthWebOrigin == origin { // demo: 允许 auth web 端跨域请求
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}
