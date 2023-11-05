package main

import "fmt"

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main021_2() {
	res := Fun1(4, 5)
	fmt.Println("res=", res)
}

func main021() {

	// 在定义匿名函数时就直接调用,这种方式匿名函数只能调用一次

	// 案例演示,求两个数的和。使用匿名函数的方式完成
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	// 将匿名函数func(n1 int,n2 int) int赋给a变量
	// 则a的数据类型就是函数类型,此时 我们可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(20, 30)
	fmt.Println("res2=", res2)
}
