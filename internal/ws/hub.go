package ws

import "websocket/internal/models"

type Hub struct {
	Rooms map[string]*models.ChatRoom
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*models.ChatRoom),
	}
}