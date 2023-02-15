package analyze

type AnalyzeServiceCacheI interface {
	GetMessageCache(input string) string
	SetWordCloudCache(UUID string, value *map[string]int) error
	SetFilterCloudCache(UUID string, value *map[string]int) error
	SetStringRankCache(UUID string, value *[]string) error
	SetAmountRankCache(UUID string, value *[]int) error
}
