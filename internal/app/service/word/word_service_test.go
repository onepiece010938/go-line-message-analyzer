package word

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-ego/gse"
	"github.com/onepiece010938/go-line-message-analyzer/internal/adapter/cache"
)

func BenchmarkSHA256(b *testing.B) {
	name := "pkcs7 簽章使用 RSA 加密演算法對資料的 SHA256 雜湊值簽章，台灣的金融機構習慣對這簽章做 base64 編碼來避免古早用 Cobol 的系統以 ASCII 字碼接收而產生所有資料第 8 bit 都是 0 而引起的驗證錯誤"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := getSHA256HashCode(name + fmt.Sprint(i))
		fmt.Println(p)
	}
	b.StopTimer()
}

func TestWordService_FilterCloud(t *testing.T) {
	type fields struct {
		wordServiceCache WordServiceCacheI
		wordSegmentor    *gse.Segmenter
	}
	type args struct {
		ctx      context.Context
		WordParm FilterWordParm
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *map[string]int
		wantErr bool
	}{
		{
			name: "ok test1",
			fields: fields{wordServiceCache: cache.NewCache(cache.InitBigCache(context.Background())),
				wordSegmentor: &gse.Segmenter{AlphaNum: true}},
			args: args{ctx: context.Background(),
				WordParm: FilterWordParm{&map[string]int{"test_name1": 1, "test_name2": 2, "test_name3": 3},
					"user1", []string{}}},
			want: &map[string]int{"test_name1": 1, "test_name2": 2, "test_name3": 3},
		},
		{
			name: "ok test2",
			fields: fields{wordServiceCache: cache.NewCache(cache.InitBigCache(context.Background())),
				wordSegmentor: &gse.Segmenter{AlphaNum: true}},
			args: args{ctx: context.Background(),
				WordParm: FilterWordParm{&map[string]int{"test_name1": 1, "test_name2": 2, "test_name3": 3},
					"user1", []string{"test_name1"}}},
			want: &map[string]int{"test_name2": 2, "test_name3": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WordService{
				wordServiceCache: tt.fields.wordServiceCache,
				wordSegmentor:    tt.fields.wordSegmentor,
			}
			i.wordSegmentor.LoadDict()
			got, err := i.FilterCloud(tt.args.ctx, tt.args.WordParm)
			if (err != nil) != tt.wantErr {
				t.Errorf("WordService.FilterCloud() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordService.FilterCloud() = %v, want %v", got, tt.want)
			}
			fmt.Println("bb", got)
			fmt.Println("aaa", tt.args.WordParm.WordCloud)
		})
	}
}
