package api

import (
	"blog/app/controller"
	"blog/app/library/client"
	"blog/app/library/jwt"
	"blog/app/request/api_request"
	"blog/app/service/user_service"
)

type AuthController struct {
	controller.Controller
}

func (c *AuthController) Login() {
	var data *api_request.LoginRequest
	if err := c.Request.Parse(&data); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	if result, err := user_service.Login(data.Account, data.Password); err != nil {
		c.ResponseFail(c.Request, err.Error())
	} else {
		accessToken, _ := jwt.CreateToken(result)
		_, _ = client.HSet("user_info", result.Id, accessToken)
		c.ResponseSuccess(c.Request, accessToken)
	}
}

func (c *AuthController) Register() {
	// 表单验证
	var request *api_request.RegisterRequest
	if err := c.Request.Parse(&request); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	if err := user_service.Register(request); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, "ok")
}

func (c *AuthController) LogOut() {
	if err := user_service.LogOut(c.Request.Session); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, "ok")
}
