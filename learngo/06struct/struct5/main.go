package main

import "fmt"

//自己实现一个构造函数
type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

//自己写一个构造函数
func newStudent(n string, age int, g string, h []string) student {
	return student{
		name:   n,
		age:    age,
		gender: g,
		hobby:  h,
	}
}
func main() {
	hobbySlice := []string{"游泳", "摄影"}
	kavin := newStudent("kavin", 18, "男", hobbySlice)
	fmt.Println(kavin.name)
	fmt.Println(kavin.age)
	fmt.Println(kavin.gender)
	fmt.Println(kavin.hobby)
}
