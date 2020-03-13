package router

import (
	"blog/app/controller/admin"
	"blog/app/middleware"
	"github.com/gogf/gf/net/ghttp"
)

func InitAdminRouter(server *ghttp.Server) {
	// TODO ----------------------后台模块接口--------------------------------
	server.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.POST("/auth/login", new(admin.AuthController), "Login")

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.AdminAuth)
			group.POST("/auth/logout", new(admin.AuthController), "LogOut")
			// 管理员
			group.Group("/manage", func(group *ghttp.RouterGroup) {
				manage := new(admin.AdminController)
				group.POST("/list", manage, "List")
				group.POST("/info/{id}", manage, "Info")
				group.POST("/create", manage, "Create")
				group.POST("/update/{id}", manage, "Update")
				group.POST("/delete/{id}", manage, "Delete")
			})
			// 角色
			group.Group("/role", func(group *ghttp.RouterGroup) {
				role := new(admin.RoleController)
				group.POST("/list", role, "List")
				group.POST("/info/{id}", role, "Info")
				group.POST("/create", role, "Create")
				group.POST("/update/{id}", role, "Update")
				group.POST("/delete/{id}", role, "Delete")
			})
			// 权限

		})

	})
}
