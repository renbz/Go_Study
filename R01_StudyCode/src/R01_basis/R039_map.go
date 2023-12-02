package main

import "fmt"

func main039() {
	var map1 map[string]string
	// make的作用就是给map分配内存空间
	map1 = make(map[string]string, 10)
	map1["ren"] = "bingzhang"

	map2 := make(map[string]string)
	map3 := map[string]string{
		"ren": "bingzhang",
	}

	fmt.Println(map1) // map[ren:bingzhang]
	fmt.Println(map2) // map[]
	fmt.Println(map3) // map[ren:bingzhang]
}
