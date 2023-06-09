package util

import "fmt"

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func HasFileSuffixes(fileSuf string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if suffix == fileSuf {
			return true
		}
	}

	return false
}
