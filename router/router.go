package router

import (
	"blog/app/controller"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	server := g.Server()
	server.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", new(controller.IndexController), "Index")
	})
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	server.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
	server.BindControllerMethod("/ws/{token}", new(controller.SocketController), "Socket")

	InitAdminRouter(server)
	InitApiRouter(server)
}
