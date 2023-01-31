package echarts

type Echarts struct {
	outpath string
}

func NewEcharts(outpath string) *Echarts {
	echarts := &Echarts{
		outpath: outpath,
	}
	return echarts
}
