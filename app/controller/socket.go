package controller

import (
	"fmt"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

type SocketController struct {
	Controller
}

var (
	socket *ghttp.WebSocket
	// 使用默认的并发安全Map
	users = gmap.New(true)
	// 使用并发安全的Set，用以用户昵称唯一性校验
	names = gset.NewStrSet(true)
)

// Msg 消息结构体
type Msg struct {
	Type string      `json:"type" v:"type@required#消息类型不能为空"`
	Data interface{} `json:"data" v:""`
	From string      `json:"name" v:""`
}

func (c *SocketController) Socket() {
	server, err := onOpen(c.Request)
	if err != nil {
		glog.Error(err)
		c.Exit()
	}
	name := c.Request.GetString("name")
	// 初始化时设置用户昵称为当前链接信息
	names.Add(name)
	users.Set(server, name)
	users.Set(name, server)

	socket = server
	_ = server.WriteMessage(ghttp.WS_MSG_BINARY, gconv.Bytes("welcome to you"))
	for {
		msgType, msg, err := server.ReadMessage()
		if err != nil {
			glog.Debug("用户下线", err.Error())
			names.Remove(name)
			users.Remove(socket)
			// 通知所有客户端当前用户已下线
			break
		}
		if err := onMessage(msgType, msg); err != nil {
			glog.Debug("接收消息失败", err.Error())
			return
		}
		// 日志记录
		g.Log().Cat("chat").Println(msg)
	}
}

func onOpen(request *ghttp.Request) (*ghttp.WebSocket, error) {
	server, err := request.WebSocket()
	if err != nil {
		return nil, err
	}
	return server, nil
}

func onMessage(msgType int, msg []byte) error {
	err := socket.WriteMessage(msgType, msg)
	if err != nil {
		return err
	}
	return nil
}

func onClone(code int, text string) error {
	fmt.Println(code, text)
	return nil
}

// 向客户端写入消息。
func push(msg []byte) error {
	return socket.WriteMessage(ghttp.WS_MSG_TEXT, msg)
}

// 向所有客户端群发消息。
func pushAll(msg []byte) error {
	users.RLockFunc(func(m map[interface{}]interface{}) {
		for user := range m {
			_ = user.(*ghttp.WebSocket).WriteMessage(ghttp.WS_MSG_TEXT, msg)
		}
	})

	return nil
}

func getUser(name interface{}) interface{} {
	return users.Get(name)
}

// 获取所有用户
func userList() (interface{}, error) {
	array := garray.NewSortedStrArray()
	names.Iterator(func(v string) bool {
		array.Add(v)
		return true
	})
	return array.Slice(), nil
}
