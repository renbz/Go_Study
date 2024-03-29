package main

import "fmt"

func main() {

	// 方式1
	// 创建结构体变量时,就直接指定字段的值
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明~", 20}

	// 在创建结构体变量时,把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
	var stu3 = Stu{
		Name: "jack",
		Age:  20,
	}
	stu4 := Stu{
		Age:  30,
		Name: "jack",
	}
	fmt.Println(stu1, stu2, stu3, stu4) // {小明 19} {小明~ 20} {jack 20} {jack 30}

	//方式2，返回结构体的指针类型(!!!)
	var stu5 *Stu = &Stu{"小王", 29} // stu5--> 地址 ---》结构体数据[xxxx,xxx]
	stu6 := &Stu{"小王~", 39}
	//在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
	var stu7 = &Stu{
		Name: "小李",
		Age:  49,
	}
	stu8 := &Stu{
		Age:  59,
		Name: "小李~",
	}
	fmt.Println(stu5, stu6, stu7, stu8)     // &{小王 29} &{小王~ 39} &{小李 49} &{小李~ 59}
	fmt.Println(*stu5, *stu6, *stu7, *stu8) // {小王 29} {小王~ 39} {小李 49} {小李~ 59}
}

type Stu struct {
	Name string
	Age  int
}
