package v1

import (
	"go-line-message-analyzer/internal/app"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup, app *app.Application) {
	v1 := router.Group("/v1")
	{
		v1.GET("/sample", SAMPLE)
	}
}
