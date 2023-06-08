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
	ChromeCtx, ChromeCancel = chromedp.NewContext(context.Background())
}
