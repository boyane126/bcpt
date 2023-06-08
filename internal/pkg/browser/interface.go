package browser

type BURL string
type FileAddr string

// ImgTextData 图文数据
type ImgTextData struct {
	Title  string
	Cover  FileAddr
	Images []FileAddr
	Desc   string
	Extra  interface{}
}

// VideoData 视频数据
type VideoData struct {
	Title string
	Cover FileAddr
	Video FileAddr
	Desc  string
	Extra interface{}
}

// AppContentMar App内容管理
type AppContentMar interface {
	// Login 登录
	Login(burl BURL, storeQrPos string) error

	// PubImgText 发布图文
	PubImgText(burl BURL, imgTextData ImgTextData) error

	// PubVideo 发布视频
	PubVideo(burl BURL, videoData VideoData) error
}

type BaseAppContentMar struct{}

func (m BaseAppContentMar) Login(burl BURL, storeQrPos string) error {
	return nil
}
func (m BaseAppContentMar) PubImgText(burl BURL, imgTextData ImgTextData) error {
	return nil
}
func (m BaseAppContentMar) PubVideo(burl BURL, videoData VideoData) error {
	return nil
}
