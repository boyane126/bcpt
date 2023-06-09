# BCPT - 内容发布工具

BCPT = Boyane of content publishing tool

调用浏览器，发布小红书、抖音内容工具

> 通过命令行调用

## 使用方法

### 小红书
```shell
# 获取登录二维码
bcptctl xiaohongshu login

# 发布图文
bcptctl xiaohongshu pub_img_text

# 发布视频
bcptctl xiaohongshu pub_video 
```

### 获取登录二维码
```shell
# -f 平台 -s 登录二维码存储路径
bcptctl login -f xiaohongshu -s ./xiaohongshu_storeQr.png
```

### 发布视频

### 发布图文
