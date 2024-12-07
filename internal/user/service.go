package user

type UserService interface {
}

type userservice struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userservice{
		repo: repo,
	}
}
