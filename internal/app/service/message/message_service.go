package message

import (
	"context"
)

type CreateMessageParm struct {
	// CreateHistogramParams postgres.CreateHistogramParams
}

func (i *MessageService) CreateMessage(ctx context.Context, param CreateMessageParm) error {
	return nil
}
