package middleware

import (
	"blog/app/consts"
	"blog/app/helpers/jwt"
	"blog/app/helpers/response"
	"github.com/gogf/gf/net/ghttp"
)

// 鉴权中间件，只有登录成功之后才能通过
func AdminAuth(r *ghttp.Request) {
	authorization := r.Request.Header.Get("Authorization")
	if authorization == "" {
		response.JsonExit(r, consts.ERROR, "请登录后操作", nil)
	}
	user, err := jwt.ParseToken(authorization)
	if err != nil {
		response.JsonExit(r, consts.ERROR, err.Error(), nil)
	}
	if user.Category != "admin" {
		response.JsonExit(r, consts.ERROR, "token错误", nil)
	}
	r.SetParam("adminId", user.Id)
	r.Middleware.Next()
}
