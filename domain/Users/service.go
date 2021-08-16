package users

import userdto "www.github.com/Ammce/banking-go-auth/dto/userDTO"

type UserService interface {
	Register(userdto.CreateUser) (*User, *error)
}

type DefaultUserService struct {
	repo UserRepository
}

func (us DefaultUserService) Register(user userdto.CreateUser) (*User, *error) {
	newUser := NewUser(user.Email, user.Password, user.Role)
	u, e := us.repo.Register(*newUser)
	if e != nil {
		return nil, e
	}
	return u, nil
}

func NewUserService(repo UserRepository) *DefaultUserService {
	return &DefaultUserService{
		repo: repo,
	}
}
