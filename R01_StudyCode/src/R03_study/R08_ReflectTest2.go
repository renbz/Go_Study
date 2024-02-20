package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 请编写一个案例。
	stu := Student067{
		Name: "tom",
		Age:  10,
	}
	reflectTest2(stu)

}

func reflectTest2(b interface{}) {

	// 通过反射获取传入的变量的 type、kind、值
	// 1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType = ", rTyp) // rType =  main.Student067

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 3. 获取 变量对应的Kind
	// (1) rVal.Kind() ==>
	Kind1 := rVal.Kind()
	// (2) rType.Kind() ==>
	Kind2 := rTyp.Kind()
	fmt.Printf("Kind =%v   Kind=%v \n", Kind1, Kind2) // Kind =struct   Kind=struct

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
