# BCPT - 内容发布工具

BCPT = Boyane of content publishing tool

调用浏览器，发布小红书、抖音内容工具

> 通过命令行调用

## 使用方法

### 小红书
```shell
# 获取登录二维码
bcptctl xiaohongshu login

# 发布图文 ./bcptctl xiaohongshu pubImgText 标题 /home/rsh/Code/go/src/bcpt/cmd/bcptctl/212957.jpeg 简介
bcptctl xiaohongshu pub_img_text

# 发布视频 ./bcptctl xiaohongshu pubVideo test视频 /home/rsh/Code/go/src/bcpt/cmd/bcptctl/test.mp4 /home/rsh/Code/go/src/bcpt/cmd/bcptctl/212957.jpeg test简介
bcptctl xiaohongshu pub_video 
```

## 作者

- [@boyane126](https://github.com/boyane126)
- 邮箱: 2628488871@qq.com