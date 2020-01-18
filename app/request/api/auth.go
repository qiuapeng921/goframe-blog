package api

type LoginRequest struct {
	Account  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
}

type RegisterRequest struct {
	Account string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	PasswordOk string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
}
