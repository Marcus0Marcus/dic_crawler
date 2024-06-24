package main

import (
	"context"
	"dic_crawler/common/logwrapper"
	"dic_crawler/common/traceid"
	"dic_crawler/ctrl/chromedppool"
	"fmt"
	"github.com/chromedp/chromedp"
)

//	func GetWordsByWord(ctx context.Context, word string) ([]string, error) {
//		var words []string
//		pageContent, err := chromedp.GetPageContentByLink(ctx, fmt.Sprintf("https://dictionary.cambridge.org/dictionary/english-chinese-traditional/%s", word))
//		if err != nil {
//			return words, err
//		}
//		words, err = htmlparse.GetAllWordsByHtml(pageContent)
//		if err != nil {
//			return words, err
//		}
//		return words, err
//	}
func main() {
	logwrapper.Init("./log", 2)
	defer logwrapper.Flush()
	//startWord := "word"
	ctx := traceid.WithTraceID(context.Background(), traceid.NewTraceID())
	//words, err := GetWordsByWord(ctx, startWord)
	//if err != nil {
	//	logwrapper.Fatal(ctx, err)
	//}
	//for _, word := range words {
	//	tWords, err := GetWordsByWord(ctx, word)
	//	if err != nil {
	//		logwrapper.Fatal(ctx, tWords)
	//	}
	//	fmt.Println("from word:", word, "got words len:", len(words), "words", words)
	//}
	pool, err := chromedppool.NewChromedpPool(5)
	if err != nil {
		logwrapper.Fatal(ctx, err)
	}
	defer pool.Shutdown()

	// 获取一个空闲的 ChromedpInstance
	instance, err := pool.GetInstance()
	if err != nil {
		logwrapper.Fatal(ctx, err)
	}

	// 运行任务
	var res string
	tasks := chromedp.Tasks{
		chromedp.Navigate("https://dictionary.cambridge.org/dictionary/english-chinese-traditional/word"),
		chromedp.Text("body", &res, chromedp.ByQuery),
	}

	err = pool.RunChromedpTask(instance, tasks)
	if err != nil {
		logwrapper.Fatal(ctx, err)
	}

	// 打印结果
	fmt.Println("Page content:", res)

	// 释放实例
	pool.ReleaseInstance(instance)
}
