package chromedp

import (
	"context"
	"dic_crawler/common/logwrapper"
	"github.com/chromedp/chromedp"
	"time"
)

func GetPageContentByLink(ctx context.Context, url string) (string, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // 禁用无头模式
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("start-maximized", true),
		chromedp.UserDataDir("./userdata"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(func(s string, i ...interface{}) {
			logwrapper.Info(ctx, s, i)
		}),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	var pageContent string
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`document.readyState`, &res), // 等待整个页面加载完成
		chromedp.OuterHTML("html", &pageContent),       // 获取整个网页的 HTML 内容
	)
	if err != nil {
		logwrapper.Error(ctx, err)
		return "", err
	}
	return pageContent, nil
}
