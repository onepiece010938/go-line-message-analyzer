package analyze

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/go-ego/gse"
)

func TestWord(t *testing.T) {
	path := "/workspace/go-line-message-analyzer/aaacacs.txt"
	f, err := os.Open(path)
	if err != nil {
		return
	}
	segmentor := &gse.Segmenter{
		AlphaNum: true,
	}

	// seg    gse.Segmenter
	// posSeg pos.Segmenter
	err = segmentor.LoadDict()
	// f.Close()
	wordcloud := SplitWordToCloudAAA("aaa", f, segmentor)
	SortWordCloudRank("aaa", wordcloud)
}
func SplitWordToCloudAAA(UUID string, file *os.File, segmenter *gse.Segmenter) map[string]int {

	WordCloudMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		if len(line) > 2 {
			_, _, content := line[0], line[1], line[2] //nolint
			fmt.Println(content)
			words := segmenter.Cut(content, false)
			for _, p := range words {
				WordCloudMap[p] = WordCloudMap[p] + 1
			}

		}
	}
	for key := range WordCloudMap {
		runeKey := []rune(key)
		if len(runeKey) == 0 || HaveSingalAtoZandSymbolWord(runeKey) {
			delete(WordCloudMap, key)
		}

	}

	return WordCloudMap
}

// func BenchmarkSortWordCloudRank(b *testing.B) {
// 	path := "D:/go-line-message-analyzer/cdcd.txt"
// 	f, err := os.Open(path)
// 	if err != nil {
// 		return
// 	}
// 	segmentor := &gse.Segmenter{
// 		AlphaNum: true,
// 	}

// 	// seg    gse.Segmenter
// 	// posSeg pos.Segmenter
// 	err = segmentor.LoadDict()
// 	// f.Close()
// 	wordcloud := SplitWordToCloud("aaa", f, segmentor)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		SortWordCloudRank("aaa", wordcloud)
// 	}
// 	b.StopTimer()

// }

func Test_HaveSingalAtoZandSymbolWord(t *testing.T) {
	type args struct {
		x []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not have a to z or Punct or Symbol",
			args: args{[]rune("安安你好早安smile")},
			want: false,
		},
		{
			name: "have word",
			args: args{[]rune("abc")},
			want: false,
		},
		{
			name: "have a",
			args: args{[]rune("a")},
			want: true,
		},
		{
			name: "have z",
			args: args{[]rune("a")},
			want: true,
		},
		{
			name: "have A",
			args: args{[]rune("a")},
			want: true,
		},
		{
			name: "have Z",
			args: args{[]rune("a")},
			want: true,
		},
		{
			name: "have Symbol",
			args: args{[]rune("!!??")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HaveSingalAtoZandSymbolWord(tt.args.x); got != tt.want {
				t.Errorf("pickZHAndENWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortWordCloudRank(t *testing.T) {
	type args struct {
		UUID      string
		wordcloud map[string]int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SortWordCloudRank(tt.args.UUID, tt.args.wordcloud)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortWordCloudRank() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SortWordCloudRank() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
