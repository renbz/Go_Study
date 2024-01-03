package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Println("测试添加用户")
	user := &User{}
	// 调用添加用户方法
	user.AddUser()
	user.AddUser2()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	t.Run("测试添加用户1:", testAddUser)
}

// 如果不是以Test的开头的  默认不执行; 我们可以将它设置成为一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户2")
}

// TestMain可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始: ")
	m.Run()
}

func TestUser_GetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录")
	user := &User{
		ID: 1,
	}
	// 调用获取user方法
	u, _ := user.GetUserByID()
	fmt.Println("得到的user信息", u)
}

func TestUser_GetUserAll(t *testing.T) {
	fmt.Println("测试查询所有记录")
	user := &User{}
	// 调用获取所有user方法
	us, _ := user.GetUserAll()
	// 遍历切片
	for k, v := range us {
		fmt.Printf("第%v个用户是%v \n", k, v)
	}

}
