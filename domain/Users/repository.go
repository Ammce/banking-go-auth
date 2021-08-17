package users

import (
	"time"

	userdto "www.github.com/Ammce/banking-go-auth/dto/userDTO"
	"www.github.com/Ammce/banking-go-auth/utils"
)

type UserRepository interface {
	Register(User) (*User, *error)
	FindBy(email string) (*User, *error)
}

func (u User) AsCreateUserResponseDTO() *userdto.CreateUserResponse {
	return &userdto.CreateUserResponse{
		ID:        u.ID.Hex(),
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewUser(email string, password string, role string) *User {
	return &User{
		Email:     email,
		Password:  utils.Hash(password),
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
