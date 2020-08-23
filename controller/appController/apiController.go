package appController

import (
	"Academic-Training-System/service/appService"
	"github.com/gin-gonic/gin"
)

func Options(ctx *gin.Context) {
	res, err := appService.FindOptions()
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
