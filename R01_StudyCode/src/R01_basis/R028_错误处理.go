package main

import "fmt"

func main028() {
	test028()
	println("执行完main方法")
}

func test028() {
	// 使用defer+recover来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数,可以捕获到异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	num3 := num1 / num2
	fmt.Printf("num3=", num3)
}
