package main

import (
	"fmt"
	"strings"
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
	conis        = 50
	users        = []string{"matthew", "sarah", "augustus", "heidi", "emilie", "peter", "giana", "adriano", "agfg"}
	distribution = make(map[string]int, len(users))
)

//根据规则分发金币,返回剩余的金币数量
func dispatchCoin() int {
	sum := 0
	//1.遍历名字切片
	for _, name := range users {
		//根据规则分发金币
		if strings.Contains(name, "e") || strings.Contains(name, "E") {
			distribution[name] = distribution[name] + 1
			sum++
		}
		if strings.Contains(name, "i") || strings.Contains(name, "I") {
			distribution[name] = distribution[name] + 2
			sum = sum + 2
		}
		if strings.Contains(name, "o") || strings.Contains(name, "O") {
			distribution[name] = distribution[name] + 3
			sum = sum + 3
		}
		if strings.Contains(name, "u") || strings.Contains(name, "U") {
			distribution[name] = distribution[name] + 4
			sum = sum + 4
		}
	}
	return conis - sum
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩余: ", left)
	for k, v := range distribution {
		fmt.Println(k, v)
	}
}
