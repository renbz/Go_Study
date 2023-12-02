package main

import "fmt"

func main042() {
	// 1. 声明一个map切片
	monsters := make([]map[string]string, 2) // 准备放两个妖怪
	// 2. 增加第一个妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}
	// 再增加 需要使用到append函数,可以动态的增加monster
	// 先定义个monster信息
	newMonster := map[string]string{
		"name": "火云邪神",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters) // [map[age:500 name:牛魔王] map[age:400 name:玉兔精] map[age:200 name:火云邪神]]
}
