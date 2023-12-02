package main

import "fmt"

func main032() {

	arr := [3]int{11, 22, 33}
	test032_1(arr)
	fmt.Println("main_1 arr=", arr) // main_1 arr= [11 22 33]
	test032_2(&arr)
	fmt.Println("main_2 arr=", arr) // main_2 arr= [88 22 33]

}

func test032_1(arr [3]int) {
	arr[0] = 88
}
func test032_2(arr *[3]int) {
	(*arr)[0] = 88
}
