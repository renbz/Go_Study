package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1. 声明一个Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2. 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3. 实现Interface
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法就是决定你使用什么标准进行排序
// 1. 按Hero的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}
func main055_1() {
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		// 将hero  append到heroes切片
		heroes = append(heroes, hero)
	}

	fmt.Println(heroes) // [{英雄~46 18} {英雄~43 58} {英雄~80 13} {英雄~29 6} {英雄~90 15} {英雄~41 54} {英雄~90 3} {英雄~25 34} {英雄~18 91} {英雄~51 24}]
	sort.Sort(heroes)
	fmt.Println(heroes) // [{英雄~90 3} {英雄~29 6} {英雄~80 13} {英雄~90 15} {英雄~46 18} {英雄~51 24} {英雄~25 34} {英雄~41 54} {英雄~43 58} {英雄~18 91}]

}

type Monkey struct {
	Name string
}

type LittleMonkey struct {
	Monkey
}

func main055() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	fmt.Println(monkey)
}
