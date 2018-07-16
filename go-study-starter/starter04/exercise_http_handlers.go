package main

import (
	"net/http"
	"log"
)

type String string

type Struct struct {
	Geeting string
	Punct   string
	Who     string
}

type Handler struct {
	String
	Struct
}

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Handle("/string", String("I'm a frayed knot."))
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
}

func main() {

	//var stri String
	//var stru Struct
	var handler Handler

	err := http.ListenAndServe("localhost:4000", handler)
	if err != nil {
		log.Fatal(err)
	}

}

/*

练习：HTTP 处理

实现下面的类型，并在其上定义 ServeHTTP 方法。在 web 服务器中注册它们来处理指定的路径。

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

例如，可以使用如下方式注册处理方法：

http.Handle("/string", String("I'm a frayed knot."))
http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})


 */