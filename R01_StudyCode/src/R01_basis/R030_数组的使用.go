package main

import "fmt"

func main030() {
	var numsArray01 [3]int = [3]int{1, 2, 3}
	var numsArray02 = [3]int{1, 2, 3}
	var numsArray03 = [...]int{1, 2, 3}
	// 可以指定元素值对应下标
	var names = [3]string{1: "tom", 0: "jack", 2: "marry"}
	fmt.Printf("numsArray01=%v,numsArray02=%v,numsArray03=%v,names=%v", numsArray01, numsArray02, numsArray03, names)

	// 从终端循环输入5个成绩,保存到float64数组,并输出
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输出第%d个元素的值\n", i+1)
		//fmt.Scanln(&score[i])
	}
	// 变量数组打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}

	for index, value := range score {
		fmt.Printf("下标:%d, value:%v\n", index, value)
	}
}
