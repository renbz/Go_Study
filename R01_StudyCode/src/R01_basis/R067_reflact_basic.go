package main

import (
	"fmt"
	"reflect"
)

func main067() {

	const (
		a = 1
		b = 2
	)

	const (
		a0 = iota
		b1
		c2
	)

	// 请编写一个案例。
	// 演示对(基本数据类型、interface{}、reflect.Value) 进行反射的基本操作

	// 先定义一个int
	var num int = 100
	reflectTest(num)
	fmt.Println("\n=========================")

	stu := Student067{
		Name: "tom",
		Age:  10,
	}
	reflectTest2(stu)

}

func reflectTest(b interface{}) {

	// 通过反射获取传入的变量的 type、kind、值
	// 1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType = ", rTyp) // rType =  int

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 将rVal转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV = %v  iV = %T", iV, iV) // num2 = 100  type = int

	// 将interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2 = %v  type = %T", num2, num2) // num2 = 100  type = int

}

func reflectTest2(b interface{}) {

	// 通过反射获取传入的变量的 type、kind、值
	// 1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType = ", rTyp) // rType =  int

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 3. 获取 变量对应的Kind
	// (1) rVal.Kind() ==>
	Kind1 := rVal.Kind()
	// (2) rType.Kind() ==>
	Kind2 := rTyp.Kind()
	fmt.Printf("Kind =%v   Kind=%v \n", Kind1, Kind2) // Kind =main.Student067   Kind=main.Student067

	iV := rVal.Interface()
	fmt.Printf("iv=%v   iv Type=%T \n", iV, iV) // iv={tom 10}   iv Type=main.Student067
	stu, ok := iV.(Student067)
	if ok {
		fmt.Printf("stu.Name=%v \n", stu.Name) // stu.Name=tom
	}
}

type Student067 struct {
	Name string
	Age  int
}
