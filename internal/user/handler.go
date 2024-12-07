package user

type authhandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *authhandler {
	return &authhandler{
		service: service,
	}
}