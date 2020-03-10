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

func (c *Controller) ResponseSuccess(request *ghttp.Request, responseData interface{}) {
	response.JsonExit(request, consts.SUCCESS, "success", responseData)
}

func (c *Controller) ResponseFail(request *ghttp.Request, msg string) {
	response.JsonExit(request, consts.ERROR, msg,"null")
}
