package ws

import "websocket/internal/models"

type Hub struct {
	Rooms      map[string]*models.ChatRoom
	Register   chan *models.Client
	Unregister chan *models.Client
	Broadcast  chan *models.Chat
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*models.ChatRoom),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
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
