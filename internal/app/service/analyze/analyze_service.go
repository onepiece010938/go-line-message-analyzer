package analyze

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/go-ego/gse"
	"github.com/onepiece010938/go-line-message-analyzer/internal/domain/analyze"
)

type CreateAnalyzeParm struct {
	// CreateHistogramParams postgres.CreateHistogramParams
	Header *multipart.FileHeader
}

type AnalyzeMessageParm struct {
}

func (i *AnalyzeService) CreateAnalyze(ctx context.Context, param CreateAnalyzeParm) error {
	fmt.Println("AnalyzeService->func CreateAnalyze()")
	i.analyzeServiceCache.GetMessageCache("")
	// var result message.MessageDomainResult
	// result = message.MessageDomainFunc("aabcccc")
	// fmt.Println(result)

	segmentor := &gse.Segmenter{ // 暫時的
		AlphaNum: true,
	}

	_ = segmentor.LoadDict()

	file, _ := param.Header.Open()
	fileName := param.Header.Filename
	// SplitWordToCloud
	analyze.SplitWordToCloud(fileName, &file, segmentor)
	return nil
}
