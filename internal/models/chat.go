package models

type Chat struct {
	ID         int    `json:"id"`
	UserId     int    `json:"user_id"`
	ChatRoomId int    `json:"chat_room_id"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
