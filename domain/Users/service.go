package users

type UserService interface {
	Register(User) (*User, *error)
}

type DefaultUserService struct {
	repo UserRepository
}

func (us DefaultUserService) Register(user User) (*User, *error) {
	u, e := us.repo.Register(user)
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
