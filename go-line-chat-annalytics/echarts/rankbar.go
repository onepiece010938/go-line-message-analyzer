package echarts

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func rankByValue(wordCountMap map[string]int) (sortKeys []string) {
	log.Println("wordCountMap", wordCountMap)
	keys := make([]string, 0, len(wordCountMap))

	for key := range wordCountMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return wordCountMap[keys[i]] > wordCountMap[keys[j]]
	})
	return keys
}

func rankTopN(topNumber int, sortKeys []string, wordCountMap map[string]int) ([]string, []opts.BarData) {
	topWords := []string{}
	// topBarData := []int{}
	topBarData := make([]opts.BarData, 0)
	i := -1
	for len(topWords) <= topNumber {
		i += 1
		sortKey := sortKeys[i]
		count := wordCountMap[sortKey]
		if sortKey == "call" || sortKey == "photo" || sortKey == "vedio" || sortKey == "voice" || sortKey == "sticker" {
			continue
		}

		topWords = append(topWords, sortKey)
		topBarData = append(topBarData, opts.BarData{Value: count})
	}
	fmt.Println("topWords", topWords)
	fmt.Println("topBarData", topBarData)
	return topWords, topBarData
}

func rederRankBar(outpath string, topWords []string, topBarData []opts.BarData) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "常用詞Top10",
		Subtitle: "It's extremely easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis(topWords).AddSeries("Category A", topBarData)
	f, _ := os.Create(outpath + "/" + "rankBar.html")
	bar.Render(f)
}

func (echarts *Echarts) RankBar(topNumber int, wordCountMap map[string]int) {
	// rank
	sortKeys := rankByValue(wordCountMap)
	// TopNumber
	topWords, topBarData := rankTopN(topNumber, sortKeys, wordCountMap)

	outpath := echarts.outpath
	// topBarData: []opts.BarData
	rederRankBar(outpath, topWords, topBarData)
}
