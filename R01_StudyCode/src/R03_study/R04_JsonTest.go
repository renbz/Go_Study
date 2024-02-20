package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	unmarshalStruct()
	unmarshalSlice()
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

type Monster struct {
	Name  string `json:"name"` // 反射机制
	Age   int    `json:"age"`
	Skill string `json:"skill"`
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

// 演示将json字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":20,\"name\":\"tom\"}]"

	//定义一个slice
	var slice []map[string]interface{}
	//反序列化，不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}
