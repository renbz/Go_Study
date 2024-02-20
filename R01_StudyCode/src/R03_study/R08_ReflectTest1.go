package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 请编写一个案例。
	// 演示对(基本数据类型、interface{}、reflect.Value) 进行反射的基本操作

	// 先定义一个int
	var num int = 100
	reflectTest(num)
}

func reflectTest(b interface{}) {

	// 通过反射获取传入的变量的 type、kind、值
	// 1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType = ", rTyp) // rType =  int

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v  type = %T", rVal, rVal) // rVal = 100  type = reflect.Value

	// 3. 类型转换  rVal.Int()
	n2 := 2 + rVal.Int()
	fmt.Println(n2) // 102
	// 将rVal转成 interface{}
	iV := rVal.Interface()
	// 将interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2 = %v  type = %T", num2, num2) // num2 = 100  type = int

}
