package ws

import (
	"websocket/internal/models"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *models.Chat
	ID       string `json:"id"`
	RoomId   string `json:"room_id"`
	Username string `json:"username"`
}

type ChatRoom struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Client    map[*Client]bool
}
