package admin

import (
	"blog/app/controller"
	"blog/app/model/admins"
	"blog/app/request/admin_request"
	"blog/app/service/admin_service"
)

type AdminController struct {
	controller.Controller
}

func (c *AdminController) List() {
	var request admin_request.AdminRequest
	if err := c.Request.Parse(&request); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	data, count, err := admin_service.GetAdminList(request)
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	responseData := make(map[string]interface{})
	responseData["result"] = data
	responseData["count"] = count
	c.ResponseSuccess(c.Request, responseData)
}

func (c *AdminController) Info() {
	id := c.Request.GetInt("id")
	admin, err := admins.GetAdminById(id)
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, admin)
}

func (c *AdminController) Create() {
	var request admin_request.AdminCreateRequest
	if err := c.Request.Parse(&request); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	result,err := admin_service.CreateAdmin(request)
	if err != nil {
		c.ResponseFail(c.Request, "添加失败")
	}
	c.ResponseSuccess(c.Request,result)
}

func (c *AdminController) Update() {
	var request admin_request.AdminUpdateRequest
	if err := c.Request.Parse(&request); err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	result := admin_service.UpdateAdmin(request)
	if !result {
		c.ResponseFail(c.Request, "更新失败")
	}
	c.ResponseSuccess(c.Request,result)
}

func (c *AdminController) Delete() {
	status := c.Request.GetParam("status")
	id := c.Request.GetInt("id")
	result, err := admins.Model.Where("id = ?", id).Update("status = ?", status)
	if err != nil {
		c.ResponseFail(c.Request, "删除失败")
	}
	c.ResponseSuccess(c.Request, result)
}
