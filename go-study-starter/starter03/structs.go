package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	/*
	结构体
		一个结构体（`struct`）就是一个字段的集合。
		（而 type 的含义跟其字面意思相符。）
	 */
	fmt.Println(Vertex{1, 2})

	fmt.Println("--------------------")

	/*
	结构体字段
	结构体字段使用点号来访问。
	 */
	v := Vertex{3, 4}
	v.X = 4
	fmt.Println(v.X)


	/*
	结构体指针
		结构体字段可以通过结构体指针来访问。
		通过指针间接的访问是透明的。
	 */

	p := &v
	p.X = 1e9
	fmt.Println(v)


	/*
结构体文法
	结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。
	使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
	特殊的前缀 & 返回一个指向结构体的指针。
	 */
	var (
		v1 = Vertex2{1, 2} // 类型为 Vertex
		v2 = Vertex2{X: 1} // Y:0 被省略
		v3 = Vertex2{}	// X:0 和 Y:0
		pp = &Vertex2{1,2}	// 类型为 *Vertex
	)

	fmt.Println(v1,v2,v3, pp)


}

type Vertex2 struct {
	X, Y int
}