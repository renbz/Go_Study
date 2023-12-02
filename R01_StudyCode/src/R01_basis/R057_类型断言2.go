package main

import "fmt"

type USB interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让phone实现USB接口
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Call() {
	fmt.Println("手机 在打电话...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
	Name string
}

// 让 camera 实现USB接口
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

// 计算机
type Computer struct {
}

// 编写一个方法Working方法, 接收一个USB接口与类型的变量
// 只要是实现了USB接口 (所谓实现USB接口,就是指实现了USB接口声明的所有方法)
func (c Computer) Working(usb USB) {
	usb.Start()
	// 如果USB是指向phone结构体变量,则还需要调用call方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}
func main057() {
	// 定义一个USB接口数组,可以存放phone和camera的结构体变量
	// 这里就体现出多态数组
	var usbArr [3]USB
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"索尼"}
	// 遍历usbArr
	// phone还有一个特有的方法call(), 请遍历USB数组,如果是phone变量
	// 除了调用USB接口声明的方法外,还需要调用phone特有方法 call -> 断言类型
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
