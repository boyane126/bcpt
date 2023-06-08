package cmd

import (
	"fmt"
	"github.com/boyane126/bcpt/internal/bcptctl/cmd/util"
	"github.com/spf13/cobra"
)

var (
	platform string
)

func NewCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "获取登录二维码",
		Long:  "用于获取系统支持内容发布平台的登录二维码，注意有效时间",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(Run(args))
		},
	}

	cmd.Flags().StringVarP(&platform, "platform", "p", "", "platform 请使用 [xiaohongshu]")

	return cmd
}

func Run(args []string) error {
	fmt.Println(args)

	return nil
}
