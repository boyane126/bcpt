package pub_video

import (
	"fmt"
	"github.com/boyane126/bcpt/internal/pkg/browser"
	"github.com/boyane126/bcpt/internal/pkg/browser/xiaohongshu"
	"github.com/boyane126/bcpt/pkg/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	burl                = browser.BURL("https://creator.xiaohongshu.com/publish/publish")
	allowImagesSuffixes = []string{ // 发布图片允许后缀名
		"jpeg", "png", "jpg",
	}
	allowVideoSuffixes = []string{
		"mp4",
	}
)

func NewCmdPubVideo() *cobra.Command {
	command := &cobra.Command{
		Use:   "pubVideo",
		Short: "发布视频",
		Long:  "用于发布视频；传递标题，视频地址，封面图地址，简介",
		Args:  cobra.MinimumNArgs(4),
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

	title, video, cover, desc := args[0], args[1], args[2], args[3]

	// 组装发布内容参数
	data := browser.VideoData{
		Title: title,
		Cover: browser.FileAddr(cover),
		Video: browser.FileAddr(video),
		Desc:  desc,
	}

	mar := xiaohongshu.NewXiaohongshu()
	if err := mar.PubVideo(burl, data); err != nil {
		return err
	}

	log.Println("发布成功")
	return nil
}

func verify(args []string) error {
	_, video, cover, _ := args[0], args[1], args[2], args[3]
	if len(video) == 0 {
		return fmt.Errorf("视频不能为空")
	}
	suf := util.GetFileSuffix(video)
	if !util.HasFileSuffixes(suf, allowVideoSuffixes) {
		return fmt.Errorf("文件后缀必须在以下之内 %v 当前文件后缀 %s", allowVideoSuffixes, suf)
	}
	if !util.HasFileExist(video) {
		return fmt.Errorf("视频文件未找到 video = %s", video)
	}

	if len(cover) > 0 {
		suf := util.GetFileSuffix(cover)
		if !util.HasFileSuffixes(suf, allowImagesSuffixes) {
			return fmt.Errorf("文件后缀必须在以下之内 %v 当前文件后缀 %s", allowImagesSuffixes, suf)
		}
		if !util.HasFileExist(cover) {
			return fmt.Errorf("封面图文件未找到 cover = %s", cover)
		}

	}

	return nil
}
