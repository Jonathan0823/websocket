package chat

type chathandler struct {
	service ChatService
}

func NewChatHandler(service ChatService) *chathandler {
	return &chathandler{
		service: service,
	}
}
