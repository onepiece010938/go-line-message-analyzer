package word

import (
	"context"

	"github.com/go-ego/gse"
)

type WordService struct {
	wordServiceCache WordServiceCacheI
	wordSegmentor    *gse.Segmenter
}

type WordServiceParam struct {
	WordServiceCache WordServiceCacheI
	WordSegmentor    *gse.Segmenter
}

func NewWordService(_ context.Context, param WordServiceParam) *WordService {
	return &WordService{
		wordServiceCache: param.WordServiceCache,
		wordSegmentor:    param.WordSegmentor,
	}
}
