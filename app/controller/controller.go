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