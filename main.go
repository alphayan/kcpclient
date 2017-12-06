package main

import (
	"fmt"

	"github.com/xtaci/kcp-go"
	"strconv"
)

func main() {
	// kcp
	conn, err := kcp.DialWithOptions(REMOTE_ADDR, nil, 0, 0)
	if err != nil {
		fmt.Println("line 41 connection error: ", err)
		return
	}
	fmt.Println("connect success")
	conn.SetNoDelay(1, 10, 2, 1)
	conn.SetWindowSize(32, 32)
	conn.SetACKNoDelay(true)
	conn.SetMtu(512)
	//return
	var i =0
	for{
		i++
		_, err = conn.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			fmt.Println("write error")
			return
		}

		b:=make([]byte,4096)
		n, err := conn.Read(b)
		if err != nil {
			fmt.Println("write error")
			return
		}
		fmt.Println(string(b[:n]))
	}

}
