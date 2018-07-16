package main

import (
	"fmt"
	"math"
)

type MyFloat float64

/*
方法（续）
	你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
	但是，不能对来自其他包的类型或基础类型定义方法。
*/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) Abs2(x int, y float64) float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	ff := MyFloat.Abs(-math.Sqrt2)
	fmt.Println(ff)

	fff := f.Abs2(1, 1.11)
	fmt.Println(fff)

	ffff := MyFloat.Abs2(f, 1, 1.11)
	fmt.Println(ffff)

}
