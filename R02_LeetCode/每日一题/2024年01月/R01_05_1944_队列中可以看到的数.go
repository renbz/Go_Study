package main

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans, stacks := make([]int, n), make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stacks) > 0 && stacks[len(stacks)-1] <= heights[i] {
			ans[i]++ // 统计能够看到的人数:比他矮的都可以看到
			stacks = stacks[:len(stacks)-1]
		}
		// 栈不为空 最后一个比他高的人,也是他可以看到的
		if len(stacks) > 0 {
			ans[i]++
		}
		stacks = append(stacks, heights[i])
	}
	return ans

}
