package auth

type AuthService interface {
}

type authservice struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *authservice {
	return &authservice{
		repo: repo,
	}
}
