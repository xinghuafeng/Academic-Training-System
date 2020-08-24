package appController

import (
	"Academic-Training-System/service/appService"
	"github.com/gin-gonic/gin"
)

func All(ctx *gin.Context) {
	res, err := appService.All()
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Find(ctx *gin.Context) {
	res, err := appService.Find(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Save(ctx *gin.Context) {
	res, err := appService.Save(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}

func Delete(ctx *gin.Context) {
	res, err := appService.Delete(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
