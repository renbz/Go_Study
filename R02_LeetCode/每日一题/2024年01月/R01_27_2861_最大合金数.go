package main

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	left, right, ans := 1, int(2e8), 0
	for left <= right {
		mid := (left + right) / 2
		var valid bool
		for i := 0; i < k; i++ {
			var spend int64
			for j := 0; j < n; j++ {
				spend += max(int64(composition[i][j])*int64(mid)-int64(stock[j]), int64(0)) * int64(cost[j])
			}
			if spend <= int64(budget) {
				valid = true
				break
			}
		}
		if valid {
			ans, left = mid, mid+1
		} else {
			right = mid - 1
		}
	}
	return ans
}
