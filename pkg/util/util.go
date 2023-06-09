package util

import "fmt"

func CheckErr(err error) {
	fmt.Println(err)
}

func HasImagesSuffixes(fileSuf string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if suffix == fileSuf {
			return true
		}
	}

	return false
}
