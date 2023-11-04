package main

import (
	"fmt"
)

func main001() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int64 = int64(i)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Printf("i type is %T\n", i)
	fmt.Printf("n1 type is %T", n1)

}
