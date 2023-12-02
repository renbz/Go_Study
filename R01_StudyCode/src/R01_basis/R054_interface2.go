package main

import "fmt"

type AInterface interface {
	Say()
}

type Student54 struct {
}

func (s Student54) Say() {
	fmt.Println("student Say i = ", 111)
}

type integer1 int

func (i integer1) Say() {
	fmt.Println("integer1 Say i = ", i)
}

func main054() {
	var i integer1 = 10
	var b AInterface = i
	b.Say() // integer1 Say i =  10

	var stu Student54
	var a AInterface = stu
	a.Say() // student Say i =  111

}
