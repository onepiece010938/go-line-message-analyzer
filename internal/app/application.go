package app

import (
	"context"
	"fmt"

	"github.com/onepiece010938/go-line-message-analyzer/internal/adapter/cache"
	serviceAnalyze "github.com/onepiece010938/go-line-message-analyzer/internal/app/service/analyze"
	serviceMessage "github.com/onepiece010938/go-line-message-analyzer/internal/app/service/message"
)

type Application struct {
	// JobService   *serviceJob.JobService
	// ImageService *serviceImage.ImageService
	AnalyzeService *serviceAnalyze.AnalyzeService
	MessageService *serviceMessage.MessageService
}

func NewApplication(ctx context.Context, cache cache.CacheI) *Application {
	fmt.Println(ctx)
	// New Service
	messageService := serviceMessage.NewMessageService(ctx, serviceMessage.MessageServiceParam{})
	// Create application
	app := &Application{
		MessageService: messageService,
		AnalyzeService: serviceAnalyze.NewAnalyzeService(ctx, serviceAnalyze.AnalyzeServiceParam{
			MessageCache: cache,
		}),
	}
	return app
}
