package orgController

import (
	"Academic-Training-System/service/orgService"
	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	res, err := orgService.Find(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Save(ctx *gin.Context) {
	res, err := orgService.Save(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Delete(ctx *gin.Context) {
	res, err := orgService.Delete(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
