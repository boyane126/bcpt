package xiaohongshu

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"

	"github.com/boyane126/bcpt/pkg/util"
)

const (
	CookiesTmp = "./tmp/xiaohongshu_cookies.tmp"
)

func LoadCookies() chromedp.ActionFunc {
	return func(ctx context.Context) error {
		// 如果cookies临时文件不存在则直接跳过
		if _, _err := os.Stat(CookiesTmp); os.IsNotExist(_err) {
			return nil
		}

		// 如果存在则读取cookies的数据
		cookiesData, err := ioutil.ReadFile(CookiesTmp)
		if err != nil {
			return err
		}

		cookiesParams := network.SetCookiesParams{}
		if err = cookiesParams.UnmarshalJSON(cookiesData); err != nil {
			return err
		}

		// 设置cookies
		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
	}
}

func CheckLoginStatus() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		var url string
		time.Sleep(2 * time.Second)
		if err = chromedp.Evaluate(`window.location.href`, &url).Do(ctx); err != nil {
			return
		}

		if strings.Contains(url, "https://creator.xiaohongshu.com/creator/home") {
			return fmt.Errorf("已经使用cookies登陆")
		}
		return
	}
}

func GetCode(storeQrPos string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		var code []byte

		// 等待一秒加载二维码
		time.Sleep(time.Second)

		// TODO 可能一秒够加载，需要继续等待

		// 截图
		if err := chromedp.Screenshot(`#page > div > div.content > div.con > div.login-box-container > div > div > div > div > div.css-dvxtzn`, &code, chromedp.ByID).Do(ctx); err != nil {
			return err
		}

		// 保存为图片文件
		if err := os.WriteFile(storeQrPos, code, 0o644); err != nil {
			log.Fatal(err)
		}

		return nil
	}
}

func SaveCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 等待二维码登陆
		if err = chromedp.WaitVisible(`#app`, chromedp.ByID).Do(ctx); err != nil {
			log.Println(err)
			return
		}

		// cookies的获取对应是在devTools的network面板中
		// 1. 获取cookies

		cookie, err := network.GetCookies().Do(ctx)
		if err != nil {
			return
		}

		// 2. 序列化
		cookiesData, err := network.GetCookiesReturns{Cookies: cookie}.MarshalJSON()
		if err != nil {
			return
		}

		// 3. 存储到临时文件
		f, err := util.CreateFile(CookiesTmp)
		if err != nil {
			return
		}
		if _, err = f.Write(cookiesData); err != nil {
			return
		}
		return
	}
}
