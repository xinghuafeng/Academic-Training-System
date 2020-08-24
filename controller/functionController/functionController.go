package functionController

import (
	"Academic-Training-System/service/functionService"
	"github.com/gin-gonic/gin"
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
