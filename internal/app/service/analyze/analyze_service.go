package analyze

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
)

type CreateAnalyzeParm struct {
	// CreateHistogramParams postgres.CreateHistogramParams
}

type AnalyzeMessageParm struct {
}

func (i *AnalyzeService) CreateAnalyze(ctx context.Context, param CreateAnalyzeParm) error {
	fmt.Println("AnalyzeService->func CreateAnalyze()")
	// var result message.MessageDomainResult
	// result = message.MessageDomainFunc("aabcccc")
	// fmt.Println(result)
	i.analyzeServiceCache.GetMessageCache("")

	return nil
}

func (i *AnalyzeService) AnalyzeTest(ctx context.Context) (string, error) {
	fmt.Println("AnalyzeService->func AnalyzeTest()")
	// var result message.MessageDomainResult
	// result = message.MessageDomainFunc("aabcccc")
	// fmt.Println(result)
	return "TESTTESTTEST", nil
}

func (i *AnalyzeService) StartAnalyze(context io.ReadCloser) (string, error) {
	var last string
	scanner := bufio.NewScanner(context)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "\t")
		// value := domain photo analyze(row)
		// if value!=nil{
		//		adapter cache photo insert
		//	}
		// value := domain message analyze(row)
		// if value!=nil{
		//		adapter cache message insert
		//	}
		fmt.Println(row)
		last = row[0]
	}

	defer context.Close()

	return last, nil
}
