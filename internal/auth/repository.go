package auth

import "database/sql"

type AuthRepository interface {
}

type authrepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authrepository {
	return &authrepository{
		db: db,
	}
}
