package controller

import (
	"Academic-Training-System/service/systemService"
	"github.com/gin-gonic/gin"
)

func Baseinfo(ctx *gin.Context) {

	res, err := systemService.All(ctx)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": res,
	})
}
