package users

// Fill with you ideas below.

func GetUserByUsername(username string) (*Entity, error) {
	return FindOne("username = ?", username)
}

func GetUserById(id uint) (*Entity, error) {
	return FindOne("id = ?", id)
}
