package main

import (
	"container/heap"
	"sort"
)

func numsGame(nums []int) []int {
	const mod = 1_000_000_007
	ans := make([]int, len(nums))
	left := hp{}  // 维护较小的一半，大根堆（小根堆取负号）
	right := hp{} // 维护较大的一半，小根堆
	for i, b := range nums {
		b -= i
		if i%2 == 0 {
			heap.Push(&right, -left.pushPop(-b))
			x := right.IntSlice[0] // 中位数
			// 原本要减去 left.sum，但由于 left 所有元素都取负号，所以负负得正改为加法
			ans[i] = (right.sum - x + left.sum) % mod
		} else {
			heap.Push(&left, -right.pushPop(b))
			ans[i] = (right.sum + left.sum) % mod
		}
	}
	return ans
}

type hp struct {
	sort.IntSlice     // 继承 Len, Less, Swap
	sum           int // 堆中元素之和
}
