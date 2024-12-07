package chat

type ChatService interface {
}

type chatservice struct {
	repo ChatRepository
}

func NewChatService(repo ChatRepository) *chatservice {
	return &chatservice{
		repo: repo,
	}
}
