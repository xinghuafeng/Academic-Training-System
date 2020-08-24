package main

import (
	"Academic-Training-System/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 发布模式
	engine := gin.Default()      // web engine
	router.Router(engine)        // 路由
	engine.Run(":9000")
}
