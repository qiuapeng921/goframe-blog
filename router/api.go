package router

import (
	"blog/app/controller/api"
	"blog/app/middleware"
	"github.com/gogf/gf/net/ghttp"
)

func InitApiRouter(server *ghttp.Server){
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
}
