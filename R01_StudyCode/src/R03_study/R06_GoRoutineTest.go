package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// 编写一个函数，每隔1秒输出 "hello,协程"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("tesst () hello,协程 " + strconv.Itoa(i))
		time.Sleep(2 * time.Second)
	}
}

func main() {

	// 获取当前系统CPU数量
	num := runtime.NumCPU()
	fmt.Println("cpu数量:", num) // cpu数量: 20
	// 设置num-1的cpu运行go程序
	runtime.GOMAXPROCS(num - 1)

	go test() // 开启了一个协程 【非顺序执行】

	for i := 1; i <= 10; i++ {
		fmt.Println(" main() hello,主线程" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
