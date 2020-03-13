package admin

import (
	"blog/app/controller"
	"blog/app/model/roles"
	"github.com/gogf/gf/util/gconv"
)

type RoleController struct {
	controller.Controller
}

func (c *RoleController) List() {
	result, err := roles.Model.Limit(0, 10).All()
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, result)
}

func (c *RoleController) Info() {
	id := c.Request.Get("id")
	result, err := roles.FindOne("id", id)
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	c.ResponseSuccess(c.Request, result)
}

func (c *RoleController) Create() {
	params := c.Request.GetMap()
	roleName := params["role_name"]
	roleDesc := params["role_desc"]
	var roleEntity roles.Entity
	roleEntity.RoleName = gconv.String(roleName)
	roleEntity.RoleDesc = gconv.String(roleDesc)
	result, err := roles.Insert(roleEntity)
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	num, _ := result.LastInsertId()
	c.ResponseSuccess(c.Request, num)
}

func (c *RoleController) Update() {
	params := c.Request.GetMap()
	id := params["id"]
	roleName := params["role_name"]
	roleDesc := params["role_desc"]
	_, err := roles.FindOne("id", id)
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	result, err2 := roles.Model.Where("id", id).Data("role_name", roleName, "role_desc", roleDesc).Update()
	if err2 != nil {
		c.ResponseFail(c.Request, err2.Error())
	}
	num, _ := result.RowsAffected()
	c.ResponseSuccess(c.Request, num)
}

func (c *RoleController) Delete() {
	id := c.Request.Get("id")
	result, err := roles.Update("status=1", "id", id)
	if err != nil {
		c.ResponseFail(c.Request, err.Error())
	}
	response, _ := result.RowsAffected()
	c.ResponseSuccess(c.Request, response)
}
