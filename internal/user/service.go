package user

type UserService interface {
}

type userservice struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userservice {
	return &userservice{
		repo: repo,
	}
}
