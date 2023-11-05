package main

import "fmt"

func main027() {
	num1 := 100
	// num1的类型int, num1的值:100, num1的地址0xc00009a1a0
	fmt.Printf("num1的类型%T, num1的值:%v, num1的地址%v\n", num1, num1, &num1)

	num2 := new(int)
	// num2的类型*int, num2的值:0xc00000a1f8, num2的地址0xc00004c058
	fmt.Printf("num2的类型%T, num2的值:%v, num2的地址%v", num2, num2, &num2)
}
