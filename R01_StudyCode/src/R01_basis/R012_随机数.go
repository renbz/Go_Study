package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main012() {

	n := rand.Intn(100)
	fmt.Println(n)

lable2:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable2
			}
		}
	}
}

func hIndex(citations []int) (h int) {
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return
}
