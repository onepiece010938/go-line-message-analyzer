package histogram

import "context"

type HistogramService struct {
	histogramServiceRepo HistogramServiceRepository
}

type HistogramServiceParam struct {
	HistogramServiceRepo HistogramServiceRepository
}

func NewHistogramService(_ context.Context, param HistogramServiceParam) *HistogramService {
	return &HistogramService{
		histogramServiceRepo: param.HistogramServiceRepo,
	}
}
