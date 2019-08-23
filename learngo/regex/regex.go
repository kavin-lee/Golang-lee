package main

import (
	"fmt"
	"regexp"
)

const text = `My email is kavin@yeah.net@yeah.net
email1 is absfds@def.ogt
email2 is ffff@qq.com
email3 is ffff@qq.com.cn
`

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match:=re.FindAllString(text,-1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
