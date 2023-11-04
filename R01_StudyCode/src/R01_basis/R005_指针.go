package main

import (
	"fmt"
)

func main005() {

	// 基本数据类型在内存布局
	var i int = 10
	// 获取i的地址 &i
	fmt.Println("i的地址=", &i) // i的地址= 0xc00000a0c8

	// 下面var ptr *int = &i 含义
	// 1. ptr是一个指针变量
	// 2. ptr的类型 *int
	// 3. ptr本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr=i的地址: %v\n", ptr) // ptr=i的地址: 0xc00000a0c8
	fmt.Printf("ptr的地址%v\n", &ptr)    // ptr的地址0xc00004a028
	fmt.Printf("ptr指向的值%v\n", *ptr)   // ptr的地址0xc00004a028
	//fmt.Println(mode.Hhh)             // ptr的地址0xc00004a028
}
