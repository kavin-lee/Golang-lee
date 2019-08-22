// server2 简易web服务器
//简易的统计一下网站的访问次数
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))

}

//处理handler的函数
func handler(w http.ResponseWriter, r *http.Request) {
	//加锁处理,避免同一时刻数据被不同的的请求修改
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)

}

//处理counter的函数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
