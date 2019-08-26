package main

import "fmt"

//方法:就是某个具体的类型才可以调用的函数
//函数:是谁都可以调用的
type people struct {
	name   string
	gender string
}

//函数指定接受者之后就是方法
//在go中约定成俗不用this也不用self,而是使用后面的类型的首字母的小写
func (p people) dream() {
	fmt.Printf("%s的梦想是赚钱!!!!", p.name)
}
func main() {
	var kavin = people{
		name:   "kavin",
		gender: "男",
	}
	kavin.dream()
}
