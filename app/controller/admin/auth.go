package admin

import (
	"blog/app/controller"
	"blog/app/model/admins"
	"blog/app/request/admin_request"
	"blog/app/service/admin_service"
	"log"
)

type AuthController struct {
	controller.Controller
}

func (c *AuthController) Login() {
	var data *admin_request.LoginRequest
	if err := c.Request.Parse(&data); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	log.Println(c.Request.GetClientIp())
	result, err := admin_service.Login(data.Username, data.Password, c.Request.GetClientIp())
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, result)
}

func (c *AuthController) LogOut() {
	adminId := c.Request.GetInt64("adminId")
	result, err := admin_service.LogOut(adminId)
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, result)
}

func (c *AuthController) Info() {
	adminId := c.Request.GetInt("adminId")
	result, err := admins.GetAdminById(adminId)
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	response := make(map[string]interface{})
	response["name"] = result.Username
	response["avatar"] = "https://wpimg.wallstcn.com/69a1c46c-eb1c-4b46-8bd4-e9e686ef5251.png"
	response["roles"] = [...]string{"admin", "edits"}
	c.ResponseSuccess(c.Request, response)
}
