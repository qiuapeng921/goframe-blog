package router

import (
	"blog/app/controller/admin"
	"blog/app/middleware"
	"github.com/gogf/gf/net/ghttp"
)

func InitAdminRouter(server *ghttp.Server){
	// TODO ----------------------后台模块接口--------------------------------
	server.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.POST("/auth/login", new(admin.AuthController), "Login")

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.AdminAuth)
			group.POST("/auth/logout", new(admin.AuthController), "LogOut")
			// 管理员
			group.Group("/manage", func(group *ghttp.RouterGroup) {
				group.POST("/info", new(admin.ManageController), "Info")
				group.POST("/list", new(admin.ManageController), "List")
			})
			// 角色
			group.Group("/role", func(group *ghttp.RouterGroup) {
				group.POST("/list", new(admin.RoleController), "List")
				group.POST("/info/{id}", new(admin.RoleController), "Info")
				group.POST("/save", new(admin.RoleController), "Save")
				group.POST("/update/{id}", new(admin.RoleController), "Update")
				group.POST("/delete/{id}", new(admin.RoleController), "Delete")
			})
			// 权限

		})

	})
}
