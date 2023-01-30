package message

import "context"

type MessageService struct {
	messageServiceRepo MessageServiceRepository
}

type MessageServiceParam struct {
	MessageServiceRepo MessageServiceRepository
}

func NewMessageService(_ context.Context, param MessageServiceParam) *MessageService {
	return &MessageService{
		messageServiceRepo: param.MessageServiceRepo,
	}
}
