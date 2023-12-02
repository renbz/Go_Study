package main

import "fmt"

func main035() {
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}
	slice := intArr[1:4]
	slice2 := slice[1:2]

	slice[1] = -1
	for _, v := range slice2 {
		fmt.Println(v)
	}

	// 常规遍历
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v   ", i, slice[i]) // slice[0] = 22   slice[1] = 33   slice[2] = 44
	}

	// for-range遍历
	for i, v := range slice {
		fmt.Printf("i=%v v=%v   ", i, v) // i=0 v=22   i=1 v=33   i=2 v=44
	}
}
