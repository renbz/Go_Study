package main

func getMaxRepetitions1(s1 string, n1 int, s2 string, n2 int) int {
	idx, cnt1, cnt2 := 0, 0, 0 // s2指针,s1计数,s2计数
	for cnt1 < n1 {            // 当s1的数量没有超过总个数，就可以继续读取s1
		for i := 0; i < len(s1); i++ { //遍历s1中每一个字符
			if s1[i] == s2[idx] { //如果相等指针后移
				if idx == len(s2)-1 { //如果到最后一个
					cnt2++  //s2计数加一
					idx = 0 //指针指回开头
				} else {
					idx++ // 指针后移
				}
			}
		}
		cnt1++ // 用了一个s1，计数
	}
	return cnt2 / n2
}
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0 // 注意此处直接使用 Ra Rb 的下标，不取模

	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}

	map1, map2 := make(map[int]int), make(map[int]int)
	ans := 0 // 注意，此处存储的是 Ra 中 Sb 的个数，而非 Ra 中 Rb 的个数

	for index1/len1 < n1 { // 遍历整个 Ra
		if index1%len1 == len1-1 { //在 Sa 末尾
			if val, ok := map1[index2%len2]; ok { // 出现了循环，进行快进
				cycleLen := index1/len1 - val/len1                 // 每个循环占多少个 Sa
				cycleNum := (n1 - 1 - index1/len1) / cycleLen      // 还有多少个循环
				cycleS2Num := index2/len2 - map2[index2%len2]/len2 // 每个循环含有多少个 Sb

				index1 += cycleNum * cycleLen * len1 // 将 Ra 快进到相应的位置
				ans += cycleNum * cycleS2Num         // 把快进部分的答案数量加上
			} else { // 第一次，注意存储的是未取模的
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}

		}

		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				ans += 1
			}
			index2 += 1
		}
		index1 += 1
	}
	return ans / n2
}
