package main

import "fmt"

func main006() {

	var n1 float32 = 10 / 4
	fmt.Println(n1) // 2

	var n2 float32 = 10.0 / 4
	fmt.Println(n2) // 2.5

	n3 := '人'
	fmt.Printf("n3 type %T, %v", n3, n3) //n3 type int32, 20154

	a := 9
	b := 2
	fmt.Printf("变换前 a= %v, b = %v \n", a, b)
	t := a
	a = b
	b = t
	fmt.Printf("变换后 a= %v, b=%v \n", a, b)
}
