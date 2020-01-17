package api

import (
	"blog/app/controller"
	"blog/app/service/user"
	"blog/app/utility"
)

type UserController struct {
	controller.Controller
}

func (c *UserController) Info() {
	response.JsonExit(c.Request, 0, "success", user.GetUserInfo(c.Request.Session))
}
