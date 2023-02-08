package analyze

import "context"

type AnalyzeService struct {
	messageCache MessageCache
}

type AnalyzeServiceParam struct {
	MessageCache MessageCache
}

func NewAnalyzeService(_ context.Context, param AnalyzeServiceParam) *AnalyzeService {
	return &AnalyzeService{
		messageCache: param.MessageCache,
	}
}
