package main

import "fmt"

func main014() {
	res := fbn(3)
	fmt.Println(res)
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
