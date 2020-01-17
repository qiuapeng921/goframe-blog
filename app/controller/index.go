package controller

import (
	"github.com/gogf/gf/net/ghttp"
)

func Index(request *ghttp.Request) {
	request.Response.Writeln("app")
}