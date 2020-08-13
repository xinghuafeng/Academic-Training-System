package controller

import (
	"github.com/gin-gonic/gin"
	"sauth/util"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(util.StatusOK, gin.H{
		"code":    "S_AuthVerify_01",
		"message": "用户权限认证通过",
		"data":    nil,
		"err":     nil,
	})
}
