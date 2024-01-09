package main

import "math"

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		d[i] = math.MaxInt
	}
	mp := map[string]int{}
	for _, e := range dictionary {
		mp[e]++
	}
	for i := 1; i <= n; i++ {
		d[i] = d[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if _, ok := mp[s[j:i]]; ok {
				d[i] = min(d[i], d[j])
			}
		}
	}
	return d[n]
}
