package main

import "fmt"

func main01() {
	sparseArray()
}

func sparseArray() {

	// 1. 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	// 2. 输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组
	// (1): 遍历chessMap, 如果我们发现有一个元素的值不为0, 创建一个node结构体
	// (2): 将其放入到对应的切片即可

	var sparseArr []ValNode
	// 标准的一个稀疏数组还应该有一个记录二维数组的规模(行和列,默认值)
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个valNode值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	fmt.Println("=============稀疏数组==========")

	// 输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d \n", i, valNode.row, valNode.col, valNode.val)
	}

	//将这个稀疏数组，存盘 d:/chessmap.data

	//如何恢复原始的数组

	//1. 打开这个d:/chessmap.data => 恢复原始数组.

	//2. 这里使用稀疏数组恢复

	// 先创建一个原始数组
	var chessMap2 [11][11]int

	// 遍历 sparseArr [遍历文件每一行]
	for i, valNode := range sparseArr {
		if i != 0 { //跳过第一行记录值
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	fmt.Println("=============稀疏数组复原==========")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}

type ValNode struct {
	row int
	col int
	val int
}
