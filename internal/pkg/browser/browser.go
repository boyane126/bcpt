// 浏览器

package browser

import (
	"context"
	"github.com/chromedp/chromedp"
)

var (
	ChromeCtx    context.Context
	ChromeCancel context.CancelFunc
)

func init() {
	ChromeCtx, ChromeCancel = chromedp.NewExecAllocator(context.Background())
	//ChromeCtx, ChromeCancel = chromedp.NewContext(ChromeCtx, chromedp.WithLogf(log.Printf)) // debug模式
}
