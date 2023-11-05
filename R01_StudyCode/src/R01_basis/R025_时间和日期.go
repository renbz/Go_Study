package main

import (
	"fmt"
	"time"
)

func main025() {

	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	// 2. 通过now获取到年月日,时分秒
	fmt.Printf("年=%v\n", now.Year())       //年=2023
	fmt.Printf("月=%v\n", now.Month())      //月=November
	fmt.Printf("月=%v\n", int(now.Month())) //月=11
	fmt.Printf("日=%v\n", now.Day())        //日=5
	fmt.Printf("时=%v\n", now.Hour())       //时=17
	fmt.Printf("分=%v\n", now.Minute())     //分=47
	fmt.Printf("秒=%v\n", now.Second())     //秒=24

	// 3.格式化日期时间
	// 方式一: 使用printf和Sprintf
	fmt.Printf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dataStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("data:%s", dataStr)
	// 方式二: 使用time.Format()方法完成
	fmt.Printf(now.Format("2006-01-02 15:04:05")) // 2023-11-5 18:10:30data:2023-11-5 18:10:302023-11-05 18:10:30
}
