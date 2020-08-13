package roleFunctionController

import (
	"github.com/gin-gonic/gin"
	"sauth/service/roleFunctionService"
)

func Save(ctx *gin.Context) {
	res, err := roleFunctionService.Save(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Find(ctx *gin.Context) {
	res, err := roleFunctionService.Find(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
