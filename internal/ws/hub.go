package ws

import "websocket/internal/models"

type Hub struct {
	Rooms map[string]*models.ChatRoom
	Register chan *models.Client
	Unregister chan *models.Client
	Broadcast chan *models.Chat
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*models.ChatRoom),
		Register: make(chan *models.Client),
		Unregister: make(chan *models.Client),
		Broadcast: make(chan *models.Chat),
	}
}

func (h *Hub) Run() {
	for _, room := range h.Rooms {
		go room.Run()
	}
}