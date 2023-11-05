package main

import "fmt"

func main015() {
	fmt.Println(getNum(5))
}

func getNum(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*getNum(n-1) + 1
	}
}
