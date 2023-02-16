package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-ego/gse"
	"github.com/onepiece010938/go-line-message-analyzer/internal/adapter/cache"
	"github.com/onepiece010938/go-line-message-analyzer/internal/app"
	"github.com/onepiece010938/go-line-message-analyzer/internal/router"

	"github.com/gin-gonic/gin"
)

var PORT int = 6666

func StartServer() {
	rootCtx, rootCtxCancelFunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	segmentor := &gse.Segmenter{ // 暫時的
		AlphaNum: true,
	}
	err := segmentor.LoadDict()
	if err != nil {
		fmt.Println(err)
	}
	cache := cache.NewCache(cache.InitBigCache(rootCtx))
	app := app.NewApplication(rootCtx, cache, segmentor)

	ginRouter := InitRouter(rootCtx, app)
	// Run server
	wg.Add(1)
	runHTTPServer(rootCtx, &wg, ginRouter, PORT)

	// Listen to SIGTERM/SIGINT to close
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)
	<-gracefulStop
	rootCtxCancelFunc()

	// Wait for all services to close with a specific timeout
	var waitUntilDone = make(chan struct{})
	go func() {
		wg.Wait()
		close(waitUntilDone)
	}()
	select {
	case <-waitUntilDone:
		log.Println("success to close all services")
	case <-time.After(10 * time.Second):
		log.Println(context.DeadlineExceeded, "fail to close all services")
	}
}

func InitRouter(rootCtx context.Context, app *app.Application) (ginRouter *gin.Engine) {
	// Set to release mode to disable Gin logger
	gin.SetMode(gin.ReleaseMode)

	// Create gin router
	ginRouter = gin.New()

	// Set general middleware
	router.SetGeneralMiddlewares(rootCtx, ginRouter)

	// Register all handlers
	router.RegisterHandlers(ginRouter, app)

	return ginRouter
}
func runHTTPServer(rootCtx context.Context, wg *sync.WaitGroup, ginRouter *gin.Engine, port int) {

	// Build HTTP server
	httpAddr := fmt.Sprintf("0.0.0.0:%d", port)
	server := &http.Server{
		Addr:    httpAddr,
		Handler: ginRouter,
	}

	// Run the server in a goroutine
	go func() {
		log.Printf("HTTP server is on http://%s", httpAddr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Panicln(err, "addr", httpAddr, "fail to start HTTP server")
		}
	}()

	// Wait for rootCtx done
	go func() {
		<-rootCtx.Done()

		// Graceful shutdown http server with a timeout
		log.Println("HTTP server is closing")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Println(err, "fail to shutdown HTTP server")
		}

		// Notify when server is closed
		log.Println("HTTP server is closed")
		wg.Done()
	}()
}
