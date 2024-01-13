package main

func repeatLimitedString(s string, repeatLimit int) string {
	const N = 26
	count := make([]int, N)
	for _, c := range s {
		count[c-'a']++
	}
	var ans []byte
	var cnt int
	for i, j := N-1, N-2; i >= 0 && j >= 0; {
		if count[i] == 0 {
			cnt = 0
			i = i - 1
		} else if cnt < repeatLimit {
			count[i]--
			ans = append(ans, 'a'+byte(i))
			cnt++
		} else if j >= i || count[j] == 0 {
			j--
		} else {
			count[j]--
			ans = append(ans, 'a'+byte(j))
			cnt = 0
		}
	}
	return string(ans)
}
