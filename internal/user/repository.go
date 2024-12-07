package user

import "database/sql"

type UserRepository interface {
}

type userrepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userrepository{
		db: db,
	}
}
