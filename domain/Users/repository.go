package users

import "www.github.com/Ammce/banking-go-auth/utils"

type UserRepository interface {
	Register(User) (*User, *error)
}

func NewUser(email string, password string, role string) *User {
	return &User{
		Email:    email,
		Password: utils.Hash(password),
		Role:     role,
	}
}
