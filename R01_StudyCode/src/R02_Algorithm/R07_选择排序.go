package main

import "fmt"

func main() {
	// 定义一个数组
	arr := [5]int{10, 34, 19, 100, 80}
	selectSort(&arr)
	fmt.Println("main函数")
	fmt.Println(arr)

}

func selectSort(arr *[5]int) {

	// 标准式的访问  (*arr)[1]=900 等价于 arr[1] = 900
	// 1. 假设 arr[0] 最大值

	for j := 0; j < len(arr)-1; j++ {

		max := arr[j]
		maxIndex := j
		// 2. 遍历后面 1 --> len(arr)-1
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		// 交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
}
