package user

import "database/sql"

type UserRepository interface {
}

type userrepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userrepository {
	return &userrepository{
		db: db,
	}
}
