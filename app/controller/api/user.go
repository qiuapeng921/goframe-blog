package api

import (
	"blog/app/controller"
	"blog/app/library/response"
	"blog/app/service/user"
	"fmt"
)

type UserController struct {
	controller.Controller
}

func (c *UserController) Info() {
	response.JsonExit(c.Request, 200, "success", user.GetUserInfo(c.Request.Session))
}

func (c *UserController) GetInfo() {
	id := c.Request.GetInt("id")
	result := user.GetUserById(id)
	if result == nil {
		fmt.Println("111111111111111")
	}
	response.JsonExit(c.Request, 200, "success", result)
}
