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
	ChromeCtx = context.Background()
	//ChromeCtx, ChromeCancel = chromedp.NewExecAllocator(ChromeCtx) // debug模式
	ChromeCtx, ChromeCancel = chromedp.NewContext(ChromeCtx)
}
