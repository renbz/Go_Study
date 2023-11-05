package main

import "fmt"

func main016() {
	fmt.Println(peach(1))
}

func peach(n int) int {
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
