package xiaohongshu

import (
	"context"
	"time"

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

	ctx, cancel := context.WithTimeout(browser.ChromeCtx, time.Second*30)
	defer cancel()

	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}

	return nil
}

func (x Xiaohongshu) PubImgText(burl browser.BURL, imgTextData browser.ImgTextData) error {
	tasks := chromedp.Tasks{
		// 登录
		LoadCookies(),
		IsLoginStatus(),
		// 发布笔记
		chromedp.Navigate(string(burl)),
		chromedp.Click(`#publish-container > div > div.header > div:nth-child(2)`),
		chromedp.SendKeys(`#publisher-dom > div > div.publisher-container > div > div.upload-container > div.video-uploader-container.upload-area > div.upload-wrapper > div > input`, string(imgTextData.Cover), chromedp.NodeVisible),
		chromedp.SendKeys(`#publisher-dom > div > div.publisher-container > div > div.img-post > div.content > div.c-input.titleInput > input`, imgTextData.Title),
		AddContent(imgTextData.Desc),
	}

	ctx, cancel := context.WithTimeout(browser.ChromeCtx, time.Minute)
	defer cancel()

	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}

	return nil
}

func (x Xiaohongshu) PubVideo(burl browser.BURL, videoData browser.VideoData) error {
	tasks := chromedp.Tasks{
		// 登录
		LoadCookies(),
		//IsLoginStatus(),
		chromedp.Navigate(string(burl)),
		chromedp.SendKeys(`#publish-container > div > div.video-uploader-container.upload-area > div.upload-wrapper > div > input`, string(videoData.Video), chromedp.NodeVisible),
		chromedp.SendKeys(`#publish-container > div > div:nth-child(3) > div.content > div.c-input.titleInput > input`, videoData.Title),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 上传封面
			if len(videoData.Cover) > 0 {
				chromedp.Click(`#publish-container > div > div:nth-child(3) > div.content > div.cover-container > button`)
				chromedp.Click(`#cover-modal-0 > div > div > div.css-t0051x.css-y1z97h.dyn.content > div > div.css-cf5fey.tab-container.tab-position-top > div.css-ckmc4o.tab-headers.header-line > div:nth-child(2)`)
				chromedp.SendKeys(`#cover-modal-0 > div > div > div.css-t0051x.css-y1z97h.dyn.content > div > div.css-cf5fey.tab-container.tab-position-top > div.css-wzyxpg > div.upload-wrapper > input`, string(videoData.Cover), chromedp.NodeVisible)
				chromedp.Click(`#cover-modal-0 > div > div > div.css-8mz9r9.footer > div > button.css-k3hpu2.css-osq2ks.dyn.btn-confirm`)
			}
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			time.Sleep(20 * time.Second)
			return nil
		}),
		AddContent(videoData.Desc),
	}

	ctx, cancel := context.WithTimeout(browser.ChromeCtx, 2*time.Minute)
	defer cancel()

	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}

	return nil
}
