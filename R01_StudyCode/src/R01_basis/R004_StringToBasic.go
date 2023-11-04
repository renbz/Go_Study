package main

import (
	"fmt"
	"strconv"
)

func main004() {

	var str string = "true"
	var b bool
	// 说明: 1. 这个函数会返回两个值(value bool, err error) 2. 只想获取value bool,不想获取err 所以使用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "11"
	var n1 int64
	// 字符串 进制  转整数的类型
	n1, _ = strconv.ParseInt(str2, 2, 64)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)

	var str3 string = "123.4567"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)

	// 如果没有转成功 默认0
	var str4 string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T n3=%v\n", n3, n3) //n3 type int64 n3=0

}
