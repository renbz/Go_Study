package main

import "fmt"

func main049() {
	var i integer = 10
	i.change()
	fmt.Println(i) // 11

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(stu) // {tom 20}
	// 实现了String() 输出对应方法。 否则输出地址
	fmt.Println(&stu) // Name=[tom] Age=[20]
}
func (i *integer) change() {
	*i = *i + 1
}
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

type integer int

type Student struct {
	Name string
	Age  int
}
