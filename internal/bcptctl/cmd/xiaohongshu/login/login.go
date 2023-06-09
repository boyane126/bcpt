package login

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/boyane126/bcpt/internal/pkg/browser"
	"github.com/boyane126/bcpt/internal/pkg/browser/xiaohongshu"
	"github.com/boyane126/bcpt/pkg/util"
)

var (
	storeQrPos        string
	defaultStoreQrPos = "./xiaohongshu_loginQRCode.png"                       // 默认存储位置
	burl              = browser.BURL("https://creator.xiaohongshu.com/login") // 登录url
)

func NewCmdLogin() *cobra.Command {
	command := &cobra.Command{
		Use:   "login",
		Short: "获取内容发布平台登录二维码",
		Long:  "用于获取系统支持内容发布平台的登录二维码，注意有效时间",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(Run(args))
		},
	}

	command.Flags().StringVarP(&storeQrPos, "storeQrPos", "s", defaultStoreQrPos, "存储登录二维码位置，默认存放在当前文件夹./xxxx_loginQRCode.png")

	return command
}

func Run(args []string) error {
	mar := xiaohongshu.NewXiaohongshu()

	defer func() {
		if err := os.RemoveAll(storeQrPos); err != nil {
			log.Println("删除二维码失败")
		} else {
			log.Println("删除二维码成功")
		}
	}()

	if err := mar.Login(burl, storeQrPos); err != nil {
		return err
	}

	log.Println("存储位置：", storeQrPos)

	return nil
}
