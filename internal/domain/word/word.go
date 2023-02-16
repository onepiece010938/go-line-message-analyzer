package word

type FilterCloud struct {
	Filter    []string
	WordCloud map[string]int
}

func GenerateFilterWordCloud(wordDict FilterCloud) (map[string]int, error) {
	returnWord := make(map[string]int)
	for key, value := range wordDict.WordCloud {
		returnWord[key] = value
	}
	for _, filter := range wordDict.Filter {
		delete(returnWord, filter)
	}
	return returnWord, nil
}
