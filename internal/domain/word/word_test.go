package word

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"github.com/go-ego/gse"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

var TestConfig = bigcache.Config{
	// number of shards (must be a power of 2)
	Shards: 1024,

	// time after which entry can be evicted  時間到後Key判定死亡 但不刪除
	LifeWindow: 20 * time.Second,
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

func Test_getCacheWordCloud(t *testing.T) {
	type args struct {
		UUID   string
		result *map[string]int
		cache  *bigcache.BigCache
	}
	// var tes int = 0
	tests := []struct {
		name        string
		SetupArgs   func(t *testing.T) args // 修改成func 有利於input傳遞參數，將會call到的func分離，原本只有 arg的Stuct
		CheckExpect func(t *testing.T, err error)
	}{
		// TODO: Add test cases.
		{
			name: "Suceess Test1 GetCache Success",
			SetupArgs: func(t *testing.T) args {
				var input args
				input.UUID = "tes1"
				input.result = &map[string]int{}
				cache, _ := bigcache.New(context.Background(), TestConfig)
				// set cache
				setCacheWordCloud(input.UUID, input.result, cache)
				input.cache = cache
				return input
			},
			CheckExpect: func(t *testing.T, err error) {
				require.Equal(t, nil, err)
			},
		},
		{
			name: "Fail Test2 GetCache",
			SetupArgs: func(t *testing.T) args {
				var input args
				input.UUID = "test2"
				input.result = &map[string]int{}
				cache, _ := bigcache.New(context.Background(), TestConfig)
				// set cache
				// setCacheWordCloud(input.UUID, input.result, cache)
				input.cache = cache
				return input
			},
			CheckExpect: func(t *testing.T, err error) {
				require.Equal(t, "Entry not found", err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.SetupArgs(t)
			err := getCacheWordCloud(args.UUID, args.result, args.cache)
			tt.CheckExpect(t, err)
			// if err := getCacheWordCloud(args.UUID, args.result, args.cache); (err != nil) != tt.wantErr {
			// 	t.Errorf("getCacheWordCloud() error = %v, wantErr %v", err, tt.wantErr)
			// }
		})
	}
}

func Test_setCacheWordCloud(t *testing.T) {
	type args struct {
		UUID    string
		Infomap *map[string]int
		cache   *bigcache.BigCache
	}
	tests := []struct {
		name      string
		SetupArgs func(t *testing.T) args
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "OK Test1 GetCache",
			SetupArgs: func(t *testing.T) args {
				var input args
				input.UUID = "tes1"
				input.Infomap = &map[string]int{}
				cache, _ := bigcache.New(context.Background(), TestConfig)
				input.cache = cache
				return input
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.SetupArgs(t)
			if err := setCacheWordCloud(args.UUID, args.Infomap, args.cache); (err != nil) != tt.wantErr {
				t.Errorf("setCacheWordCloud() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// use sonic
func BenchmarkWordCloud(b *testing.B) {

	testCache, err := bigcache.New(context.Background(), TestConfig)
	if err != nil {
		log.Fatal("cannot create :", err)
	}
	segmentor := &gse.Segmenter{
		AlphaNum: true,
	}
	err = segmentor.LoadDict()
	if err != nil {
		log.Fatal("cannot create :", err)
	}
	path := "/workspace/go-line-message-analyzer/internal/test_en_line.txt"
	b.ResetTimer() // 如果某個函數再回圈內 又不需要計算benchmark，則在該函數上下一行加入 b.StopTimer() b.StartTimer()

	for i := 0; i < b.N; i++ {
		WordCould(path, testCache, segmentor)
	}
	b.StopTimer()
}
func BenchmarkWordCount(b *testing.B) {

	segmentor := &gse.Segmenter{
		AlphaNum: true,
	}
	err := segmentor.LoadDict()
	if err != nil {
		log.Fatal("cannot create :", err)
	}
	path := "/workspace/go-line-message-analyzer/internal/test_en_line.txt"

	wordCountTmp := make(map[string]int)
	// senderMap map
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// b.StopTimer()
		for scanner.Scan() {
			// do something with a line
			line := strings.Split(scanner.Text(), "\t")
			if len(line) > 2 {
				_, _, content := line[0], line[1], line[2]
				words := segmentor.Slice(content)
				b.StartTimer()
				wordCount(&wordCountTmp, words)
				b.StopTimer()
			}
		}
	}

}
func BenchmarkBigCacheUseJSON(b *testing.B) {
	// 前置作業
	testCache, err := bigcache.New(context.Background(), TestConfig)
	if err != nil {
		log.Fatal("cannot create :", err)
	}
	segmentor := &gse.Segmenter{
		AlphaNum: true,
	}
	err = segmentor.LoadDict()
	if err != nil {
		log.Fatal("cannot create :", err)
	}
	data := PrePare(testCache, segmentor)
	b.ResetTimer() // 直接reset benchmark 可省略上面運行時間
	for i := 0; i < b.N; i++ {
		JSONSetCacheWordCloud("testUUID", &data, testCache)
		JSONGetCacheWordCloud("testUUID", &data, testCache)
	}
	b.StopTimer()
}

func JSONGetCacheWordCloud(UUID string, result *map[string]int, cache *bigcache.BigCache) error {

	// fmt.Print(cache)
	prefixKey := UUID + ":wordCloud"
	CacheValue, err := cache.Get(prefixKey)
	if err != nil {
		return err
	}
	err = json.Unmarshal(CacheValue, &result)
	if err != nil {
		return err
	}
	return err

}
func JSONSetCacheWordCloud(UUID string, Infomap *map[string]int, cache *bigcache.BigCache) error {

	prefixKey := UUID + ":wordCloud"
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
func PrePare(cache *bigcache.BigCache, segmenter *gse.Segmenter) map[string]int {
	path := "/workspace/go-line-message-analyzer/internal/test_en_line.txt"
	segmentor := &gse.Segmenter{
		AlphaNum: true,
	}

	// seg    gse.Segmenter
	// posSeg pos.Segmenter
	err := segmentor.LoadDict()
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

			words := segmentor.Slice(content)
			wordCount(&wordCountMap, words)
		}
	}
	return wordCountMap
}
