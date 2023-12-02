package main

import "fmt"

func main046() {
	var a A
	var b B
	a = A(b) // 可以转换,但是要求:就是结构体的字段要完全一样(名字 个数 类型)
	fmt.Println(a, b)
}

type A struct {
	Num int
}
type B struct {
	Num int
}
