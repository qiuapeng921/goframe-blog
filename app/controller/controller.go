package controller

import (
	"blog/app/consts"
	"blog/app/helpers/response"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	gmvc.Controller
}

func (c *Controller) ResponseJson(request *ghttp.Request, err int, msg string, data ...interface{}) {
	response.JsonExit(request, err, msg, data)
}

func (c *Controller) ResponseSuccess(request *ghttp.Request, data ...interface{}) {
	c.ResponseJson(request, consts.SUCCESS, "success", data...)
}

func (c *Controller) ResponseFail(request *ghttp.Request, msg string) {
	c.ResponseJson(request, consts.ERROR, msg, nil)
}
