package main

import "fmt"

func main034() {

	// 演示切片的使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	// 声明/定义一个切片.
	// intArr[1:3] 标识slice引用到intArr这个数组的下标为1 到下标为3(不包含3)
	slice := intArr[1:3]
	slice[1] = 0

	fmt.Println("intArr=", intArr)            // intArr= [1 22 0 66 99]
	fmt.Println("slice 的元素是 =", slice)        // slice 的元素是 = [22 0]
	fmt.Println("slice 的元素个数是 =", len(slice)) // slice 的元素个数是 = 2
	fmt.Println("slice 的容量是 =", cap(slice))   // slice 的容量是 = 4

	var slice1 []float64 = make([]float64, 5, 10)
	slice1[1] = 10
	slice1[3] = 30
	fmt.Println(slice1)
	fmt.Println("slice1 的大小:", len(slice1)) // slice1 的大小: 5
	fmt.Println("slice1 的容量:", cap(slice1)) // slice1 的容量: 10

	var slice2 []string = []string{"tom", "jack", "mary"}
	fmt.Println(slice2)
	fmt.Println("slice2 的大小:", len(slice2)) // slice2 的大小: 3
	fmt.Println("slice2 的容量:", cap(slice2)) // slice2 的容量: 3

}
