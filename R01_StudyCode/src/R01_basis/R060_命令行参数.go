package main

import (
	"flag"
	"fmt"
	"os"
)

func main060() {
	fmt.Println("命令行参数有", len(os.Args))
	// 遍历 os.Args切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

}

func flagTest() {
	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int
	//&user 就是接收用户命今行中输入的 -u 后面的参数值
	//"u",就是 -u 指定参数/nn默认值
	//"用户名,默认为空"说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
	//这里有一个非常重要的操作,转换，必须调用该方法
	flag.Parse()

	// test.exe  -u root -pwd 123456 -h 127.0.0.1 -port 3306
}
