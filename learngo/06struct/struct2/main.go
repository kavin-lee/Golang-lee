package main

import "fmt"

//结构体
//创建新的类型需要使用type关键字

type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//结构体实例化方法1
	var kavin = student{
		name:   "kavin",
		age:    24,
		gender: "男",
		hobby:  []string{"篮球", "摄影", "游泳"},
	}
	//结构体访问属性
	fmt.Println(kavin)
	fmt.Println(kavin.name)
	fmt.Println(kavin.age)
	fmt.Println(kavin.gender)
	fmt.Println(kavin.hobby)

	//结构体实例化方法1
	//struct 是值类型
	//如果初始化时没有给属性(字段)设置对应的初始值,那么对应属性就是
	//其类型的默认值
	var allen = student{}
	fmt.Println(allen.name)
	fmt.Println(allen.age)
	fmt.Println(allen.gender)
	fmt.Println(allen.hobby)

	//实例化方法2
	var tom = new(student)
	tom.name = "tom"
	tom.age = 18
	tom.gender = "男"
	fmt.Println(tom.name, tom.age, tom.gender)
}
