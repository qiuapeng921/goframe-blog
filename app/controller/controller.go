package controller

import (
	"blog/app/consts"
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

func (c *Controller) ResponseSuccess(request *ghttp.Request, data ...interface{}) {
	c.ResponseJson(request, consts.SUCCESS, consts.GetMsgByCode(consts.SUCCESS), data)
}

func (c *Controller) ResponseFail(request *ghttp.Request, msg ...string) {
	c.ResponseJson(request, consts.ERROR, consts.GetMsgByCode(consts.ERROR))
}

func (c *Controller) Render(path string, data ...interface{}) {
	c.View.Assign("name", "GoFrame")
	_ = c.View.Display(path)
	c.Exit()
}
