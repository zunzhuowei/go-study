package main

import (
	"fmt"
)

func main() {

	/*
	slice
		一个 slice 会指向一个序列的值，并且包含了长度信息。
		[]T 是一个元素类型为 T 的 slice。
	 */
	p := [] int{2,3,5,7,11, 13}

	fmt.Println("p====", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n",i,p[i])
	}

	fmt.Println("-------------------------------------------------")

	/*
	对 slice 切片
		slice 可以重新切片，创建一个新的 slice 值指向相同的数组。
		表达式
		s[lo:hi]
		表示从 lo 到 hi-1 的 slice 元素，含两端。因此
		s[lo:lo]
		是空的，而
		s[lo:lo+1]
		有一个元素。
	 */
	pp := [] int {2,3,5,6,77, 11}
	fmt.Println("pp ====== ",pp)
	fmt.Println("pp[1,4] === ",pp[1:4])
	fmt.Println("pp[:3] === ",pp[:3])
	fmt.Println("pp[4:] === ",pp[4:])


}
