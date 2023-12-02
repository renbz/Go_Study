package main

import "fmt"

func main045() {

	// 第一种方式
	var p1 Person
	fmt.Println(p1)

	// 第二种方式
	p2 := Person{}
	p2.Name = "tom"
	fmt.Println(p2)

	// 第三种方式
	var p3 *Person = new(Person)
	// 因为p3是一个指针,因此标准的给字段赋值方式
	(*p3).Name = "smith" // 为了书写方便也可以写成  p3.Name = "smith" ; 底层会处理成*p3
	(*p3).Age = 30
	fmt.Println(*p3)

	//第四种方式
	var p4 *Person = &Person{}
	// 因为person是一个指针,因此标准的访问字段方法
	(*p4).Name = "scott" // 为了书写方便也可以写成  p4.Name = "scott" ; 底层会处理成*p4.Name =
	fmt.Println(*p4)

}

type Person struct {
	Name string
	Age  int
}
