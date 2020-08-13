package orgController

import (
	"github.com/gin-gonic/gin"
	"sauth/service/orgService"
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
