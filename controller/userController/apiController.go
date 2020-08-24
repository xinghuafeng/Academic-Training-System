package userController

import (
	"Academic-Training-System/util"
	"github.com/gin-gonic/gin"
)

func GetOwnInfo(ctx *gin.Context) {
	userName, err := ctx.Cookie(util.AuthUserCookieId)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": userName,
	})
}
