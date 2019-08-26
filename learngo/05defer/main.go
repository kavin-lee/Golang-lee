package main

import "fmt"

// defer
// 延时执行函数,先注册的后执行
// 也就是说,在函数即将要结束的时候,但是还没有结束的时候执行
// 作用及用途:一般用来做一些资源回收方面的工作,
func f1() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println("hello 沙河")
}
func main() {
	f1()
}
