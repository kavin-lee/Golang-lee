package main

import (
	"fmt"
	"time"
)

func TimestampDemo(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year() //年
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//内置的time包
func main() {
	//time.Time  struct
	now := time.Now()
	fmt.Printf("%#v\n", now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	tm := now.Unix() + 3600

	//时间的打印
	TimestampDemo(tm)
}
