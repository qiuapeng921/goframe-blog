package middleware

import (
	"blog/app/consts"
	"blog/app/library/response"
	"blog/app/service/admin"
	"github.com/gogf/gf/net/ghttp"
)

// 鉴权中间件，只有登录成功之后才能通过
func AdminAuth(r *ghttp.Request) {
	if admin.IsSignedIn(r.Session) {
		r.Middleware.Next()
	} else {
		response.JsonExit(r, consts.ERROR, "您暂未登录")
	}
}