// 文件类工具函数

package util

import (
	"os"
	"path"
)

// CreateFile 创建文件 文件夹不存在就创建文件夹
func CreateFile(file string) (*os.File, error) {
	dir := path.Dir(file)

	if len(dir) > 0 {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModePerm)
}
