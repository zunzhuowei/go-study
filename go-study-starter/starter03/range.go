package main

import "fmt"

var (
	pow = [] int{1,2,4,8,16,32,64, 128}

)

func main() {
		/*
		range
			for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。
	 */
	for i,v := range pow {
		fmt.Printf("2**%d  =  %d\n",i, v)
	}


/*

range（续）
	可以通过赋值给_来忽略序号和值。
	如果只需要索引值，去掉“, value”的部分即可。
 */

 	pow1 := make([]int , 10)

	for i := range pow1  {
		pow1[i] = 1 << uint(i)

	}
	for _,value := range pow1{
		fmt.Printf("%d\n", value)

	}
	
}
