package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/boyane126/bcpt/internal/bcptctl/cmd/util"
)

var (
	platform   string
	storeQrPos string
)

func NewCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "获取内容发布平台登录二维码",
		Long:  "用于获取系统支持内容发布平台的登录二维码，注意有效时间",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(Run(args))
		},
	}

	cmd.Flags().StringVarP(&platform, "platform", "p", "", "请使用 [xiaohongshu]")
	cmd.Flags().StringVarP(&storeQrPos, "storeQrPos", "s", "", "存储登录二维码位置，默认存放在当前文件夹./xxxx_loginQRCode.png")

	return cmd
}

func Run(args []string) error {

	mar, err := NewServer(Platform(platform))
	if err != nil {
		return err
	}
	if len(storeQrPos) > 0 {
		mar.storeQr = storeQrPos
	}

	defer func() {
		if err := os.RemoveAll(mar.storeQr); err != nil {
			log.Println("删除二维码失败")
		} else {
			log.Println("删除二维码成功")
		}
	}()

	if err = mar.app.Login(mar.loginUrl, mar.storeQr); err != nil {
		return err
	}

	log.Println("存储位置：", mar.storeQr)

	return nil
}
