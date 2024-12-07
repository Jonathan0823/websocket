package models

type Chat struct {
	ID         int    `json:"id"`
	UserId     int    `json:"user_id"`
	ChatRoomId string `json:"chat_room_id"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
