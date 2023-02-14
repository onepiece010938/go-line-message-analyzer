package word

import (
	"context"
)

type CreateWordParm struct {
	// CreateHistogramParams postgres.CreateHistogramParams
}

func (i *WordService) WordCloud(ctx context.Context, username string) (*map[string]int, error) {

	var value *map[string]int
	err := i.wordServiceCache.GetWordCloudCache(username, value)
	if err == nil {
		return value, err
	}
	/*
		do something or not
	*/
	return nil, err
}
