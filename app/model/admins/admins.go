package admins

// Fill with you ideas below.

func GetAdminByUsername(username string) (*Entity, error) {
	return  FindOne("username = ?", username)

}

func GetAdminById(id int) (*Entity, error) {
	return FindOne("id = ?", id)
}

func GetAdminListPage(page, limit int) (admins []*Entity, count int, err error) {
	count, err = Model.Count()
	admins, err = Model.Page(page, limit).All()
	return
}
