package users

type UserRepository interface {
	Register(User) (*User, *error)
}
