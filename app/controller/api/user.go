package api

import (
	"blog/app/controller"
	"blog/app/service/user_service"
	"github.com/gogf/gf/util/gconv"
)

type UserController struct {
	controller.Controller
}

func (c *UserController) Info() {
	id := c.Request.GetParam("userId")
	userInfo, err := user_service.GetUserById(gconv.Uint(id))
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, userInfo)
}

func (c *UserController) GetInfo() {
	id := c.Request.GetInt("id")
	userInfo, err := user_service.GetUserById(gconv.Uint(id))
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, userInfo)
}
