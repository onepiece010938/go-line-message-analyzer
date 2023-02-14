package analyze

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/go-ego/gse"
)

func pickZHAndENWord(x []rune) bool {
	for _, voca := range x {
		if unicode.IsSymbol(voca) || unicode.IsPunct(voca) {
			return true
		}

		if (voca > 'a' && voca < 'z') || (voca > 'A' && voca < 'Z') {
			if len(x) == 1 {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
func SplitWordToCloud(UUID string, file *os.File, segmenter *gse.Segmenter) map[string]int {
	// segmenter *gse.Segmenter
	defer file.Close()

	WordCloudMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	// wordsSlice := []string{}
	for scanner.Scan() {

		line := strings.Split(scanner.Text(), "\t")
		// ppp := scanner.Text()
		// fmt.Println(ppp)
		if len(line) > 2 {
			_, _, content := line[0], line[1], line[2]
			// unicode.Is(unicode.Han)
			// words := segmenter.Slice(content)
			words := segmenter.Cut(content, false)
			// fmt.Println(hihi)
			for _, p := range words {
				WordCloudMap[p] = WordCloudMap[p] + 1
			}

		}
	}
	// clear single(1,s,...or ! A)  map

	for key, _ := range WordCloudMap {

		runeKey := []rune(key)
		// len(runeKey) == 1 || len(runeKey) == 0
		if len(runeKey) == 0 || pickZHAndENWord(runeKey) {
			delete(WordCloudMap, key)
		}
	}

	return WordCloudMap
	// delete(WordCloudMap, "貼圖")
	// delete(WordCloudMap, "照片")
	// fmt.Println(WordCloudMap)
	// fmt.Println(WordCloudMap["貼圖"])
	// func SpecialLetters(letter rune) (bool, []rune) {
	// 	if unicode.IsPunct(letter) || unicode.IsSymbol(letter) {
	// 		var chars []rune
	// 		chars = append(chars, '\\', letter)
	// 		return true, chars
	// 	}
	// 	return false, nil
	// }
}

func SortWordCloudRank(UUID string, wordcloud map[string]int) {

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
	for i := 0; i < len(keySort); i++ {
		fmt.Println(keySort[i], valueSort[i])
	}
	// fmt.Println(wordcloud,)
}
