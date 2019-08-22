// 搭建一个建议的web服务器,
// 页面显示访问的url路径
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) //没一个请求需要有一个handler来进行处理
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))

}

//根据每一个不同的url,分别进行url的函数处理
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
