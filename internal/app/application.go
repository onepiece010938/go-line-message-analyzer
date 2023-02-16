package app

import (
	"context"

	"github.com/go-ego/gse"
	"github.com/onepiece010938/go-line-message-analyzer/internal/adapter/cache"
	serviceAnalyze "github.com/onepiece010938/go-line-message-analyzer/internal/app/service/analyze"
	serviceMessage "github.com/onepiece010938/go-line-message-analyzer/internal/app/service/message"
	serviceWord "github.com/onepiece010938/go-line-message-analyzer/internal/app/service/word"
)

type Application struct {
	// JobService   *serviceJob.JobService
	// ImageService *serviceImage.ImageService
	AnalyzeService *serviceAnalyze.AnalyzeService
	MessageService *serviceMessage.MessageService
	WordService    *serviceWord.WordService
}

func NewApplication(ctx context.Context, cache cache.CacheI, segmentor *gse.Segmenter) *Application {

	// Create application
	app := &Application{
		MessageService: serviceMessage.NewMessageService(ctx, serviceMessage.MessageServiceParam{
			MessageServiceCache: cache,
		}),
		AnalyzeService: serviceAnalyze.NewAnalyzeService(ctx, serviceAnalyze.AnalyzeServiceParam{
			AnalyzeServiceCache: cache,
		}),
		WordService: serviceWord.NewWordService(ctx, serviceWord.WordServiceParam{
			WordServiceCache: cache,
			WordSegmentor:    segmentor,
		}),
	}
	return app
}
