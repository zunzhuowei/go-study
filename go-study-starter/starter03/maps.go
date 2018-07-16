package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//key  value
var m map[string]Vertex

func main() {
	/*
		map
			map 映射键到值。
			map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
	*/
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	fmt.Println("---------------------------")

	/*
		map 的文法
			map 的文法跟结构体文法相似，不过必须有键名。
	*/
	var mm = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
		"Google1": { //如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。
			37.422021, -122.08408,
		},
	}

	fmt.Println(mm)

	fmt.Println("---------------------------")

	/*
	修改 map
		在 map m 中插入或修改一个元素：
		m[key] = elem
		获得元素：
		elem = m[key]
		删除元素：
		delete(m, key)
		通过双赋值检测某个键存在：
		elem, ok = m[key]
		如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
		同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	*/
	mmm := make(map[string]int)
	mmm["Answer"] = 42
	fmt.Println("The value:", mmm["Answer"])

	mmm["Answer"] = 48
	fmt.Println("The value:", mmm["Answer"])

	delete(mmm, "Answer")
	fmt.Println("The value:", mmm["Answer"])

	v, ok := mmm["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
