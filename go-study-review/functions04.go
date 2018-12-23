package main

import (
	"fmt"
	"strconv"
)

/*
函数可以没有参数或接受多个参数。

在这个例子中，`add` 接受两个 int 类型的参数。

注意类型在变量名 _之后_。

（参考 这篇关于 Go 语法定义的文章了解类型以这种形式出现的原因。）
*/
func add(x int, y int) int {
	z := x + y
	return z
}

func add2(x, y int) int {
	return x + y

}

func swap(x, y int) (string, string) {
	//strconv.FormatInt(int64(y),10)
	return strconv.Itoa(y), strconv.Itoa(x)

}

func split(x int) (y, z int) {
	y = x / 2
	z = x % 2
	return
}

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(add2(1, 3))
	fmt.Println(swap(1, 3))
	a, b := swap(1, 3)
	fmt.Println(a, b)

	fmt.Println(split(13))
}
