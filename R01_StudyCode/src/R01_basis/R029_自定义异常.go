package main

import (
	"errors"
	"fmt"
)

func main029() {
	test029()
}

func test029() {
	err := readConf("config.ini1")
	if err != nil {
		// 如果读取文件发生错误,就输出这个错误,并终止程序
		panic(err)
	}
	fmt.Printf("test029()继续执行")
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取...
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误")
	}
}
