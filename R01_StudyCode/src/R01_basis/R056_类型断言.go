package main

import "fmt"

type Point struct {
	x int
	y int
}

func main056() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point // ok
	// 如何将a赋值给一个Point变量?
	var b Point
	// b = a 不可以
	// b = a.(point) // 可以
	b = a.(Point)
	fmt.Println(b)

	/*	// 类型断言的其他案例
		var x interface{}
		var b2 float32 = 1.1
		x = b2 // 空接口, 可以接收任意类型
		// x-> float32 [使用类型断言]
		y := x.(float32)
		fmt.Printf("y 的类型是 %T 值是=%v", y, y) // y 的类型是 float32 值是=1.1*/

	// 类型断言(带检测)
	var x interface{}
	var b2 float32 = 1.1
	x = b2 // 空接口, 可以接收任意类型
	// x-> float32 [使用类型断言]
	y, ok := x.(float32)
	if ok == true {
		fmt.Println("convert success \n")      // convert success
		fmt.Printf("y 的类型是 %T 值是=%v \n", y, y) // y 的类型是 float32 值是=1.1
	} else {
		fmt.Println("convert fail")
	}

	if yy, ok := x.(float32); ok {
		fmt.Println("convert success")         // convert success
		fmt.Printf("yy 的类型是 %T 值是=%v", yy, yy) // yy 的类型是 float32 值是=1.1
	} else {
		fmt.Println("convert fail")
	}

}
