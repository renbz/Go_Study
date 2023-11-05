package main

import "fmt"

func main017() {
	num := 20
	test017(&num)
	fmt.Println(num) // 30
}

func test017(n *int) {
	*n = *n + 10
	fmt.Println(*n) // 30
}
