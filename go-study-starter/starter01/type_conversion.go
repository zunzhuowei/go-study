package main

import (
	"fmt"
	"math"
)

func main() {
	//var x, y int= 3,4
	var x, y = 3, 4
	//var f float64 = math.Sqrt(float64(x * x + y * y))
	f := math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
	z := int(f)
	fmt.Println(x, y, z)

}
