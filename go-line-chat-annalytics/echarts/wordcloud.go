package echarts

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateWCData(wordCountMap map[string]int) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range wordCountMap {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func wcBase(wordCountMap map[string]int) *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "basic WordCloud example",
		}))

	wc.AddSeries("wordcloud", generateWCData(wordCountMap)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
				}),
		)
	return wc
}

func wcCardioid(wordCountMap map[string]int) *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "cardioid shape"}),
	)

	wc.AddSeries("wordcloud", generateWCData(wordCountMap)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "cardioid",
				}),
		)
	return wc
}

func wcStar(wordCountMap map[string]int) *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "star shape",
		}))

	wc.AddSeries("wordcloud", generateWCData(wordCountMap)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "star",
				}),
		)
	return wc
}

func wcTriangle(wordCountMap map[string]int) *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "triangle shape",
		}))

	wc.AddSeries("wordcloud", generateWCData(wordCountMap)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "triangle",
				}),
		)
	return wc
}

func (echarts *Echarts) Wordcloud(wordCountMap map[string]int) {
	page := components.NewPage()
	// Optional: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow"
	page.AddCharts(
		wcBase(wordCountMap),
		wcCardioid(wordCountMap),
		wcStar(wordCountMap),
		wcTriangle(wordCountMap),
	)

	f, err := os.Create(echarts.outpath + "/" + "wordcloud.html")

	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
