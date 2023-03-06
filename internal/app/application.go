package app

import (
	"context"

	_ "github.com/go-ego/gse"
	"github.com/line/line-bot-sdk-go/linebot"
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
	LineBotClient  *linebot.Client
}

func NewApplication(ctx context.Context, cache cache.CacheI, lineBotClient *linebot.Client) *Application {

	// Create application
	app := &Application{
		LineBotClient: lineBotClient,
		MessageService: serviceMessage.NewMessageService(ctx, serviceMessage.MessageServiceParam{
			MessageServiceCache: cache,
		}),
		AnalyzeService: serviceAnalyze.NewAnalyzeService(ctx, serviceAnalyze.AnalyzeServiceParam{
			AnalyzeServiceCache: cache,
		}),
		WordService: serviceWord.NewWordService(ctx, serviceWord.WordServiceParam{
			WordServiceCache: cache,
		}),
	}
	return app
}
