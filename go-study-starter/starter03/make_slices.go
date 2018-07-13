package main

import "fmt"

func printSlice(s string, x [] int) {
	fmt.Printf("%s len = %d cap = %d %v\n",s,len(x),cap(x), x)

}

/*
 slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：
	a := make([]int, 5)  // len(a)=5
	为了指定容量，可传递第三个参数到 `make`：
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
 */
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	fmt.Println("---------------")

	b := make([]int,0, 5)
	printSlice("b", b)

	fmt.Println("---------------")
	c := b[:2]
	printSlice("c", c)

	fmt.Println("---------------")
	d := c[2:5]
	printSlice("d", d)

	fmt.Println("+++++++++++++++++++++++")

	/*
	nil slice
		slice 的零值是 `nil`。
		一个 nil 的 slice 的长度和容量是 0。
	 */
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

}
