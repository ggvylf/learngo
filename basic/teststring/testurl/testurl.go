package main

import (
	"fmt"
	"strings"
)

//判断前缀
func urlProcess(url string) string {
	result := strings.HasPrefix(url, "https")
	if !result {
		url = fmt.Sprintf("https://%s", url)
	}
	return url
}

//判断后缀
func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func main() {
	var (
		url  string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Println(url)
	fmt.Println(path)
}
