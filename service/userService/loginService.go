package userService

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sauth/util"
	"sauth/model/userModel"
	"net/http"
)

/*
	登录验证
	@return count
	"0": 验证失败  |  "1": 验证成功
*/
func LoginVerify(ctx *gin.Context) (string, error) {
	tenantId := ctx.PostForm("tenantId")
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")
	password = util.Md5Encrypt(password)
	fmt.Println(password)
	res, err := userModel.FindCountByUserNameAndPasswordAndTenantId(userName, password, tenantId)
	if nil != err {
		return "", err
	}
	count := string(res[0]["count"])
	if "1" == count { // 设置 cookie
		cookie := http.Cookie{
			Name:     util.AuthUserCookieId,
			Value:    userName,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			// MaxAge=0 means no 'Max-Age' attribute specified.(会话级别 cookie)
			// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
			// MaxAge>0 means Max-Age attribute present and given in seconds
			MaxAge: 0,
		}
		http.SetCookie(ctx.Writer, &cookie)
	}
	return count, nil
}

func Logout(ctx *gin.Context) error {
	userName, err := ctx.Cookie(util.AuthUserCookieId)
	cookie := http.Cookie{
		Name:     util.AuthUserCookieId,
		Value:    userName,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		// MaxAge=0 means no 'Max-Age' attribute specified.(会话级别 cookie)
		// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
		// MaxAge>0 means Max-Age attribute present and given in seconds
		MaxAge: -1,
	}
	http.SetCookie(ctx.Writer, &cookie)
	return err
}

/*
	权限验证
*/
func AuthVerify(ctx *gin.Context) (string, error) {
	token, err := ctx.Cookie(util.AuthUserCookieId)
	res, _ := userModel.FindCountByUserName(token) // 数据库验证，后面考虑用缓存验证
	count := string(res[0]["count"])
	return count, err
}
