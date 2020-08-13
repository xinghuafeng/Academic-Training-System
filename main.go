package main

import (
	"github.com/gin-gonic/gin"
	"sauth/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)              // 发布模式
	engine := gin.Default()                   // web engine
	router.Router(engine)                     // 路由
	engine.Run(":9000")
}
