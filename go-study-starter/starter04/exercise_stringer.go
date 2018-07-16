package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

/*
练习：Stringers
	让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址。
	例如，`IPAddr{1,`2,`3,`4}` 应当输出 `"1.2.3.4"`。
*/
func (ipaddr IPAddr) String() string {
	return strconv.Itoa(int(ipaddr[0])) + "." + strconv.Itoa(int(ipaddr[1])) + "." +
		strconv.Itoa(int(ipaddr[2])) + "." + strconv.Itoa(int(ipaddr[3]))
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for key, a := range addrs {
		fmt.Printf("%v : %v \n", key, a)
	}

}

/*
   #string到int
   int,err:=strconv.Atoi(string)
   #string到int64
   int64, err := strconv.ParseInt(string, 10, 64)
   #int到string
   string:=strconv.Itoa(int)
   #int64到string
   string:=strconv.FormatInt(int64,10)
*/
