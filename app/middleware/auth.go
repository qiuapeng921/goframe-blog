package middleware

import (
	"blog/app/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	if r.Request.Header.Get("Authorization") != "" {
		r.Middleware.Next()
	} else {
		_ = r.Response.WriteJson(response.JsonResponse{
			Code:    1,
			Message: "您暂未登录",
		})
		r.Exit()
	}
}
