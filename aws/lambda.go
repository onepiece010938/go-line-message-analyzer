package aws

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/onepiece010938/go-line-message-analyzer/cmd/server"
	"github.com/onepiece010938/go-line-message-analyzer/internal/adapter/cache"
	"github.com/onepiece010938/go-line-message-analyzer/internal/app"
)

type DeployLambda struct {
	GinLambda        *ginadapter.GinLambda
	CacheLambda      *cache.Cache
	LineClientLambda *linebot.Client
	Ssmsvc           *SSM
}

func NewDeployLambda(rootCtx context.Context) *DeployLambda {
	d := &DeployLambda{}
	d.initSSM()
	d.initLineClient()
	d.initCache(rootCtx)
	d.initGinAdapter(rootCtx)
	return d
}

func (d *DeployLambda) initSSM() {
	d.Ssmsvc = NewSSMClient()
}

func (d *DeployLambda) initLineClient() {
	lineSecret, err := d.Ssmsvc.Param("CHANNEL_SECRET", true).GetValue()
	if err != nil {
		log.Println(err)
	}
	lineAccessToken, err := d.Ssmsvc.Param("CHANNEL_ACCESS_TOKEN", true).GetValue()
	if err != nil {
		log.Println(err)
	}
	d.LineClientLambda, err = linebot.New(lineSecret, lineAccessToken)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DeployLambda) initCache(rootCtx context.Context) {
	d.CacheLambda = cache.NewCache(cache.InitBigCache(rootCtx))
}

func (d *DeployLambda) initGinAdapter(rootCtx context.Context) {
	app := app.NewApplication(rootCtx, d.CacheLambda, d.LineClientLambda)
	ginRouter := server.InitRouter(rootCtx, app)
	d.GinLambda = ginadapter.New(ginRouter)
}

func (d *DeployLambda) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return d.GinLambda.ProxyWithContext(ctx, request)
}
