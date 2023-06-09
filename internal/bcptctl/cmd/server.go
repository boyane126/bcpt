package cmd

import (
	"fmt"
	"github.com/boyane126/bcpt/internal/pkg/browser"
	"github.com/boyane126/bcpt/internal/pkg/browser/xiaohongshu"
)

type Platform string

const (
	Xiaohongshu = Platform("xiaohongshu")
)

type Server struct {
	platform Platform
	loginUrl browser.BURL
	storeQr  string
	app      browser.AppContentMar
}

func NewServer(platform Platform) (s Server, err error) {
	switch platform {
	case Xiaohongshu:
		s.platform = Xiaohongshu
		s.loginUrl = "https://creator.xiaohongshu.com/login"
		s.storeQr = "./xiaohongshu_loginQRCode.png"
		s.app = xiaohongshu.NewXiaohongshu()
	default:
		return s, fmt.Errorf("平台不支持")
	}
	return s, nil
}
