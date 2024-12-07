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
	Validate(user models.User) bool
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
	existingUser := r.IsUserExists(user)
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

func (r *authrepository) Validate(user models.User) bool {
	var dbUser models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE email=$1", user.Email).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Email, &dbUser.Password)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	return err == nil
}

func (r *authrepository) Login(user models.User) (string, error) {
	if !r.Validate(user) {
		return "", errors.New("invalid credentials")
	}

	token, err := GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *authrepository) IsUserExists(user models.User) bool {
	var dbUser models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE email=$1", user.Email).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Email, &dbUser.Password)
	return err == nil
}
