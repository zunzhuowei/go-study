package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorites number is :", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Nextafter(2.1, 3.5))

}
