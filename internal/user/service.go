package user

import "websocket/internal/models"

type UserService interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUsers() ([]models.User, error)
}

type userservice struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userservice {
	return &userservice{
		repo: repo,
	}
}

func (s *userservice) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userservice) GetUserByID(id string) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userservice) GetUsers() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}