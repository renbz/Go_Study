package main

import (
	"fmt"
)

func main011() {

	c := 5
	for i := 0; i < c; i++ {
		//打印空格 c-i-1个
		for j := 0; j < c-i-1; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("*")
		for j := 0; j < 2*i-1; j++ {
			if i == c-1 {
				fmt.Print("*")
			} else {
				fmt.Printf(" ")
			}
		}
		if i > 0 {
			fmt.Print("*")
		}
	}
}
