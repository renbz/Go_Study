package main

import (
	"fmt"
	"go/types"
)

type Point struct {
	x int
	y int
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300
	TypeJudge(n1, n2, n3, name, address, n4)
	/**
	  第0个param is a float32 值为:1.1
	  第1个param is a float64 值为:2.3
	  第2个param is a 整数 值为:30
	  第3个param is a string 值为:tom
	  第4个param is a string 值为:北京
	  第5个param is a 整数 值为:300
	*/

}

func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%d个param is a bool 值为:%v\n", i, x)
		case float32:
			fmt.Printf("第%d个param is a float32 值为:%v\n", i, x)
		case float64:
			fmt.Printf("第%d个param is a float64 值为:%v\n", i, x)
		case int, int32, int64:
			fmt.Printf("第%d个param is a 整数 值为:%v\n", i, x)
		case types.Nil:
			fmt.Printf("第%d个param is a Nil 值为:%v\n", i, x)
		case string:
			fmt.Printf("第%d个param is a string 值为:%v\n", i, x)
		default:
			fmt.Printf("第%d个param is other 值为:%v\n", i, x)
		}
	}
}
