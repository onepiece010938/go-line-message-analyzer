package cache

import (
	"context"
	"fmt"
	"math/rand"

	"testing"

	"github.com/allegro/bigcache/v3"
)

func TestCache_SetWordCloudCache(t *testing.T) {
	type fields struct {
		cache *bigcache.BigCache
	}
	type args struct {
		UUID   string
		result *map[string]int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				cache: tt.fields.cache,
			}
			if err := c.SetWordCloudCache(tt.args.UUID, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Cache.SetWordCloudCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkSetWordCloudCache(b *testing.B) {
	// Cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Second))
	Result := make(map[string]int)
	for i := 0; i < 10000; i++ {
		testNumber := "test" + fmt.Sprint(i)
		Result[testNumber] = rand.Intn(1000)
	}
	initCache := InitBigCache(context.Background())
	var Cache = Cache{cache: initCache}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UUID := fmt.Sprint(i)
		Cache.SetWordCloudCache(UUID, &Result)
	}
	b.StopTimer()
}
