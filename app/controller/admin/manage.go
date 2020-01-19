package admin

import (
	"blog/app/controller"
)

type ManageController struct {
	controller.Controller
}

func (c *ManageController) Info() {
	c.ResponseSuccess(c.Request, c.Session.Get("admin_info"))
}

func (c *ManageController) List() {

}
