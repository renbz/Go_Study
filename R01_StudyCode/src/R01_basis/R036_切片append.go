package main

import "fmt"

func main036() {

	var slice = []int{100, 200, 300}
	// 通过append直接给slice3追加具体元素
	slice1 := append(slice, 400, 500, 600)
	slice2 := append(slice, slice...)
	slice2[0] = -1
	fmt.Printf("slice ", slice)   // slice %!(EXTRA []int=[100 200 300])
	fmt.Printf("slice1 ", slice1) // slice1 %!(EXTRA []int=[100 200 300 400 500 600])
	fmt.Printf("slice2 ", slice2) // slice2 %!(EXTRA []int=[-1 200 300 100 200 300])

	var a []int = []int{1, 2, 3, 4, 5}
	var slice3 = make([]int, 10)
	fmt.Println(slice3) // [0 0 0 0 0 0 0 0 0 0]
	a[0] = 111
	copy(slice3, a)
	fmt.Println(slice3) // [1 2 3 4 5 0 0 0 0 0]
}
