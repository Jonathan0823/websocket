package models

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Chat
	ID       string `json:"id"`
	RoomId   string `json:"room_id"`
	Username string `json:"username"`
}
