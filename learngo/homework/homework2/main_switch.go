package main

import (
	"fmt"
)

//分金币作业
/**
50枚金币:分配给几个人:matthew,sarah,augustus,heidi,emilie,peter,
giana,adriano,agfg
分配规则如下:
a.名字中包含e或者E:1枚金币
b.名字中包含i或者I:2枚金币
c.名字中包含o或者O:3枚金币
d.名字中包含u或者U:4枚金币
写一个程序,计算每个用户分到多少个金币,以及最后剩余多少金币?

**/

var (
	conis1        = 50
	users1        = []string{"matthew", "sarah", "augustus", "heidi", "emilie", "peter", "giana", "adriano", "agfg"}
	distribution1 = make(map[string]int, len(users1))
)

//根据规则分发金币,返回剩余的金币数量
func dispatchCoin1() int {
	sum := 0
	//1.遍历名字切片
	for _, name := range users1 {
		//根据规则分发金币
		for _, char := range name {
			switch char {
			case 'E', 'e':
				sum = sum + 1
				distribution1[name] = distribution1[name] + 1
			case 'i', 'I':
				sum = sum + 2
				distribution1[name] = distribution1[name] + 2
			case 'o', 'O':
				sum = sum + 3
				distribution1[name] = distribution1[name] + 3
			case 'u', 'U':
				sum = sum + 4
				distribution1[name] = distribution1[name] + 4
			}
		}
	}
	return conis1 - sum
}

func main() {
	left := dispatchCoin1()
	fmt.Println("剩余: ", left)
	for k, v := range distribution1 {
		fmt.Println(k, v)
	}
}
