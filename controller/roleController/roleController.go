package roleController

import (
	"Academic-Training-System/service/roleService"
	"github.com/gin-gonic/gin"
)

func All(ctx *gin.Context) {
	res, err := roleService.All()
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Find(ctx *gin.Context) {
	res, err := roleService.Find(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Save(ctx *gin.Context) {
	res, err := roleService.Save(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Delete(ctx *gin.Context) {
	res, err := roleService.Delete(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
