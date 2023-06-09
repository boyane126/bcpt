package xiaohongshu

import (
	"github.com/chromedp/chromedp"

	"github.com/boyane126/bcpt/internal/pkg/browser"
)

func NewXiaohongshu() browser.AppContentMar {
	return Xiaohongshu{}
}

type Xiaohongshu struct {
	browser.BaseAppContentMar
}

func (x Xiaohongshu) Login(burl browser.BURL, storeQrPos string) error {
	tasks := chromedp.Tasks{
		// 加载cookies
		LoadCookies(),
		// 打开小红书登录界面
		chromedp.Navigate(string(burl)),
		// 判断一下是否已经登录
		CheckLoginStatus(),
		// 点击扫码登录按钮
		chromedp.Click(`#page > div > div.content > div.con > div.login-box-container > div > div > div > div > img`),
		// 获取二维码
		GetCode(storeQrPos),
		// 若二维码登录后，浏览器会自动跳转到用户信息页面
		SaveCookies(),
	}

	defer browser.ChromeCancel()

	if err := chromedp.Run(browser.ChromeCtx, tasks); err != nil {
		return err
	}

	return nil
}
