package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听...")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close() // 延时关闭listen

	//循环等待客户端来链接我
	for {
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err = ", err)
		} else {
			fmt.Printf("accept success con = %v ; 客户端ip = %v \n", conn, conn.RemoteAddr())
		}
		// 这里准备起一个协程 为客户端服务
	}

}
