package main

import "fmt"

func main018() {
	a := getSum
	fmt.Printf("a的类型%T, getSum类型是%T\n", a, getSum)
	res := a(10, 40) // 等价 res:= getSum(10,40)
	fmt.Println("res=", res)

	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2) // 110
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型,因此在Go中,函数可以作为形参,并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func call(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
func main018_2() {
	res1, _ := call(10, 20)
	fmt.Printf("res1=%d", res1)
}

// 支持0到多个参数
func sum1(n1 int, args ...int) int {
	sum := n1
	//变量args
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
