package main

import "fmt"

func main002() {

	var n1 int32 = 12
	var n3 int8
	var n4 int8
	n4 = int8(n1) + 127 // 结果不是127+12
	//n3 = int8(n1) + 128 // 【编译不通过】
	fmt.Println(n3)
	fmt.Println(n4)
}
