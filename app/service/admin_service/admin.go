package admin_service

import (
	"blog/app/helpers/client"
	"blog/app/helpers/jwt"
	"blog/app/model/admins"
	"blog/app/request/admin_request"
	"errors"
)

// 用户登录，成功返回用户信息
func Login(username, password string) (accessToken string, err error) {
	var result *admins.Entity
	result, err = admins.GetAdminByUsername(username)
	if err != nil {
		return
	}
	if result == nil {
		err = errors.New("账号不存在")
		return
	}
	if result.Password != password {
		err = errors.New("账号或密码错误")
		return
	}
	accessToken, err = jwt.GenerateToken(result.Id, result.Username, "admin")
	_, _ = client.HSet("admin_token", result.Id, accessToken)
	return
}

func LogOut(adminId int64) (interface{}, error) {
	return client.HDel("admin_token", adminId)
}

func GetAdminList(request admin_request.AdminRequest) (adminResult []*admins.Entity, count int, err error) {
	page, limit := 1, 10
	if request.Page > 1 {
		page = request.Page
	}
	if request.Limit > 10 {
		limit = request.Limit
	}
	adminResult, count, err = admins.GetAdminListPage(page, limit)
	return
}

func CreateAdmin(request admin_request.AdminCreateRequest) (id int64, err error) {
	var adminEntity admins.Entity
	adminEntity.Username = request.UserName
	adminEntity.Password = request.Password
	adminEntity.Phone = request.Phone
	result, err := admins.Insert(&adminEntity)
	if err != nil {
		err = errors.New(err.Error())
		return
	}
	id, err = result.LastInsertId()
	return
}

func UpdateAdmin(request admin_request.AdminUpdateRequest) bool {
	return true
}
