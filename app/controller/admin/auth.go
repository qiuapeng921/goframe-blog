package admin

import (
	"blog/app/controller"
	adminRequest "blog/app/request/admin"
	"blog/app/service/admin"
)

type AuthController struct {
	controller.Controller
}

func (c *AuthController) Login() {
	var data *adminRequest.LoginRequest
	if err := c.Request.Parse(&data); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	result, err := admin.Login(data.Username, data.Password,c.Session)
	if  err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, result)
}

func (c *AuthController) LogOut() {
	if err := admin.LogOut(c.Request.Session); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, c.Request.GetParam("admin_id"))
}
