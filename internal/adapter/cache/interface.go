package cache

type CacheI interface {
	GetMessageCache(input string) string
}
