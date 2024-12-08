package ws

import (
	"websocket/internal/models"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Rooms      map[string]*ChatRoom
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *models.Chat
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*ChatRoom),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *models.Chat),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.Rooms[client.RoomId]; ok {
				room := h.Rooms[client.RoomId]
				if _, ok := room.Client[client]; !ok {
					room.Client[client] = true
				}
			}
		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.RoomId]; ok {
				room := h.Rooms[client.RoomId]
				if _, ok := room.Client[client]; ok {
					if len(room.Client) != 0 {
						h.Broadcast <- &models.Chat{
							ChatRoomId: client.RoomId,
							Message:    client.Username + " left",
						}
					}

					delete(room.Client, client)
					close(client.Message)
				}
			}
		case message := <-h.Broadcast:
			if _, ok := h.Rooms[message.ChatRoomId]; ok {
				room := h.Rooms[message.ChatRoomId]
				for client := range room.Client {
					select {
					case client.Message <- message:
					default:
						close(client.Message)
						delete(room.Client, client)
					}
				}
			}
		}
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Message {
		if err := c.Conn.WriteJSON(message); err != nil {
			return
		}
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return
			}
			break
		}

		msg := &models.Chat{
			UserId:     c.ID,
			ChatRoomId: c.RoomId,
			Message:    string(message),
		}

		hub.Broadcast <- msg
	}
}
