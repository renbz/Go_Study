package main

import "sort"

func minimumRemoval(beans []int) int64 {
	n := len(beans)
	sort.Ints(beans)
	total := 0
	for _, bean := range beans {
		total += bean
	}
	// 最少需要移除的豆子数
	res := total
	for i := 0; i < n; i++ {
		res = min(res, total-beans[i]*(n-i))
	}
	return int64(res)
}
