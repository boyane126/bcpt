package pub_img_text //nolint:typecheck

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/boyane126/bcpt/internal/pkg/browser"
	"github.com/boyane126/bcpt/internal/pkg/browser/xiaohongshu"
	"github.com/boyane126/bcpt/pkg/util"
)

var (
	burl                = browser.BURL("https://creator.xiaohongshu.com/publish/publish")
	allowImagesSuffixes = []string{ // 发布图片允许后缀名
		"jpeg", "png", "jpg",
	}
)

func NewCmdPubImgText() *cobra.Command {
	command := &cobra.Command{
		Use:   "pubImgText",
		Short: "发布图文",
		Long:  "用于发布图文；传递标题，封面图地址，简介",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(Run(args))
		},
	}

	return command
}

func Run(args []string) error {
	if err := verify(args); err != nil {
		return err
	}

	title, cover, desc := args[0], args[1], args[2]
	// TODO 多图暂时放弃

	// 组装发布内容参数
	data := browser.ImgTextData{
		Title: title,
		Cover: browser.FileAddr(cover),
		Desc:  desc,
	}

	mar := xiaohongshu.NewXiaohongshu()
	if err := mar.PubImgText(burl, data); err != nil {
		return err
	}

	log.Println("发布成功")
	return nil
}

func verify(args []string) error {
	_, cover, _ := args[0], args[1], args[2]
	if len(cover) == 0 {
		return fmt.Errorf("封面图不能为空")
	}
	suf := util.GetFileSuffix(cover)
	if !util.HasImagesSuffixes(suf, allowImagesSuffixes) {
		return fmt.Errorf("文件后缀必须在以下之内 %v 当前文件后缀 %s", allowImagesSuffixes, suf)
	}
	if !util.HasFileExist(cover) {
		return fmt.Errorf("封面图文件未找到 cover = %s", cover)
	}

	return nil
}
