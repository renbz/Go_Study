package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔1秒输出 "hello,world"
// 在主线程中也每隔一秒输出"hello,golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

// 编写一个函数，每隔1秒输出 "hello,协程"
func test062() {
	for i := 1; i <= 10; i++ {
		fmt.Println("tesst () hello,协程 " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main062() {

	go test062() // 开启了一个协程 【非顺序执行】

	for i := 1; i <= 10; i++ {
		fmt.Println(" main() hello,主线程" + strconv.Itoa(i))
		time.Sleep(time.Second)

	}

	// 获取当前系统CPU数量
	num := runtime.NumCPU()
	fmt.Println("cpu数量:", num) // cpu数量: 20
	// 设置num-1的cpu运行go程序
	runtime.GOMAXPROCS(num - 1)
}
