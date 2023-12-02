package main

import (
	"fmt"
	"math/rand"
)

func main033() {

	var intArr [5]int
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
}
