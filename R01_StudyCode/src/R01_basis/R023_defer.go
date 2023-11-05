package main

import "fmt"

var Name string = "tom"

func main023() {
	res := sum(10, 20)
	fmt.Println("res=", res)
	fmt.Println("res=", Name)
	/**
	ok3 res= 30
	ok2 n2= 20
	ok1 n1= 10
	res= 30    */
}
func sum(n1 int, n2 int) int {
	// 当执行到defer时, 暂时不执行会将defer后面的语句压入到独立的栈(defer栈)
	// 当函数执行完毕后,再从defer栈,按照先入后出的当时出栈执行
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}
