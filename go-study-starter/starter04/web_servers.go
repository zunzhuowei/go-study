package main

import (
	"net/http"
	"log"
	"fmt"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}

}

/*

Web 服务器

包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

在这个例子中，类型 Hello 实现了 `http.Handler`。

访问 http://localhost:4000/ 会看到来自程序的问候。

 */