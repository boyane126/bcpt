package util

import "testing"

func TestCreateFile(t *testing.T) {
	CookiesTmp := "./temp/aaa.txt"
	f, err := CreateFile(CookiesTmp)
	if err != nil {
		return
	}
	cookiesData := []byte("hello world")
	if _, err = f.Write(cookiesData); err != nil {
		return
	}
}
