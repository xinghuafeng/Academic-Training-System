package appController

import (
	"github.com/gin-gonic/gin"
	"sauth/service/appService"
)

func Options(ctx *gin.Context) {
	res, err := appService.FindOptions()
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
