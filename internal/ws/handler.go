package ws

import (
	"net/http"
	"websocket/internal/models"

	"github.com/gin-gonic/gin"
	"strconv"
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

	h.hub.Rooms[req.ID] = &models.ChatRoom{
		ID:     req.ID,
		Name:   req.Name,
		Client: make(map[*models.Client]bool),
	}

	c.JSON(http.StatusCreated, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *wshandler) HandleWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientIdStr := c.Query("userId")
	clientId, err := strconv.Atoi(clientIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}
	roomId := c.Query("roomId")
	username := c.Query("username")

	client := &models.Client{
		Conn:    conn,
		Message: make(chan *models.Chat, 10),
		ID:      clientIdStr,
		RoomId:  roomId,
		Username: username,
	}

	m:= &models.Chat{
		Message:  "New user joined",
		ChatRoomId: roomId,
		UserId: clientId,
	}		

	h.hub.Register <- client

	h.hub.Broadcast <- m
}
