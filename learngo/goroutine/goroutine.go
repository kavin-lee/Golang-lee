//go语言的并发执行
package main

import (
	"fmt"
	"time"
)

//非抢占式的多任务
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { //数据访问冲突
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
				//a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
