package api

import (
	"blog/app/controller"
	"blog/app/request/api"
	"blog/app/service/user"
	"blog/app/utility"
)

type AuthController struct {
	controller.Controller
}

type LoginRequest struct {
	api.LoginRequest
}

func (c *AuthController) Login() {
	var data *LoginRequest
	if err := c.Request.GetStruct(&data); err != nil {
		response.JsonExit(c.Request, 1, err.Error())
	}
	if err := user.Login(data.Account, data.Password, c.Request.Session); err != nil {
		response.JsonExit(c.Request, 1, err.Error())
	} else {
		response.JsonExit(c.Request, 0, "ok")
	}
}

func (c *AuthController) Register() {
	// 表单验证
	var request *api.RegisterRequest
	if err := c.Request.GetStruct(&request); err != nil {
		response.JsonExit(c.Request, 1, err.Error())
	}
	if err := user.Register(request); err != nil {
		response.JsonExit(c.Request, 1, err.Error())
	}
	response.JsonExit(c.Request, 0, "ok")
}

func (c *AuthController) LogOut() {
	if err := user.LogOut(c.Request.Session); err != nil {
		response.JsonExit(c.Request, 1, "")
	}
	response.JsonExit(c.Request, 0, "ok")
}
