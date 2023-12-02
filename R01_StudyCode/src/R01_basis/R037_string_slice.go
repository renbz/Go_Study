package main

import "fmt"

func main037() {
	str := "hello@12345"
	slice := str[6:]
	fmt.Println("slice = ", slice) // slice =  12345

	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Printf("str = ", str) // str = %!(EXTRA string=zello@12345)

	// 细节: 我们转成byte后,可以处理英文和数字,但是不能处理中文
	// 原因: []byte字节来处理,而一个汉字是3个字节,因此会出现乱码
	// 解决方法: 将string转成[]rune即可, 因为 []rune是按字符存储
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str = ", str) // str =  北ello@12345
}
