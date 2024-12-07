package chat

import "database/sql"

type ChatRepository interface {
}

type chatrepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *chatrepository {
	return &chatrepository{
		db: db,
	}
}
