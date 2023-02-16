package analyze

import (
	"context"
	"fmt"
	"mime/multipart"

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

	seg := i.analyzeSegmentor
	file, _ := param.Header.Open()
	fileName := param.Header.Filename
	// SplitWordToCloud
	analyze.SplitWordToCloud(fileName, &file, seg)
	return nil
}
