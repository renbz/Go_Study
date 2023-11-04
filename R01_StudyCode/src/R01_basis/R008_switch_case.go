package main

import "fmt"

func main008() {

	var n1 int32 = 20
	var n2 int32 = 20
	switch n1 {
	case n2, 5, 10:
		fmt.Println("ok1")
	default:
		fmt.Println("默认")
	}

	// switch穿透
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok10")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("ok20")
	case 30:
		fmt.Println("ok30")
	default:
		fmt.Println("没有匹配到")

	}
}
