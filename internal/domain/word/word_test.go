package word

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/stretchr/testify/require"
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
