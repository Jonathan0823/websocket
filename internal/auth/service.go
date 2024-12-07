package auth

type AuthService interface {
}

type authservice struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authservice{
		repo: repo,
	}
}
