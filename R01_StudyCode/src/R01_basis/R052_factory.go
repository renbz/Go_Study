package main

import (
	"R01_StudyCode/src/R01_basis/mode"
	"fmt"
)

func main052() {
	var stu = mode.NewStudent("tom~", 99.9)
	fmt.Println(*stu)                                        // {tom~ 99.9}
	fmt.Println("name=", stu.Name, "score=", stu.GetScore()) // name= tom~ score= 99.9
}

type Goods struct {
	Name  string
	Price int
}

type Book struct {
	Goods  // 这里就是嵌套匿名结构体Goods
	Writer string
	// 匿名字段基本数据类型. 如果一个结构体有int类型的匿名字段,不能再有第二个
	// 如果需要有多个int字段, 则必须给int字段指定名字
	int
	n int
}
