package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}


/*
函数定义
 */
func main() {
	println(add(10,20))
	fmt.Println("---------------")
	fmt.Println(add(10, 20))

	fmt.Println(add2(10, 45))

}
