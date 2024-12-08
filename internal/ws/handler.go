package ws

import (
	"net/http"
	"websocket/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type wshandler struct {
	hub *Hub
}

func NewWsHandler(hub *Hub) *wshandler {
	return &wshandler{
		hub: hub,
	}
}

type RoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *wshandler) CreateRoom(c *gin.Context) {
	var req RoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &ChatRoom{
		ID:     req.ID,
		Name:   req.Name,
		Client: make(map[*Client]bool),
	}

	c.JSON(http.StatusCreated, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *wshandler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomId := c.Param("roomId")
	clientIdStr := c.Query("userId")
	username := c.Query("username")

	client := &Client{
		Conn:     conn,
		Message:  make(chan *models.Chat, 10),
		ID:       clientIdStr,
		RoomId:   roomId,
		Username: username,
	}

	m := &models.Chat{
		Message:    "New user joined",
		ChatRoomId: roomId,
		UserId:     clientIdStr,
	}

	h.hub.Register <- client

	h.hub.Broadcast <- m

	go client.WriteMessage()
	client.ReadMessage(h.hub)
}

func (h *wshandler) GetRooms(c *gin.Context) {
	rooms := make([]*RoomReq, 0)
	for _, room := range h.hub.Rooms {
		rooms = append(rooms, &RoomReq{
			ID:   room.ID,
			Name: room.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *wshandler) GetClients(c *gin.Context) {
	roomId := c.Param("roomId")
	clients := make([]string, 0)
	if room, ok := h.hub.Rooms[roomId]; ok {
		for client := range room.Client {
			clients = append(clients, client.Username)
		}
	}

	c.JSON(http.StatusOK, clients)
}