package analyze

import (
	"os"
	"testing"

	"github.com/go-ego/gse"
)

func TestWord(t *testing.T) {
	path := "D:/go-line-message-analyzer/cdcd.txt"
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
	segmentor.LoadDict("D:/Users/2327/go/pkg/mod/github.com/go-ego/gse@v0.80.2/data/dict/en/dict.txt")
	// f.Close()
	wordcloud := SplitWordToCloud("aaa", f, segmentor)
	SortWordCloudRank("aaa", wordcloud)
}

func BenchmarkSortWordCloudRank(b *testing.B) {
	path := "D:/go-line-message-analyzer/cdcd.txt"
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
	wordcloud := SplitWordToCloud("aaa", f, segmentor)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortWordCloudRank("aaa", wordcloud)
	}
	b.StopTimer()

}
