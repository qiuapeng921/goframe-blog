package api

import (
	"blog/app/controller"
	"blog/app/library/response"
	"blog/app/service/user"
)

type UserController struct {
	controller.Controller
}

func (c *UserController) Info() {
	response.JsonExit(c.Request, 0, "success", user.GetUserInfo(c.Request.Session))
}
