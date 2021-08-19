package users

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

type DuplicateEmail struct{}

func (m *DuplicateEmail) Error() string {
	return "this email is already taken"
}
