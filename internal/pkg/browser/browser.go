// 浏览器

package browser

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

var (
	ChromeCtx    context.Context
	ChromeCancel context.CancelFunc
)

func init() {
	ChromeCtx = context.Background()
	//ChromeCtx, ChromeCancel = chromedp.NewExecAllocator(ChromeCtx) // debug模式
	//ChromeCtx, ChromeCancel = chromedp.NewRemoteAllocator(ChromeCtx, "ws://127.0.0.1:9222") // 无头浏览器
	//ChromeCtx, ChromeCancel = chromedp.NewContext(ChromeCtx)

	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用false  正式使用用true
		chromedp.WindowSize(1920, 1024), // 调整浏览器大小
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	ChromeCtx, ChromeCancel = chromedp.NewExecAllocator(context.Background(), options...)
	ChromeCtx, ChromeCancel = chromedp.NewContext(ChromeCtx, chromedp.WithLogf(log.Printf)) // 会打开浏览器并且新建一个标签页进行操作

	//ChromeCtx, ChromeCancel = chromedp.NewRemoteAllocator(ChromeCtx, "ws://127.0.0.1:9222") //使用远程调试，可以结合下面的容器使用
	//ChromeCtx, ChromeCancel = chromedp.NewContext(ChromeCtx)                                // WithTargetID可以指定一个标签页进行操作

}
