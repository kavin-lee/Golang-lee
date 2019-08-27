package main

import "fmt"

//结构体内嵌模拟"继承"

type animal struct {
	name string
}

//定义一个动物会动的方法
func (a *animal) move() {
	fmt.Println("%s会动~", a.name)
}

//定义一个狗 结构体
type dog struct {
	feet int
	*animal
}

//定义了一个狗的方法
func (d *dog) wangwang() {
	fmt.Printf("%s 在叫: 汪汪汪~\n", d.name)
}
func main() {
	var a = dog{
		feet:   4,
		animal: &animal{name: "旺财"},
	}
	a.wangwang()
	a.move()
}
