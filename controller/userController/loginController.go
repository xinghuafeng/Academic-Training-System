package userController

import (
	"github.com/gin-gonic/gin"
	"sauth/service/userService"
)

const (
	LoginURI     = "/auth/api/v1/user/login"
	LoginAddress = "http://127.0.0.1:8889/sauth/index.html"
)

/*
	登录
*/
func Login(ctx *gin.Context) {
	res, err := userService.LoginVerify(ctx)
	ctx.JSON(200, gin.H{
		"data": res,
		"err":  err,
	})
}

/*
	登出
*/
func Logout(ctx *gin.Context) {
	err := userService.Logout(ctx)
	ctx.JSON(200, gin.H{
		"data": nil,
		"err":  err,
	})
}

/*
	cookie 权限认证
*/
func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if LoginURI == ctx.Request.URL.String() { // 进行登录时不进行cookie认证
			ctx.Next()
			return
		}
		count, err := userService.AuthVerify(ctx) // 访问权限验证
		if "0" == count { // 认证失败
			ctx.JSON(200, gin.H{
				"code":    "F_AuthVerify_01",
				"message": "用户权限认证失败",
				"data":    LoginAddress,
				"err":     err,
			})
			ctx.Abort() // 终止计划
			return
		}
		ctx.Next()
	}
}
