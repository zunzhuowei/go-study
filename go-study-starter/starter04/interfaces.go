package main

import (
	"math"
	"fmt"
)

/*
接口
	接口类型是由一组方法定义的集合。
	接口类型的值可以存放实现这些方法的任何值。
	注意： 列子代码的 22 行存在一个错误。 由于 Abs 只定义在 *Vertex（指针类型） 上，
	所以 Vertex（值类型） 不满足 `Abser`。
 */

//1) 定义接口
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())
}

//2）接口实现类1
type MyFloat float64

//2.1）接口实现类1,实现方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

//3）接口实现类2
type Vertex struct {
	X, Y float64
}

//3.1）接口实现类2,实现方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
