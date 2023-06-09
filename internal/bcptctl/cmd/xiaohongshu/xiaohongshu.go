package xiaohongshu

import (
	"github.com/boyane126/bcpt/internal/bcptctl/cmd/xiaohongshu/pub_video"
	"github.com/spf13/cobra"

	"github.com/boyane126/bcpt/internal/bcptctl/cmd/xiaohongshu/login"
	"github.com/boyane126/bcpt/internal/bcptctl/cmd/xiaohongshu/pub_img_text"
)

func NewCmdXiaohongshu() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xiaohongshu",
		Short: "小红书平台内容管理",
		Long:  "小红书平台内容管理，用于操作小红书内容发布命令集合",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(
		login.NewCmdLogin(),
		pub_img_text.NewCmdPubImgText(),
		pub_video.NewCmdPubVideo(),
	)

	return cmd
}
