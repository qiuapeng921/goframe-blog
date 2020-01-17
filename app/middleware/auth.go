package middleware

import (
	"blog/app/library/response"
	"blog/app/service/user"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	glog.Debug("2222222222222222")
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
