package main

import "fmt"

func main047() {
	var stu1 Student047
	var stu2 Stu047
	stu2 = Stu047(stu1)
	fmt.Println(stu1, stu2) // { 0}  { 10}

	var i integer047 = 10
	var j = 20
	j = int(i)        // j=i 不可以
	fmt.Println(i, j) // 10   10
}

type Stu047 Student047

type Student047 struct {
	Name string
	Age  int
}

type integer047 int
