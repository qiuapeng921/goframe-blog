package router

import (
	"blog/app/controller"
	"blog/app/controller/api"
	"blog/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	server := g.Server()
	server.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", controller.Index)
	})

	server.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.Group("/auth", func(group *ghttp.RouterGroup) {
			group.ALL("/login", new(api.AuthController), "Login")
			group.ALL("/register", new(api.AuthController), "Register")
		})
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/auth/logout", new(api.AuthController), "LogOut")
			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.Auth)
				group.ALL("/info", new(api.UserController), "Info")
			})
		})
	})
}
