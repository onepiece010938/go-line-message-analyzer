package v1

import (
	"fmt"
	"net/http"

	"github.com/onepiece010938/go-line-message-analyzer/internal/app"
	"github.com/onepiece010938/go-line-message-analyzer/internal/app/service/analyze"

	"github.com/gin-gonic/gin"
)

func StartAnalyze(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		err := app.AnalyzeService.CreateAnalyze(ctx, analyze.CreateAnalyzeParm{})
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, "")
	}

}
func FakeStartAnalyze(app *app.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		_, header, err := c.Request.FormFile("fileupload")
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(p)
		app.AnalyzeService.CreateAnalyze(ctx, analyze.CreateAnalyzeParm{Header: header})
		c.JSON(http.StatusOK, "")
		//curl --form "fileupload=@my-file.txt" https://example.com/resource.cgi
	}

}
