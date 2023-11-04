package main

import (
	"fmt"
	"go/types"
)

func main009() {

	var x interface{}
	var y = 10.0
	x = y
	// 可以判断x这个空接口 指向的数据类型
	switch i := x.(type) {
	case types.Nil:
		fmt.Printf("x 的类型~ :%T", i)
	case int:
		fmt.Printf("x 是int型")
	case float64:
		fmt.Printf("x 是float64型")
	case func(int) float64:
		fmt.Printf("x 是func(int)型")
	case bool, string:
		fmt.Printf("x 是bool或string型")
	default:
		fmt.Printf("未知型 %T", i)
	}

}
