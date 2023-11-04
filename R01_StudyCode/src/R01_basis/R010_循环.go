package main

import "fmt"

func main010() {

	// 循环第一种
	for i := 1; i <= 10; i++ {
		fmt.Printf("你好1")
	}

	// 循环第二种
	j := 1
	for j <= 10 {
		fmt.Println("你好2")
		j++
	}

	// 循环第三种
	k := 1
	for { // 等价于 for ;;
		if k <= 10 {
			fmt.Printf("你好3")
		} else {
			break
		}
		k++
	}

	// 字符串变量方式
	var str string = "hello world!" //包含中文会乱码
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i]) //使用到下标
	}

	// 字符串遍历方式2-for-range
	for index, val := range str {
		fmt.Printf("index=%d , val=%c \n", index, val)
	}
	//index=11 , val=!
	//index=12 , val=上
	//index=15 , val=海

	var str2 string = "hello world!北京"
	str3 := []rune(str2) // 把string转成切片
	for i := 0; i < len(str3); i++ {
		fmt.Printf("index:%d, %c \n", i, str3[i]) //使用到下标
	}
	//index:11, !
	//index:12, 北
	//index:13, 京

}
