package user

import (
	"database/sql"
	"errors"
	"websocket/internal/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (models.User, error)
	GetByEmail(email string) (models.User, error)
}

type userrepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userrepository {
	return &userrepository{
		db: db,
	}
}

func (r *userrepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userrepository) GetByID(id string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *userrepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE email=$1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}
