package main

import "fmt"

func main038() {

	// 声明二维数组
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v] = %v  ", i, j, v2)
		}
		fmt.Println()
	}

	/**
	1 2 3
	4 5 6
	arr[0][0] = 1  arr[0][1] = 2  arr[0][2] = 3
	arr[1][0] = 4  arr[1][1] = 5  arr[1][2] = 6
	*/
}
