package middleware

import (
	"blog/app/consts"
	"blog/app/helpers/jwt"
	"blog/app/helpers/response"
	"github.com/gogf/gf/net/ghttp"
)

// 鉴权中间件，只有登录成功之后才能通过
func ApiAuth(r *ghttp.Request) {
	authorization := r.Request.Header.Get("Authorization")
	if authorization == "" {
		response.JsonExit(r, consts.ERROR, "您暂未登录")
	}
	user, err := jwt.ParseToken(authorization)
	if err != nil {
		response.JsonExit(r, consts.ERROR, err.Error())
	}
	if user.Category != "api"{
		response.JsonExit(r, consts.ERROR, "token错误")
	}
	r.SetParam("userId", user.Id)
	r.Middleware.Next()
}
