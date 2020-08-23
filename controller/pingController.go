package controller

import (
	"Academic-Training-System/util"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(util.StatusOK, gin.H{
		"code":    "S_AuthVerify_01",
		"message": "用户权限认证通过",
		"data":    nil,
		"err":     nil,
	})
}
