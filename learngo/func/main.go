package main

//函数
import "fmt"

//不带参数的返回值
func f1() {
	fmt.Println("hello 沙河!")
}

//两个参数的返回值
func f2(name1, name2 string) {
	fmt.Println(name1)
	fmt.Println(name2)
}

//可变参数,参数可为 0 或者 多个
func f3(names ...string) {
	fmt.Println(names) //[]string
}

//无返回值
func f4() {
	fmt.Println("哈哈 正那")
}

//一个返回值
func f5(a, b int) int {
	return a + b
}

//多返回值,必须要用括号扩起来,用英文逗号分割
func f6(a, b int) (int, int) {
	return a + b, a - b
}

//带返回名的函数
func f7(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

//匿名函数
var f = func(name string) {
	fmt.Println("hello", name)
}

//闭包:函数调用了它外层的变量
func closure(key string) func(string) {
	//key := "沙河有沙子!"
	return func(name string) {
		fmt.Println("hello", name)
		fmt.Println(key)
	}
}
func main() {
	//f("豪杰")
	//func(name string) {
	//	fmt.Println("hekllo", name)
	//}("zhuzhuz")
	//闭包
	m := closure("沙河有沙子!")
	fmt.Printf("%T", m)

	//调用闭包函数
	m("哈接")

}
