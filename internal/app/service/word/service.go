package word

import (
	"context"
)

type WordService struct {
	wordServiceCache WordServiceCacheI
}

type WordServiceParam struct {
	WordServiceCache WordServiceCacheI
}

func NewWordService(_ context.Context, param WordServiceParam) *WordService {
	return &WordService{
		wordServiceCache: param.WordServiceCache,
	}
}
