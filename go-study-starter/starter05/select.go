package main

import (
	"fmt"
	//"time"
)

func fibonacci(c, quit chan int) {
	//time.Sleep(time.Minute)

	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//支线程
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	/*go func() {
		time.Sleep(1000)
		quit <- 0
	}()*/

	//主线程
	fibonacci(c, quit)
}

/*
	1.当执行main方法的时候，可能先执行支线程的闭包函数，也有可能先执行主线程的fibonacci(c, quit);
	如果先执行支线程，
 */

/*
select
	select 语句使得一个 goroutine 在多个通讯操作上等待。
	select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。
 */
