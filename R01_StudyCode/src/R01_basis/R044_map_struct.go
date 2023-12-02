package main

import "fmt"

func main044() {

	students := make(map[string]Stu044, 10)
	students["no1"] = Stu044{"Tom", 10, 99.4}
	students["no2"] = Stu044{"Bob", 18, 80.4}
	students["no3"] = Stu044{"Mary", 25, 70}
	fmt.Println(students) // map[no1:{Tom 10 99.4} no2:{Bob 18 80.4} no3:{Mary 25 70}]
}

type Stu044 struct {
	name  string
	age   int
	grade float32
}
