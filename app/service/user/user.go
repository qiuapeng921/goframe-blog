package user

import (
	"blog/app/library/client"
	"blog/app/model/user/user"
	"blog/app/request/api"
	"errors"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

// 用户注册
func Register(data *api.RegisterRequest) error {
	// 输入参数检查
	if e := gvalid.CheckStruct(data, nil); e != nil {
		return errors.New(e.String())
	}
	// 账号唯一性数据检查
	if !CheckAccount(data.Account) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", data.Account))
	}
	// 将输入参数赋值到数据库实体对象上
	var entity *user.Entity
	if err := gconv.Struct(data, &entity); err != nil {
		return err
	}
	// 记录账号创建/注册时间
	entity.CreateTime = gtime.Now()
	if _, err := user.Save(entity); err != nil {
		return err
	}
	return nil
}

// 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains("user_info")
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func Login(account, password string, session *ghttp.Session) error {
	result, err := user.FindOne("account=? and password=?", account, password)
	if err != nil {
		return err
	}
	if result == nil {
		return errors.New("账号或密码错误")
	}

	client.HSet("user_info", result.Id, gconv.String(result))
	return session.Set("user_info", result)
}

// 用户注销
func LogOut(session *ghttp.Session) error {
	return session.Remove("user_info")
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckAccount(account string) bool {
	if i, err := user.FindCount("account", account); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 获得用户信息详情
func GetUserInfo(session *ghttp.Session) (user *user.Entity) {
	_ = session.GetStruct("user_info", &user)
	return
}
