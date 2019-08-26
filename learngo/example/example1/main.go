package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		url  string
		path string
	)
	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(url)
	fmt.Println(url)
	fmt.Println(path)
}

func pathProcess(path string) string {
	//判断 path 是不是以/结尾
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func urlProcess(url string) string {
	//判断 url 是不是以http:// 开头
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}
