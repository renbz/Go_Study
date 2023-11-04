package main

import (
	"fmt"
	"strconv"
)

func main003() {

	var num1 int = 99
	var num2 float32 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string
	// 使用第一中方式转换 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %v\n", str, str) //str type string str = 99

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %v\n", str, str) //str type string str = 23.455999

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %v\n", str, str) //str type string str = true

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str = %v\n", str, str) //str type string str = h

	// 第二种方式 strconv 函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="99"

	// strconv.FormatFloat(num4, 'f', 10, 64)
	// 说明: 'f' 格式 10:表示小数保留10位  64:表示小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="23.4560000000"

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="true"

	// strconv包中有一个函数Itoa
	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str) //str type string str="4567"

}
