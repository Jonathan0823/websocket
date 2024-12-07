package user

import "database/sql"

type AuthRepository interface {
}

type authrepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authrepository{
		db: db,
	}
}
