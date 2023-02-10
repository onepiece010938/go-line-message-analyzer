package main_test

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/allegro/bigcache/v3"
	"github.com/bytedance/sonic"
	"github.com/go-ego/gse"
)

func BenchmarkPrepare(b *testing.B) {
	// data := PrePare()
	for i := 0; i < b.N; i++ {
		PrePare()
	}
}
func BenchmarkBigCache(b *testing.B) {
	data := PrePare()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SetCacheWordCloud("testBigCache", &data)
		GetCacheWordCloud("testBigCache")
	}
	b.StopTimer()
}
func BenchmarkJSONMarshalToBigCache(b *testing.B) {
	data := PrePare()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JSONSetCacheWordCloud("testBigCache", &data)
		JSONGetCacheWordCloud("testBigCache")
	}

	b.StopTimer()
}
func BenchmarkSONICBigCache(b *testing.B) {
	data := PrePare()
	fmt.Println(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// SONICSetCacheWordCloud("testBigCache", &data)
		SONICGetCacheWordCloud("testBigCache")
	}
	b.StopTimer()
}
func BenchmarkSONICShortBigCache(b *testing.B) {
	data := PrePare()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SetSONICShortWordCloud("testBigCache", &data)
		// GetSONICShortWordCloud("testBigCache")
	}
	b.StopTimer()
}

var (
	WordCacheKey    = "WordCould:" // + UUID
	MessageCacheKey = "Message:"
	CallCacheKey    = "Call:"
	keyInfo         = "KeyInfo"
	config          = bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted  時間到後Key判定死亡 但不刪除
		LifeWindow: 60 * time.Second,
		// LifeWindow: 10 * time.Minute,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		// CleanWindow: 5 * time.Minute,
		CleanWindow: 30 * time.Second,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: true,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}
)
var cache *bigcache.BigCache

func NewCache() (*bigcache.BigCache, error) {
	cache, initErr := bigcache.New(context.Background(), config)
	if initErr != nil {
		log.Fatal(initErr)
	}

	return cache, nil
}
func TestBigCache(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// ctx := context.Background()
	cache, initErr := bigcache.New(ctx, config)
	if initErr != nil {
		log.Fatal(initErr)
	}
	cache.Set("test1", []byte("aaa"))
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop channel")
				return
			default:
				// cache.Get()
				// fmt.Println("still working")
				GetSONICShortWordCloud("test1", cache)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	GetSONICShortWordCloud("aaa", cache)
	time.Sleep(5 * time.Second)
	cache.Set("test1", []byte("aaa"))
	GetSONICShortWordCloud("aaa", cache)
	time.Sleep(5 * time.Second)
}
func wordSegment(path string, segmenter *gse.Segmenter) map[string]int {

	wordCountMap := map[string]int{}
	// open file
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	// wordsSlice := []string{}
	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), "\t")
		if len(line) > 2 {
			_, _, content := line[0], line[1], line[2]

			// wordCountMap
			words := segmenter.Slice(content)
			wordCount(&wordCountMap, words)
		}
	}
	// scaner後寫入

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wordCountMap

}
func PrePare() map[string]int {
	path := "/workspace/go-line-message-analyzer_ken/test_en_line.txt"
	segmentor := &gse.Segmenter{
		AlphaNum: true,
	}

	// seg    gse.Segmenter
	// posSeg pos.Segmenter
	err := segmentor.LoadDict()
	cache, err = NewCache()
	if err != nil {
		log.Fatal("cannot create cache:", err)
	}
	wordCountMap := make(map[string]int)
	// open file
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	// wordsSlice := []string{}
	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), "\t")
		if len(line) > 2 {
			_, _, content := line[0], line[1], line[2]

			// wordCountMap
			words := segmentor.Slice(content)
			wordCount(&wordCountMap, words)
		}
	}
	return wordCountMap
}
func isDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

func wordCount(wordCountMap *map[string]int, wordsSlice []string) {
	for _, word := range wordsSlice {
		rune_word := []rune(word) // for chinese
		if len(rune_word) <= 1 {
			continue
		} else if isDigit(word) {
			continue
		} else if word == "" {
			continue
		} else {
			(*wordCountMap)[word] = (*wordCountMap)[word] + 1
		}
	}

}
func GetCacheWordCloud(UUID string) (*map[string]int, error) {

	// fmt.Print(cache)
	prefixKey := WordCacheKey + UUID + ":"
	CacheValue, err := cache.Get(keyInfo + prefixKey)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	ressultMap := make(map[string]int)

	splitCache := strings.Split(string(CacheValue), " ")
	for _, key := range splitCache[:len(splitCache)-1] {
		CacheValue, err = cache.Get(prefixKey + key)
		intCacheValue, err := strconv.Atoi(string(CacheValue))
		if err != nil {
			return nil, err
		}
		ressultMap[key] = intCacheValue
	}

	return &ressultMap, err
}
func SetCacheWordCloud(UUID string, Infomap *map[string]int) error {

	prefixKey := WordCacheKey + UUID + ":"
	wordCloudKeyInfo := ""
	for key, value := range *Infomap {
		valueString := strconv.Itoa(value) // to byte
		cacheKey := prefixKey + key
		err := cache.Set(cacheKey, []byte(valueString))
		if err != nil {
			fmt.Println("set value error")
			return err
		}
		wordCloudKeyInfo = wordCloudKeyInfo + key + " " // 用空白區隔 取得時依造空白判斷
	}

	err := cache.Set(keyInfo+prefixKey, []byte(wordCloudKeyInfo)) // 有bug 這個key有機會存在 但對應到的key會不再
	if err != nil {

		return err
	}
	return nil
}
func GetSONICShortWordCloud(UUID string, cache *bigcache.BigCache) error {

	// fmt.Print(cache)
	prefixKey := UUID
	CacheValue, err := cache.Get(prefixKey)
	if err != nil {
		fmt.Print(err)
		return err
	}
	p := ""
	err = sonic.Unmarshal(CacheValue, &p)
	// ressultKeyMap := make(map[string]int)

	// splitCache := strings.Split(p, " ")

	fmt.Println("live")
	return err
}
func SetSONICShortWordCloud(UUID string, Infomap *map[string]int) error {

	prefixKey := WordCacheKey + UUID + ":"
	wordCloudKeyInfo := ""
	for key, value := range *Infomap {

		cacheKey := prefixKey + key
		str, err := sonic.Marshal(value)
		if err != nil {
			fmt.Println("set Marshal error")
			return err
		}
		err = cache.Set(cacheKey, str)
		if err != nil {
			fmt.Println("set value error")
			return err
		}
		wordCloudKeyInfo = wordCloudKeyInfo + key + " " // 用空白區隔 取得時依造空白判斷
	}
	str, err := sonic.Marshal(wordCloudKeyInfo)
	if err != nil {
		fmt.Println("set Marshal error")
		return err
	}
	err = cache.Set(keyInfo+prefixKey, str) // 有bug 這個key有機會存在 但對應到的key會不再
	if err != nil {

		return err
	}
	return nil
}
func SONICGetCacheWordCloud(UUID string) (*map[string]int, error) {

	// fmt.Print(cache)
	prefixKey := WordCacheKey + UUID + ":"
	CacheValue, err := cache.Get(prefixKey)
	ressultMap := make(map[string]int)
	if err != nil {
		return &ressultMap, err
	}
	err = sonic.Unmarshal(CacheValue, &ressultMap)
	if err != nil {
		fmt.Println()
		return &ressultMap, err
	}
	return &ressultMap, err

}
func SONICSetCacheWordCloud(UUID string, Infomap *map[string]int) error {

	prefixKey := WordCacheKey + UUID + ":"
	str, err := sonic.Marshal(Infomap)
	if err != nil {

		return err
	}
	err = cache.Set(prefixKey, str)
	if err != nil {

		return err
	}
	return nil
}

func JSONGetCacheWordCloud(UUID string) (*map[string]int, error) {

	// fmt.Print(cache)
	prefixKey := WordCacheKey + UUID + ":"
	CacheValue, err := cache.Get(prefixKey)
	ressultMap := make(map[string]int)
	if err != nil {

		return &ressultMap, err
	}
	err = json.Unmarshal(CacheValue, &ressultMap)
	if err != nil {
		fmt.Println()
		return &ressultMap, err
	}
	return &ressultMap, err

}
func JSONSetCacheWordCloud(UUID string, Infomap *map[string]int) error {

	prefixKey := WordCacheKey + UUID + ":"
	str, err := json.Marshal(Infomap)
	if err != nil {

		return err
	}
	err = cache.Set(prefixKey, str)
	if err != nil {

		return err
	}
	return nil
}
