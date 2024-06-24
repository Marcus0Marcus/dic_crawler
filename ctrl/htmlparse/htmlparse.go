package htmlparse

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

func getWordByText(text string) []string {
	wordRegex := regexp.MustCompile(`\b[a-zA-Z]+\b`)

	// 查找所有匹配的单词
	words := wordRegex.FindAllString(text, -1)

	// 使用 map 保存单词并去重
	wordMap := make(map[string]bool)
	for _, word := range words {
		// 转换为小写，避免大小写不同的单词重复
		lowercaseWord := strings.ToLower(word)
		wordMap[lowercaseWord] = true
	}

	// 提取去重后的单词列表
	var uniqueWords []string
	for word := range wordMap {
		uniqueWords = append(uniqueWords, word)
	}
	return uniqueWords
}
func GetAllWordsByHtml(pageContent string) ([]string, error) {
	var allWords []string
	allTextMp := make(map[string]bool)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageContent))
	if err != nil {
		return allWords, err
	}
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		text := strings.ToLower(strings.TrimSpace(s.Text()))
		if text != "" {
			words := getWordByText(text)
			for _, word := range words {
				if _, ok := allTextMp[word]; !ok {
					allTextMp[word] = true
					allWords = append(allWords, word)
				}
			}

		}
	})

	return allWords, nil
}
