package main

import "fmt"

type USB1 interface {
	Start()
	Stop()
}

type Phone1 struct {
	Name string
}

// 让phone实现USB接口
func (p Phone1) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone1) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera1 struct {
	Name string
}

// 让 camera 实现USB接口
func (c Camera1) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera1) Stop() {
	fmt.Println("相机停止工作...")
}

// 计算机
type Computer1 struct {
}

// 编写一个方法Working方法, 接收一个USB接口与类型的变量
// 只要是实现了USB接口 (所谓实现USB接口,就是指实现了USB接口声明的所有方法)
func (c Computer1) Working(usb USB) {
	usb.Start()
	usb.Stop()
}

func main053() {
	// 定义一个USB接口数组,可以存放phone和camera的结构体变量
	// 这里就体现出多态数组
	var usbArr [3]USB1
	usbArr[0] = Phone1{"vivo"}
	usbArr[1] = Phone1{"小米"}
	usbArr[2] = Camera1{"索尼"}

	// 测试
	// 先创建结构体变量
	computer := Computer1{}
	phone := Phone1{}

	camera := Camera1{}
	computer.Working(phone)
	computer.Working(camera)
}
