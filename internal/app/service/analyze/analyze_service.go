package analyze

import (
	"context"
	"fmt"
)

type CreateAnalyzeParm struct {
	// CreateHistogramParams postgres.CreateHistogramParams
}

type AnalyzeMessageParm struct {
}

func (i *AnalyzeService) CreateAnalyze(ctx context.Context, param CreateAnalyzeParm) error {
	fmt.Println("AnalyzeService->func CreateAnalyze()")
	// var result message.MessageDomainResult
	// result = message.MessageDomainFunc("aabcccc")
	// fmt.Println(result)
	i.messageCache.GetMessageCache("")

	return nil
}
