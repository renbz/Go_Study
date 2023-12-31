package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

// 通过反射，修改
// num  int 的值
// 修改 student 的值

func reflect01(b interface{}) {
	// 获取到reflect.Valu
	rVal := reflect.ValueOf(b)
	// num
	rVal.Elem().SetInt(20)

}

func main068() {
	conn, err := redis.Dial("tcp", "82.157.50.241:6379")
	if err != nil {
		fmt.Println("redis.Dial err = ", err)
		return
	}
	defer conn.Close()
	fmt.Println("conn suc...", conn)
}
