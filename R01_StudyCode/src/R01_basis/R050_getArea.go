package main

import "fmt"

func main050() {
	var methodUtil MethodUtil
	areaResult := methodUtil.area(100.5, 10)
	fmt.Println(areaResult)
}

func (mu MethodUtil) area(len float64, width float64) float64 {
	return len * width
}

type MethodUtil struct {
}
