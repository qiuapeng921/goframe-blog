package middleware

import (
	"blog/app/service/user"
	response "blog/app/utility"
	"github.com/gogf/gf/net/ghttp"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	if user.IsSignedIn(r.Session) {
		r.Middleware.Next()
	} else {
		_ = r.Response.WriteJson(response.JsonResponse{
			Code:    1,
			Message: "您暂未登录",
		})
		r.Exit()
	}
}
