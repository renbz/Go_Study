package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main024() {
	// 1. 统计字符串长度,按字节len(str) [utf-8 汉字占3个字节]
	str := "hello北"
	fmt.Println("str len=", len(str)) // str len= 8

	// 2. 字符串遍历,同时处理有中文的问题 r:=[]rune(str)  切片
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 3. 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果:", n)
	}

	// 4. 整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T", str, str)

	// 5. 字符串转byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// 6. byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	// 7. 10进制转2、8、16进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制=%v\n", str)

	// 8. 查找子串是否在指定的字符串
	b := strings.Contains("seafood", "food")
	fmt.Printf("b=%v\n", b)

	// 9. 统计一个字符串有几个指定的子串
	count := strings.Count("abcabccbaabc", "abc")
	fmt.Printf("count=%v\n", count)

	// 10. 不区分大小写的字符串比较 (==是区分比较)
	fold := strings.EqualFold("abc", "ABC")
	fmt.Printf("fold=%v\n", fold)
	fmt.Println("结果:", "abc" == "ABC")

	// 11. 返回子串第一次出现的index值
	index := strings.Index("NLT_ABC", "ABC")
	fmt.Printf("index=%v", index)

}
