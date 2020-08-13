package userController

import (
	"github.com/gin-gonic/gin"
	"sauth/service/userService"
)

func Delete(ctx *gin.Context) {
	res, err := userService.Delete(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Save(ctx *gin.Context) {
	res, err := userService.Save(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func All(ctx *gin.Context) {
	res, err := userService.All(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Find(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"err":  nil,
		"data": nil,
	})
}
