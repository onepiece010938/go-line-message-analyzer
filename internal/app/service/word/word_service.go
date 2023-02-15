package word

import (
	"context"
)

func (i *WordService) WordCloud(ctx context.Context, username string) (*map[string]int, error) {

	var value *map[string]int
	err := i.wordServiceCache.GetWordCloudCache(username, value)
	if err != nil {
		return nil, err
	}
	return value, err
}
func (i *WordService) FilterCloud(ctx context.Context, username string) (*map[string]int, error) {
	var value *map[string]int
	err := i.wordServiceCache.GetFilterCloudCache(username, value)
	if err != nil {
		return nil, err
	}
	return value, err
}
func (i *WordService) StringRank(ctx context.Context, username string) (*[]string, error) {
	var value *[]string
	err := i.wordServiceCache.GetStringRankCache(username, value)
	if err != nil {
		return nil, err
	}
	return value, err
}
func (i *WordService) AmountRank(ctx context.Context, username string) (*[]int, error) {
	var value *[]int
	err := i.wordServiceCache.GetAmountRankCache(username, value)
	if err != nil {
		return nil, err
	}
	return value, err
}
