package main

import "slices"

func minimumTime(nums1 []int, nums2 []int, x int) int {
	s1, s2, n := 0, 0, len(nums1)
	id := make([]int, n)
	for i := range id {
		id[i] = i
		s1 += nums1[i]
		s2 += nums2[i]
	}
	// 对下标数组排序，避免破坏 nums1 和 nums2 的对应关系
	slices.SortFunc(id, func(i, j int) int { return nums2[i] - nums2[j] })
	f := make([]int, n+1)
	for i, p := range id {
		a, b := nums1[p], nums2[p]
		for j := i + 1; j > 0; j-- {
			f[j] = max(f[j], f[j-1]+a+b*j)
		}
	}

	for t, v := range f {
		if s1+s2*t-v <= x {
			return t
		}
	}
	return -1
}
