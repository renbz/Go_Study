package main

import (
	"encoding/json"
	"fmt"
)

func main061() {
	unmarshalMap()
}

func unmarshalStruct() {

	str := "{\"name\":\"牛魔王\",\"age\":500,\"skill\":\"飞天\"}"
	var monster Monster
	// 反序列化
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println(monster)
}
func unmarshalMap() {

	str := "{\"name\":\"牛魔王\",\"age\":500,\"skill\":\"飞天\"}"
	// 定义一个map
	var a map[string]interface{}

	// 反序列化。 反序列化map不需要make, 因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("反序列化后: ", a) // 反序列化后:  map[age:500 name:牛魔王 skill:飞天]
}

type Monster struct {
	Name  string `json:"name"` // 反射机制
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func marshalStruct() string {
	// 创建一个Monster变量
	monster := Monster{"牛魔王", 500, "飞天"}
	// 将monster变量序列化为json字符串
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr)) // jsonStr {"name":"牛魔王","age":500,"skill":"飞天"}
	return string(jsonStr)
}
