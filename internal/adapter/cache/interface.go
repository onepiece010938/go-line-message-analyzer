package cache

type CacheI interface {
	GetMessageCache(input string) string

	SetWordCloudCache(UUID string, value *map[string]int) error
	SetFilterCloudCache(UUID string, value *map[string]int) error
	SetStringRankCache(UUID string, value *[]string) error
	SetAmountRankCache(UUID string, value *[]int) error

	GetWordCloudCache(UUID string, value *map[string]int) error
	GetFilterCloudCache(UUID string, value *map[string]int) error
	GetStringRankCache(UUID string, value *[]string) error
	GetAmountRankCache(UUID string, value *[]int) error

	DeleteWordCloudCache(UUID string) error
	DeleteFilterCloudCache(UUID string) error
	DeleteStringRankCache(UUID string) error
	DeleteAmountRankCache(UUID string) error
}
