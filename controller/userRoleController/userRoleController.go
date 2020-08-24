package userRoleController

import (
	"Academic-Training-System/service/userRoleService"
	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	res, err := userRoleService.Find(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Save(ctx *gin.Context) {
	res, err := userRoleService.Save(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
