package auth

import (
	"database/sql"
	"errors"
	"websocket/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	Register(user models.User) error
	Login(user models.User) (string, error)
}

type authrepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authrepository {
	return &authrepository{
		db: db,
	}
}

func (r *authrepository) Register(user models.User) error {
	var existingUser bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", user.Email).Scan(&existingUser)
	if err != nil {
		return err
	}
	if existingUser {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (r *authrepository) Login(user models.User) (string, error) {
	var hashedPassword string
	err := r.db.QueryRow("SELECT password FROM users WHERE email=$1", user.Email).Scan(&hashedPassword)
	if err != nil {
		return "", errors.New("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
