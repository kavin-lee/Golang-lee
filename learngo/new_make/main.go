package main

import "fmt"

func main() {
	//var a *int //a是一个int类型的指针
	var a = new(int) //得到一个int类型的指针
	//var b *string
	//var c *[3]int
	*a = 10
	fmt.Println(*a) //打印a指针对应的值

	var c = new([3]int)
	fmt.Println(c)
	//(*c)[0]=1
	c[0] = 1
	fmt.Println(*c)
}
