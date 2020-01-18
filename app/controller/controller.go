package controller

import (
	"blog/app/library/response"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	gmvc.Controller
}

// json返回
func (c *Controller) ResponseJson(request *ghttp.Request, err int, msg string, data ...interface{}) {
	response.JsonExit(request, err, msg, data)
}

func (c *Controller) ResponseSuccess(request *ghttp.Request,msg string, data ...interface{}) {
	c.ResponseJson(request, 200, msg, data)
}

func (c *Controller) ResponseFail(request *ghttp.Request,msg string) {
	c.ResponseJson(request, 100, msg)
}
