package main

import "fmt"

//结构体的匿名字段
type student struct {
	name string
	string
	int
}

func main() {
	var stu1 = student{
		name:   "kavin",
		string: "hah",
		int:    45,
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.string)
	fmt.Println(stu1.int)
}
