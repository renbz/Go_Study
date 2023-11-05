package main

import (
	"fmt"
	"strconv"
	"time"
)

func main026() {
	/*for i := 0; i <= 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}*/

	//unix和unixNano的使用
	now := time.Now()
	fmt.Printf("unix时间戳:%v  unixnano时间戳:%v\n", now.Unix(), now.UnixMilli())
	//unix时间戳:1699180013  unixnano时间戳:1699180013634

	currTime1 := time.Now().Unix()
	test03()
	currTime2 := time.Now().Unix()
	fmt.Printf("执行test03()耗费时间为%v秒", currTime2-currTime1)
}
func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
