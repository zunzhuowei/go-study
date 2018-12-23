package main

import "fmt"

/*
变量
var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。

就像在这个例子中看到的一样，`var` 语句可以定义在包或函数级别。
*/

/*
短声明变量
在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。

函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
*/
var c, python, java bool

var i, j int = 1, 2

func main() {
	var i int

	fmt.Println(i, c, python, java)

	var c, java, python = 1, "yes", 1.2
	golang := "adfads"

	fmt.Println(c, java, python, golang)

}
