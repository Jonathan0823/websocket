package user

type userhandler struct {
	service UserService
}

func UserHandler(service UserService) *userhandler {
	return &userhandler{
		service: service,
	}
}