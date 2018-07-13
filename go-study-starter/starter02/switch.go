package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
switch
	一个结构体（`struct`）就是一个字段的集合。
	除非以 fallthrough 语句结束，否则分支会自动终止。
 */
func main() {

	fmt.Println("Go runs on")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("------------------------")

	/*

switch 的执行顺序
	switch 的条件从上到下的执行，当匹配成功的时候停止。
	（例如，
	switch i {
	case 0:
	case f():
	}
	当 i==0 时不会调用 `f`。）

	 */
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println("------------------------")

	/*

没有条件的 switch
	没有条件的 switch 同 `switch true` 一样。
	这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	 */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Godd afternoon")
	default:
		fmt.Println("Good evening")


	}

}
