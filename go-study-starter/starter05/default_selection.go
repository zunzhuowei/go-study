package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

/*
默认选择
	当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。
	为了非阻塞的发送或者接收，可使用 default 分支：

	select {
	case i := <-c:
		// 使用 i
	default:
		// 从 c 读取会阻塞
	}
 */
