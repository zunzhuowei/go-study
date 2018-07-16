package main

import "fmt"

func swap(x, y string) (string, string, string) {
	return x, y, x + y
}

/*
多个返回值
*/
func main() {
	a, b, c := swap("10", "20")
	fmt.Println(a, b, c)
}
