package main

import "fmt"

func main007() {

	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Scanln(&sal)
	fmt.Scanln(&isPass)
	fmt.Printf("name=%v, age=%d, float=%f, isPass=%t \n", name, age, sal, isPass)

	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("name=%v, age=%d, float=%f, isPass=%t", name, age, sal, isPass)

}
