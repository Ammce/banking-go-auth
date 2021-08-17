package users

import userdto "www.github.com/Ammce/banking-go-auth/dto/userDTO"

type UserService interface {
	Register(userdto.CreateUser) (*userdto.CreateUserResponse, *error)
}

type DefaultUserService struct {
	repo UserRepository
}

func (us DefaultUserService) Register(user userdto.CreateUser) (*userdto.CreateUserResponse, *error) {
	newUser := NewUser(user.Email, user.Password, user.Role)
	u, e := us.repo.Register(*newUser)
	if e != nil {
		return nil, e
	}
	return u.AsCreateUserResponseDTO(), nil
}

func NewUserService(repo UserRepository) *DefaultUserService {
	return &DefaultUserService{
		repo: repo,
	}
}
