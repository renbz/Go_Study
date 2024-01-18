package main

func maximumNumberOfStringPairs(words []string) (ans int) {
	mapP := map[int]int{}
	for _, word := range words {
		ans += mapP[int(word[1])*100+int(word[0])]
		mapP[int(word[0])*100+int(word[1])] = 1
	}
	return ans
}
