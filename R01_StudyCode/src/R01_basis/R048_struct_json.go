package main

import (
	"encoding/json"
	"fmt"
)

func main048() {
	// 创建一个Monster变量
	monster := Monster048{"牛魔王", 500, "飞天"}
	// 将monster变量序列化为json字符串
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr)) // jsonStr {"name":"牛魔王","age":500,"skill":"飞天"}

}

type Monster048 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}
