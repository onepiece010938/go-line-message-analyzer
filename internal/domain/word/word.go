package word

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/allegro/bigcache/v3"
	"github.com/bytedance/sonic"
	"github.com/go-ego/gse"
)

type Wordcloud struct {
	wordCount map[string]int
	//
}

func WordCould(path string, cache *bigcache.BigCache, segmenter *gse.Segmenter) Wordcloud {
	wordCountTmp := make(map[string]int)
	// senderMap map
	senderMap := map[string]int{}
	sendTimeMap := map[int]int{}
	for hour := 0; hour < 24; hour++ {
		sendTimeMap[hour] = 0
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// Assuming the path is a UUID
	err = getCacheWordCloud(path, &wordCountTmp, cache)
	if err == nil {
		// get data ,directly return
		return Wordcloud{wordCount: wordCountTmp}
	}
	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), "\t")
		if len(line) > 2 {
			timeString, sender, content := line[0], line[1], line[2]
			/*
				don't use Yoda condition error : go-staticcheck
				if "上午"==timeString[0:6] {
					fmt.Println(timeString)
				} else if "下午" == timeString[0:6] {
					fmt.Println(timeString)
				}
				一個中文字佔3個長度
			*/
			if timeString[0:6] == "上午" || timeString[0:6] == "AM" {
				timeString = timeString[6:] + " AM"
			} else if timeString[0:6] == "下午" || timeString[0:6] == "PM" {
				timeString = timeString[6:] + " PM"
			}
			// timeString
			const timeLayout = "03:04 PM"
			timeStr, err := time.Parse(timeLayout, timeString)
			if err != nil {
				fmt.Println("ERR", err)
			}
			sendTimeMap[timeStr.Hour()] = sendTimeMap[timeStr.Hour()] + 1
			if sender != "" {
				senderMap[sender] = senderMap[sender] + 1
			}

			words := segmenter.Slice(content)
			wordCount(&wordCountTmp, words)
		}
	}
	setCacheWordCloud(path, &wordCountTmp, cache)

	return Wordcloud{wordCount: wordCountTmp}
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
func isDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

func getCacheWordCloud(UUID string, result *map[string]int, cache *bigcache.BigCache) error {
	prefixKey := UUID + ":wordCloud"
	CacheValue, err := cache.Get(prefixKey)
	if err != nil {
		return err
	}
	err = sonic.Unmarshal(CacheValue, &result)
	if err != nil {
		fmt.Println()
		return err
	}
	return err
}
func setCacheWordCloud(UUID string, Infomap *map[string]int, cache *bigcache.BigCache) error {
	prefixKey := UUID + ":wordCloud"
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
