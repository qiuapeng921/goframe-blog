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
		c.ResponseFail(c.Request, err.Error())
	}
	if res, err := user.Login(data.Account, data.Password); err != nil {
		c.ResponseFail(c.Request, err.Error())
	} else {
		c.ResponseSuccess(c.Request, "登录成功", res.Id)
	}
}

func (c *AuthController) Register() {
	// 表单验证
	var request *api.RegisterRequest
	if err := c.Request.GetStruct(&request); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	if err := user.Register(request); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, "ok")
}

func (c *AuthController) LogOut() {
	if err := user.LogOut(c.Request.Session); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, "ok")
}
