package main

import "fmt"

func main041() {
	map1 := make(map[string]string)
	map1["no1"] = "北京"
	map1["no2"] = "天津"
	map1["no3"] = "上海"
	fmt.Println(map1) // map[no1:北京 no2:天津 no3:上海]
	// 修改
	map1["no3"] = "重庆"
	fmt.Println(map1) // map[no1:北京 no2:天津 no3:重庆]
	// 删除
	delete(map1, "no3")
	delete(map1, "no4")
	fmt.Println(map1) // map[no1:北京 no2:天津]
	// 查找
	val, ok := map1["no1"]
	if ok {
		fmt.Printf("存在no1 key; 值为%v \n", val) // 存在no1 key; 值为北京
	} else {
		fmt.Printf("没有no1 key\n")
	}
	// 遍历
	for k, v := range map1 {
		fmt.Printf("key= %v , val= %v \n", k, v) // key= no1 , val= 北京  \n  key= no2 , val= 天津
	}
}
