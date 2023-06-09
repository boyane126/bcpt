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
	return nil
}
