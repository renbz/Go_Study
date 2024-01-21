package main

import "strconv"

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n, maxNumInx, idx1, idx2 := len(s), len(s)-1, -1, -1
	for i := n - 1; i >= 0; i-- {
		if s[i] > s[maxNumInx] {
			maxNumInx = i
		} else if s[i] < s[maxNumInx] {
			idx1, idx2 = i, maxNumInx
		}
	}
	if idx1 < 0 {
		return num
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]
	v, _ := strconv.Atoi(string(s))
	return v
}
