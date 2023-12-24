package main

import (
	"fmt"
	"reflect"
)

// 通过反射，修改
// num  int 的值
// 修改 student 的值

func reflect01(b interface{}) {
	// 获取到reflect.Valu
	rVal := reflect.ValueOf(b)
	// num
	rVal.Elem().SetInt(20)

}

func main068() {

	var num int = 10
	reflect01(&num)
	fmt.Printf("num=", num) // num=%!(EXTRA int=20)

}
