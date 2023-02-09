package app

import (
	"context"

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

	// Create application
	app := &Application{
		MessageService: serviceMessage.NewMessageService(ctx, serviceMessage.MessageServiceParam{
			MessageServiceCache: cache,
		}),
		AnalyzeService: serviceAnalyze.NewAnalyzeService(ctx, serviceAnalyze.AnalyzeServiceParam{
			AnalyzeServiceCache: cache,
		}),
	}
	return app
}
