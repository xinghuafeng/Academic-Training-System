package userController

import (
	"github.com/gin-gonic/gin"
	"sauth/util"
)

func GetOwnInfo(ctx *gin.Context) {
	userName, err := ctx.Cookie(util.AuthUserCookieId)
	ctx.JSON(200, gin.H{
		"err":  err,
		"data": userName,
	})
}
