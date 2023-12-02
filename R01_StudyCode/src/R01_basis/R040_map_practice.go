package main

import "fmt"

func main040() {
	// 案例
	/**
	课堂练习: 演示一个key-value 的value是map的案例
	比如:我们要存放3个学生信息，每个学生有 name和sex 信息
	思路:map[string]map[string]string
	*/
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街~"

	studentMap["stu02"] = make(map[string]string, 3) //这句话不能少!!
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江~"
	fmt.Println(studentMap)          // map[stu01:map[address:北京长安街~ name:tom sex:男] stu02:map[address:上海黄浦江~ name:mary sex:女]]
	fmt.Println(studentMap["stu02"]) // map[address:上海黄浦江~ name:mary sex:女]

	// 遍历
	for k1, v1 := range studentMap {
		fmt.Printf("k1 = %v \n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2 = %v   v2 = %v \n", k2, v2)
		}
	}
	/**
	k1 = stu02
	         k2 = name   v2 = mary
	         k2 = sex   v2 = 女
	         k2 = address   v2 = 上海黄浦江~
	k1 = stu01
	         k2 = name   v2 = tom
	         k2 = sex   v2 = 男
	         k2 = address   v2 = 北京长安街
	*/

}
