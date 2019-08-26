package main

import "fmt"

//结构体的嵌套
type address struct {
	province string
	city     string
}
type student struct {
	name string
	age  int
	addr address //嵌套了别的结构体
}

func main() {
	var stu1 = student{
		name: "kavin",
		age:  15,
		addr: address{
			province: "河北",
			city:     "雄安",
		},
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.age)
	fmt.Println(stu1.addr.province)
	fmt.Println(stu1.addr.city)
}
