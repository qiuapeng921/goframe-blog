package response

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func json(request *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	err := request.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
	if err != nil {
		glog.Debug(err.Error())
	}
	request.Exit()
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(request *ghttp.Request, code int, msg string, data ...interface{}) {
	json(request, code, msg, data...)
}
