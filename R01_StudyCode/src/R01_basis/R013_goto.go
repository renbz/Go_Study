package main

import (
	"fmt"
	"math/rand"
)

func main013() {

	// 演示goto使用
	fmt.Println("ok1")
	if rand.Intn(10) > 5 {
		goto lable12
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
lable12:
	fmt.Println("ok5")

}
