package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"unicode"

	"github.com/go-ego/gse"
	chart "github.com/go-line-chat-annalytics/echarts"
)

func NewGse() (*gse.Segmenter, error) {
	segmentor := &gse.Segmenter{
		AlphaNum: true,
	}
	err := segmentor.LoadDict()
	if err != nil {
		log.Println("fail to load default dict", err)
	}
	err = segmentor.LoadDict("customized_dict.txt")
	if err != nil {
		log.Println("fail to load customized_dict", err)
	}
	return segmentor, nil
}

func main() {
	// 加载默认词典
	segmentor, err := NewGse()
	if err != nil {
		log.Fatal("cannot create Segmenter:", err)
	}

	// 分詞
	source := "../source/line_jill.txt"
	_, _, wordCountMap := wordSegment(source, segmentor)
	// sendTimeMap, senderMap, wordCountMap := word(source)
	// TODO: sendTimeMap, senderMap

	// echart
	outpath := "./chartResults"
	echart := chart.NewEcharts(outpath)

	// rank
	topNumber := 10
	echart.RankBar(topNumber, wordCountMap)

	// wordcloud
	echart.Wordcloud(wordCountMap)
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

func wordSegment(path string, segmenter *gse.Segmenter) (map[int]int, map[string]int, map[string]int) {
	// sendTimeMap map
	sendTimeMap := map[int]int{}
	for hour := 0; hour < 24; hour++ {
		sendTimeMap[hour] = 0
	}

	// senderMap map
	senderMap := map[string]int{}

	// content map
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
			timeString, sender, content := line[0], line[1], line[2]
			// timeString
			const timeLayout = "03:04 PM"
			timeStr, err := time.Parse(timeLayout, timeString)
			if err != nil {
				fmt.Println("ERR", err)
			}
			sendTimeMap[timeStr.Hour()] = sendTimeMap[timeStr.Hour()] + 1

			// senderMap
			if sender != "" {
				senderMap[sender] = senderMap[sender] + 1
			}
			// wordCountMap
			words := segmenter.Slice(content)
			wordCount(&wordCountMap, words)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sendTimeMap, senderMap, wordCountMap

}
