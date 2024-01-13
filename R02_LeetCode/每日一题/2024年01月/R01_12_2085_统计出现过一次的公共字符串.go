package main

func countWords(words1 []string, words2 []string) (ans int) {
	w1, w2 := make(map[string]int), make(map[string]int)
	for _, w := range words1 {
		w1[w]++
	}
	for _, w := range words2 {
		w2[w]++
	}

	for w, cnt := range w1 {

		if cnt == 1 && w2[w] == 1 {
			ans++
		}
	}
	return ans
}
