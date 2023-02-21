package word

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/onepiece010938/go-line-message-analyzer/internal/domain/word"
)

type FilterWordParm struct {
	// CreateHistogramParams postgres.CreateHistogramParams
	WordCloud   *map[string]int
	UserName    string
	FilterWords []string
}

func (i *WordService) WordCloud(ctx context.Context, WordParm FilterWordParm) (*map[string]int, error) {

	var value *map[string]int
	err := i.wordServiceCache.GetWordCloudCache(WordParm.UserName, value)
	if err != nil {
		return nil, err
	}
	return value, err
}
func (i *WordService) FilterCloud(ctx context.Context, WordParm FilterWordParm) (*map[string]int, error) {
	var value *map[string]int

	filterWords := WordParm.UserName
	for _, filter := range WordParm.FilterWords {
		filterWords = filterWords + filter
	}
	filterCloudKey := getSHA256HashCode(filterWords)

	err := i.wordServiceCache.GetFilterCloudCache(filterCloudKey, value)
	if err == nil {
		return value, err
	}

	Cloud := word.GenerateFilterWordCloud(word.FilterCloud{Filter: WordParm.FilterWords, WordCloud: *WordParm.WordCloud})

	err = i.wordServiceCache.SetFilterCloudCache(filterCloudKey, &Cloud)
	if err != nil {
		return nil, err
	}

	return &Cloud, err
}
func (i *WordService) StringRank(ctx context.Context, WordParm FilterWordParm) (*[]string, error) {
	var value *[]string
	err := i.wordServiceCache.GetStringRankCache(WordParm.UserName, value)
	if err != nil {
		return nil, err
	}
	return value, err
}
func (i *WordService) AmountRank(ctx context.Context, WordParm FilterWordParm) (*[]int, error) {
	var value *[]int
	err := i.wordServiceCache.GetAmountRankCache(WordParm.UserName, value)
	if err != nil {
		return nil, err
	}
	return value, err
}

func getSHA256HashCode(stringMessage string) string {

	message := []byte(stringMessage) //字符串转化字节数组
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New() //sha-256加密
	//hash := sha512.New() //SHA-512加密
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

}
