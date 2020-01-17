package api

import (
	"blog/app/controller"
	"blog/app/request/api"
	"blog/app/service/user"
)

type AuthController struct {
	controller.Controller
}

func (c *AuthController) Login() {
	var data *api.LoginRequest
	if err := c.Request.GetStruct(&data); err != nil {
		c.ResponseJson(c.Request, 1, err.Error())
	}
	if err := user.Login(data.Account, data.Password, c.Request.Session); err != nil {
		c.ResponseJson(c.Request, 1, err.Error())
	} else {
		c.ResponseJson(c.Request, 0, "ok")
	}
}

func (c *AuthController) Register() {
	// 表单验证
	var request *api.RegisterRequest
	if err := c.Request.GetStruct(&request); err != nil {
		c.ResponseJson(c.Request, 1, err.Error())
	}
	if err := user.Register(request); err != nil {
		c.ResponseJson(c.Request, 1, err.Error())
	}
	c.ResponseJson(c.Request, 0, "ok")
}

func (c *AuthController) LogOut() {
	if err := user.LogOut(c.Request.Session); err != nil {
		c.ResponseJson(c.Request, 1, err.Error())
	}
	c.ResponseJson(c.Request, 0, "ok")
}