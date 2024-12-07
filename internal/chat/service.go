package chat

type ChatService interface {
}

type chatservice struct {
	repo ChatRepository
}

func NewChatService(repo ChatRepository) ChatService {
	return &chatservice{
		repo: repo,
	}
}
