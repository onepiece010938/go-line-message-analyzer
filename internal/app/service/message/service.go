package message

import "context"

type MessageService struct {
	// messageServiceRepo MessageServices
}

type MessageServiceParam struct {
	// MessageServiceRepo MessageServices
}

func NewMessageService(_ context.Context, param MessageServiceParam) *MessageService {
	return &MessageService{
		// messageServiceRepo: param.MessageServiceRepo,
	}
}
