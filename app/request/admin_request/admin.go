package admin_request

type AdminRequest struct {
	Page  int
	Limit int
}

type AdminCreateRequest struct {
	UserName string
	Password string
	Phone    string
}

type AdminUpdateRequest struct {
}
