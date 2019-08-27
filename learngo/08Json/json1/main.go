package main

import (
	"encoding/json"
	"fmt"
)

//学生
//type Student struct {
//	ID     int
//	Gender string
//	Name   string
//}
//定义源信息
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

func main() {
	var stu1 = Student{
		ID:     1,
		Gender: "男",
		Name:   "kavin",
	}
	//序列化:把编程语言中的数据转换成 json 格式的字符串
	val, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("JSON格式化错误!")
		fmt.Println(err) //把[]byte转成字符串
	}
	fmt.Println(string(val))
	fmt.Printf("%#v", string(val))

	str := "{\"ID\":1,\"Gender\":\"男\",\"Name\":\"kavin\"}"

	//反序列化:把满足JSON格式的字符串转换成 当前编程语言里面的对象
	var stu2 = &Student{}
	json.Unmarshal([]byte(str), stu2)
	fmt.Println(stu2)
	fmt.Printf("%T\n", stu2)
}
