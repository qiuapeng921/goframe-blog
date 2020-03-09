package admin

import (
	"blog/app/controller"
	"blog/app/request/admin_request"
	"blog/app/service/admin_service"
)

type AuthController struct {
	controller.Controller
}

func (c *AuthController) Login() {
	var data *admin_request.LoginRequest
	if err := c.Request.Parse(&data); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	result, err := admin_service.Login(data.Username, data.Password)
	if  err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, result)
}

func (c *AuthController) LogOut() {
	if err := admin_service.LogOut(c.Request.Session); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, c.Request.GetParam("admin_id"))
}
