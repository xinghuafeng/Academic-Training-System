package functionController

import (
	"github.com/gin-gonic/gin"
	"sauth/service/functionService"
)

func All(ctx *gin.Context) {
	res, err := functionService.All()
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Find(ctx *gin.Context) {
	res, err := functionService.Find(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
