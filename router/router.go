package router

import (
	"blog/app/controller"
	"blog/app/controller/admin"
	"blog/app/controller/api"
	"blog/app/middleware"
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

	// TODO ----------------------前台模块接口--------------------------------
	server.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Group("/auth", func(group *ghttp.RouterGroup) {
			group.POST("/login", new(api.AuthController), "Login")
			group.POST("/register", new(api.AuthController), "Register")
		})
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.POST("/auth/logout", new(api.AuthController), "LogOut")
			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.ApiAuth)
				group.POST("/info", new(api.UserController), "Info")
				group.POST("/getInfo/{id}", new(api.UserController), "GetInfo")
			})
		})
	})

	// TODO ----------------------后台模块接口--------------------------------
	server.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Group("/auth", func(group *ghttp.RouterGroup) {
			group.POST("/login", new(admin.AuthController), "Login")
		})
		group.POST("/auth/logout", new(admin.AuthController), "LogOut")
		group.Group("/manage", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.AdminAuth)
			group.POST("/info", new(admin.ManageController), "Info")
			group.POST("/list", new(admin.ManageController), "List")
		})
	})

}
