package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	rCharCount, mCharCount := [26]int{}, [26]int{}

	for _, v := range ransomNote {
		rCharCount[v-'a']++
	}
	for _, v := range magazine {
		mCharCount[v-'a']++
	}
	for i := 0; i <= 25; i++ {
		if rCharCount[i] > mCharCount[i] {
			return false
		}
	}
	return true
}
