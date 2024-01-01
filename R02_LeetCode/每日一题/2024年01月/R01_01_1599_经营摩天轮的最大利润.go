package main

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	var (
		n                              = len(customers)
		waitCnt, profit, mxProfit, ans int
	)
	// 特判
	if boardingCost*4 <= runningCost {
		return -1
	}
	//处理顾客流 、 剩余顾客
	for i := 0; i < n || waitCnt > 0; i++ {
		//顾客还在陆续来
		if i < n {
			waitCnt += customers[i]
		}
		playCnt := min(waitCnt, 4)
		waitCnt -= playCnt
		profit += boardingCost*playCnt - runningCost
		if profit > mxProfit {
			ans = i + 1
			mxProfit = profit
		}
	}
	if mxProfit == 0 {
		return -1
	}
	return ans
}
