package admin_service

import (
	"blog/app/helpers/client"
	"blog/app/helpers/jwt"
	"blog/app/model/admin"
	"errors"
	"github.com/gogf/gf/net/ghttp"
)

// 用户登录，成功返回用户信息
func Login(username, password string) (entity *admin.Entity, error error) {
	result, err := admin.FindOne("username", username)
	if err != nil {
		return
	}
	if result == nil {
		return nil, errors.New("账号不存在")
	}
	if result.Password != password {
		return nil, errors.New("账号或密码错误")
	}
	accessToken, err := jwt.GenerateToken(result.Id,result.Username,"admin")
	_, _ = client.HSet("admin_token", result.Id, accessToken)
	return result, err
}

// 用户注销
func LogOut(session *ghttp.Session) error {
	return session.Remove("admin_info")
}

// 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains("admin_info")
}
