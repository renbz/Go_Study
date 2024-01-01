package main

import "fmt"

func main() {
	arr := [7]int{23, 0, 12, 56, 34, -1, 55}
	fmt.Println("原始数组=", arr)
	fmt.Println("main 函数")
	InsertSort(&arr)
	//fmt.Printf("插入排序耗时%d秒", end-start)
	fmt.Println(arr)
}

func InsertSort(arr *[7]int) {

	for i := 1; i < len(arr); i++ {
		// 给第i个元素找到合适的位置并插入
		insertVal := arr[i]
		insertIndex := i - 1

		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}

}
