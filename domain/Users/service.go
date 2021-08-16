package users

type UserService interface {
	Register(User) (*User, *error)
}

type DefaultUserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *DefaultUserService {
	return &DefaultUserService{
		repo: repo,
	}
}
