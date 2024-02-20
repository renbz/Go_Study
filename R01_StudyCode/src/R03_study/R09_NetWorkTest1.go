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
		go process(conn)
	}

}

func process(conn net.Conn) {
	// 这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		// 等待客户端通过conn发送消息;如果客户端没有write[发送],那么携程就阻塞在这里
		//fmt.Println("服务器在等待客户端发送信息")
		n, err := conn.Read(buf) // 从conn读取
		if err != nil {
			fmt.Printf("客户端退出 err = %v", err)
			return
		}
		//3. 显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:n]))
	}
}
