package main

import (
	"fmt"
	"net"
)

func main069() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}

	// 功能一: 客户端可以发送单行数据，然后退出 

	fmt.Printf("conn = %v 成功", conn)

}
