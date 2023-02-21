package analyze

import (
	"bufio"
	"mime/multipart"
	"sort"
	"strings"
	"unicode"

	"github.com/go-ego/gse"
)

func HaveSingalAtoZandSymbolWord(x []rune) bool {
	for _, voca := range x {
		if unicode.IsSymbol(voca) || unicode.IsPunct(voca) {
			return true
		}

		if (voca >= 'a' && voca <= 'z') || (voca >= 'A' && voca <= 'Z') {
			if len(x) == 1 {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
func SplitWordToCloud(UUID string, file *multipart.File, segmenter *gse.Segmenter) map[string]int {

	WordCloudMap := make(map[string]int)
	scanner := bufio.NewScanner(*file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		if len(line) > 2 {
			_, _, content := line[0], line[1], line[2] //nolint

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

func SortWordCloudRank(UUID string, wordcloud map[string]int) ([]string, []int) {

	keySort := make([]string, 0, len(wordcloud))
	valueSort := make([]int, 0, len(wordcloud))

	for key := range wordcloud {
		keySort = append(keySort, key)
		valueSort = append(valueSort, wordcloud[key])
	}

	sort.SliceStable(keySort, func(i, j int) bool {
		return wordcloud[keySort[i]] > wordcloud[keySort[j]]
	})
	sort.SliceStable(valueSort, func(i, j int) bool {
		return valueSort[i] > valueSort[j]
	})
	return keySort, valueSort
}
