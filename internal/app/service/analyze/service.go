package analyze

import (
	"context"

	"github.com/go-ego/gse"
)

type AnalyzeService struct {
	analyzeServiceCache AnalyzeServiceCacheI
	analyzeSegmentor    *gse.Segmenter
}

type AnalyzeServiceParam struct {
	AnalyzeServiceCache AnalyzeServiceCacheI
	AnalyzeSegmentor    *gse.Segmenter
}

func NewAnalyzeService(_ context.Context, param AnalyzeServiceParam) *AnalyzeService {
	return &AnalyzeService{
		analyzeServiceCache: param.AnalyzeServiceCache,
		analyzeSegmentor:    param.AnalyzeSegmentor,
	}
}
